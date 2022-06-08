package nomad

import (
	"bytes"
	"fmt"
	"path"
	"text/template"

	"github.com/devbookhq/orchestration-services/api/internal/api"
	nomadAPI "github.com/hashicorp/nomad/api"
)

const (
	baseDockerfile = `
FROM alpine:3.16

RUN apk update && apk upgrade
RUN apk add util-linux openrc openssh

COPY devbookd /usr/bin/devbookd
COPY devbookd-init /etc/init.d/devbookd
COPY provision-env.sh provision-env.sh
RUN chmod +x provision-env.sh`
)

func (n *NomadClient) DeleteEnv(codeSnippetID string) error {
	tempName := path.Join(templatesDir, "firecracker-env-deleters.hcl")
	temp, err := template.ParseFiles(tempName)
	if err != nil {
		return fmt.Errorf("failed to parse template file '%s': %s", tempName, err)
	}
	temp = template.Must(temp, err)

	tempVars := struct {
		CodeSnippetID string
		FCEnvsDisk    string
	}{
		CodeSnippetID: codeSnippetID,
		FCEnvsDisk:    fcEnvsDisk,
	}
	var spec bytes.Buffer
	if err := temp.Execute(&spec, tempVars); err != nil {
		return fmt.Errorf("failed to `temp.Execute()`: %s", err)
	}

	job, err := n.client.Jobs().ParseHCL(spec.String(), false)
	if err != nil {
		return fmt.Errorf("failed to parse template '%s': %s", tempName, err)
	}

	_, _, err = n.client.Jobs().Register(job, &nomadAPI.WriteOptions{})
	if err != nil {
		return fmt.Errorf("failed to register 'firecracker-env-deleters/%s' job: %s", tempVars.CodeSnippetID, err)
	}

	return nil
}

func (n *NomadClient) CreateEnv(codeSnippetID string, envTemplate string, deps []string) error {
	dockerfileName := fmt.Sprintf("%s.Dockerfile", envTemplate)
	tname := path.Join(templatesDir, "env-templates", dockerfileName)
	dockerfileTemp, err := template.ParseFiles(tname)
	if err != nil {
		return fmt.Errorf("failed to parse template file '%s': %s", tname, err)
	}

	dockerfileTemp = template.Must(dockerfileTemp, err)

	dockerfileVars := struct {
		Deps           []string
		BaseDockerfile string
	}{
		Deps:           deps,
		BaseDockerfile: baseDockerfile,
	}
	var dockerfile bytes.Buffer
	if err := dockerfileTemp.Execute(&dockerfile, dockerfileVars); err != nil {
		return fmt.Errorf("failed to `dockerfileTemp.Execute()`: %s", err)
	}

	tname = path.Join(templatesDir, "firecracker-envs.hcl")
	envsJobTemp, err := template.New("firecracker-envs.hcl").Funcs(
		template.FuncMap{
			"escapeHCL": escapeHCL,
		},
	).ParseFiles(tname)
	if err != nil {
		return fmt.Errorf("failed to parse template file '%s': %s", tname, err)
	}

	envsJobTemp = template.Must(envsJobTemp, err)

	jobVars := struct {
		CodeSnippetID string
		Dockerfile    string
		FCEnvsDisk    string
		APIKey        string
	}{
		CodeSnippetID: codeSnippetID,
		Dockerfile:    dockerfile.String(),
		FCEnvsDisk:    fcEnvsDisk,
		APIKey:        api.APIAdminKey,
	}
	var jobDef bytes.Buffer
	if err := envsJobTemp.Execute(&jobDef, jobVars); err != nil {
		return fmt.Errorf("failed to `envsJobTemp.Execute()`: %s", err)
	}

	job, err := n.client.Jobs().ParseHCL(jobDef.String(), false)
	if err != nil {
		return fmt.Errorf("failed to parse the `firecracker-envs` HCL job file: %s", err)
	}

	_, _, err = n.client.Jobs().Register(job, &nomadAPI.WriteOptions{})
	if err != nil {
		return fmt.Errorf("failed to register 'firecracker-envs/%s' job: %s", jobVars.CodeSnippetID, err)
	}

	return nil
}

func (n *NomadClient) UpdateEnv(codeSnippetID string, session *api.Session) error {
	tname := path.Join(templatesDir, "firecracker-envs-updater.hcl")
	envsJobTemp, err := template.New("firecracker-envs-updater.hcl").ParseFiles(tname)
	if err != nil {
		return fmt.Errorf("failed to parse template file '%s': %s", tname, err)
	}

	envsJobTemp = template.Must(envsJobTemp, err)

	var sessionID string
	if session != nil {
		sessionID = session.SessionID
	} else {
		sessionID = ""
	}

	jobVars := struct {
		CodeSnippetID string
		FCEnvsDisk    string
		SessionID     string
		APIKey        string
	}{
		CodeSnippetID: codeSnippetID,
		FCEnvsDisk:    fcEnvsDisk,
		SessionID:     sessionID,
		APIKey:        api.APIAdminKey,
	}
	var jobDef bytes.Buffer
	if err := envsJobTemp.Execute(&jobDef, jobVars); err != nil {
		return fmt.Errorf("failed to `envsJobTemp.Execute()`: %s", err)
	}

	job, err := n.client.Jobs().ParseHCL(jobDef.String(), false)
	if err != nil {
		return fmt.Errorf("failed to parse the `firecracker-envs-updater` HCL job file: %s", err)
	}

	_, _, err = n.client.Jobs().Register(job, &nomadAPI.WriteOptions{})
	if err != nil {
		return fmt.Errorf("failed to register 'firecracker-envs-updater/%s' job: %s", jobVars.CodeSnippetID, err)
	}

	return nil
}
