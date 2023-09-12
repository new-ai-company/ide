package nomad

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/nomad/api"
)

const (
	shortNodeIDLength = 8
	taskRunningState  = "running"
	taskDeadState     = "dead"
	defaultTaskName   = "start"
)

var (
	nomadAddress = os.Getenv("NOMAD_ADDRESS")
	nomadToken   = os.Getenv("NOMAD_TOKEN")
)

type NomadClient struct {
	client *api.Client
}

func InitNomadClient() *NomadClient {
	config := api.Config{
		Address:  nomadAddress,
		SecretID: nomadToken,
	}

	client, err := api.NewClient(&config)
	if err != nil {
		log.Fatalf("Error determining current working dir\n> %s\n", err)
	}

	return &NomadClient{
		client: client,
	}
}

func (n *NomadClient) Close() {
	n.client.Close()
}

type JobInfo struct {
	name   string
	evalID string
	index  uint64
}

func (n *NomadClient) getJobEventStream(ctx context.Context, job JobInfo, meta api.QueryMeta) (<-chan *api.Events, error) {
	topics := map[api.Topic][]string{
		api.TopicAllocation: {job.name},
	}

	streamCtx, streamCancel := context.WithCancel(ctx)
	defer streamCancel()

	eventCh, err := n.client.EventStream().Stream(streamCtx, topics, job.index, &api.QueryOptions{
		Filter:     fmt.Sprintf("EvalID == \"%s\"", job.evalID),
		AllowStale: true,
		// The following commented field could probably be used for improving the event stream handling.
		// WaitIndex:  meta.LastIndex,
		// WaitTime:   timeout,
		NextToken: meta.NextToken,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get Nomad event stream for: %+w", err)
	}

	return eventCh, nil
}

func (n *NomadClient) WaitForJobStart(ctx context.Context, job JobInfo, meta api.QueryMeta) (*api.Allocation, error) {
	jobEvents, err := n.getJobEventStream(ctx, job, meta)
	if err != nil {
		return nil, err
	}

	for event := range jobEvents {
		for _, e := range event.Events {
			alloc, allocErr := e.Allocation()
			if allocErr != nil {
				return nil, fmt.Errorf("cannot retrieve allocations for '%s' job: %+w", job.name, allocErr)
			}

			if alloc.TaskStates[defaultTaskName] == nil {
				continue
			}

			if alloc.TaskStates[defaultTaskName].State == taskDeadState {
				return nil, fmt.Errorf("allocation is %s for '%s' job", alloc.TaskStates[defaultTaskName].State, job.name)
			}

			if alloc.TaskStates[defaultTaskName].State == taskRunningState {
				return alloc, nil
			}

			continue
		}
	}

	return nil, fmt.Errorf("failed retrieving allocations")
}

func (n *NomadClient) WaitForJobFinish(ctx context.Context, job JobInfo, meta api.QueryMeta) (*api.Allocation, error) {
	jobEvents, err := n.getJobEventStream(ctx, job, meta)
	if err != nil {
		return nil, err
	}

	for event := range jobEvents {
		for _, e := range event.Events {
			alloc, allocErr := e.Allocation()
			if allocErr != nil {
				return nil, fmt.Errorf("cannot retrieve allocations for '%s' job: %+w", job.name, allocErr)
			}

			if alloc.TaskStates[defaultTaskName] == nil {
				continue
			}

			if alloc.TaskStates[defaultTaskName].State == taskDeadState {
				if alloc.TaskStates[defaultTaskName].Failed {
					return nil, fmt.Errorf("allocation is %s for '%s' job", alloc.TaskStates[defaultTaskName].State, job.name)
				} else {
					return alloc, nil
				}
			}

			continue
		}
	}

	return nil, fmt.Errorf("failed retrieving allocations")
}
