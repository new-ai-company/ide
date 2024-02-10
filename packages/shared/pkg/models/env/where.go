// Code generated by ent, DO NOT EDIT.

package env

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/internal"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Env {
	return predicate.Env(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Env {
	return predicate.Env(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldUpdatedAt, v))
}

// TeamID applies equality check predicate on the "team_id" field. It's identical to TeamIDEQ.
func TeamID(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldTeamID, v))
}

// Dockerfile applies equality check predicate on the "dockerfile" field. It's identical to DockerfileEQ.
func Dockerfile(v string) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldDockerfile, v))
}

// Public applies equality check predicate on the "public" field. It's identical to PublicEQ.
func Public(v bool) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldPublic, v))
}

// BuildID applies equality check predicate on the "build_id" field. It's identical to BuildIDEQ.
func BuildID(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldBuildID, v))
}

// BuildCount applies equality check predicate on the "build_count" field. It's identical to BuildCountEQ.
func BuildCount(v int32) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldBuildCount, v))
}

// SpawnCount applies equality check predicate on the "spawn_count" field. It's identical to SpawnCountEQ.
func SpawnCount(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldSpawnCount, v))
}

// LastSpawnedAt applies equality check predicate on the "last_spawned_at" field. It's identical to LastSpawnedAtEQ.
func LastSpawnedAt(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldLastSpawnedAt, v))
}

// Vcpu applies equality check predicate on the "vcpu" field. It's identical to VcpuEQ.
func Vcpu(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldVcpu, v))
}

// RAMMB applies equality check predicate on the "ram_mb" field. It's identical to RAMMBEQ.
func RAMMB(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldRAMMB, v))
}

// FreeDiskSizeMB applies equality check predicate on the "free_disk_size_mb" field. It's identical to FreeDiskSizeMBEQ.
func FreeDiskSizeMB(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldFreeDiskSizeMB, v))
}

// TotalDiskSizeMB applies equality check predicate on the "total_disk_size_mb" field. It's identical to TotalDiskSizeMBEQ.
func TotalDiskSizeMB(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldTotalDiskSizeMB, v))
}

// KernelVersion applies equality check predicate on the "kernel_version" field. It's identical to KernelVersionEQ.
func KernelVersion(v string) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldKernelVersion, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldUpdatedAt, v))
}

// TeamIDEQ applies the EQ predicate on the "team_id" field.
func TeamIDEQ(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldTeamID, v))
}

// TeamIDNEQ applies the NEQ predicate on the "team_id" field.
func TeamIDNEQ(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldTeamID, v))
}

// TeamIDIn applies the In predicate on the "team_id" field.
func TeamIDIn(vs ...uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldTeamID, vs...))
}

// TeamIDNotIn applies the NotIn predicate on the "team_id" field.
func TeamIDNotIn(vs ...uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldTeamID, vs...))
}

// DockerfileEQ applies the EQ predicate on the "dockerfile" field.
func DockerfileEQ(v string) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldDockerfile, v))
}

// DockerfileNEQ applies the NEQ predicate on the "dockerfile" field.
func DockerfileNEQ(v string) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldDockerfile, v))
}

// DockerfileIn applies the In predicate on the "dockerfile" field.
func DockerfileIn(vs ...string) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldDockerfile, vs...))
}

// DockerfileNotIn applies the NotIn predicate on the "dockerfile" field.
func DockerfileNotIn(vs ...string) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldDockerfile, vs...))
}

// DockerfileGT applies the GT predicate on the "dockerfile" field.
func DockerfileGT(v string) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldDockerfile, v))
}

// DockerfileGTE applies the GTE predicate on the "dockerfile" field.
func DockerfileGTE(v string) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldDockerfile, v))
}

// DockerfileLT applies the LT predicate on the "dockerfile" field.
func DockerfileLT(v string) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldDockerfile, v))
}

// DockerfileLTE applies the LTE predicate on the "dockerfile" field.
func DockerfileLTE(v string) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldDockerfile, v))
}

// DockerfileContains applies the Contains predicate on the "dockerfile" field.
func DockerfileContains(v string) predicate.Env {
	return predicate.Env(sql.FieldContains(FieldDockerfile, v))
}

// DockerfileHasPrefix applies the HasPrefix predicate on the "dockerfile" field.
func DockerfileHasPrefix(v string) predicate.Env {
	return predicate.Env(sql.FieldHasPrefix(FieldDockerfile, v))
}

// DockerfileHasSuffix applies the HasSuffix predicate on the "dockerfile" field.
func DockerfileHasSuffix(v string) predicate.Env {
	return predicate.Env(sql.FieldHasSuffix(FieldDockerfile, v))
}

// DockerfileEqualFold applies the EqualFold predicate on the "dockerfile" field.
func DockerfileEqualFold(v string) predicate.Env {
	return predicate.Env(sql.FieldEqualFold(FieldDockerfile, v))
}

// DockerfileContainsFold applies the ContainsFold predicate on the "dockerfile" field.
func DockerfileContainsFold(v string) predicate.Env {
	return predicate.Env(sql.FieldContainsFold(FieldDockerfile, v))
}

// PublicEQ applies the EQ predicate on the "public" field.
func PublicEQ(v bool) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldPublic, v))
}

// PublicNEQ applies the NEQ predicate on the "public" field.
func PublicNEQ(v bool) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldPublic, v))
}

// BuildIDEQ applies the EQ predicate on the "build_id" field.
func BuildIDEQ(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldBuildID, v))
}

// BuildIDNEQ applies the NEQ predicate on the "build_id" field.
func BuildIDNEQ(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldBuildID, v))
}

// BuildIDIn applies the In predicate on the "build_id" field.
func BuildIDIn(vs ...uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldBuildID, vs...))
}

// BuildIDNotIn applies the NotIn predicate on the "build_id" field.
func BuildIDNotIn(vs ...uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldBuildID, vs...))
}

// BuildIDGT applies the GT predicate on the "build_id" field.
func BuildIDGT(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldBuildID, v))
}

// BuildIDGTE applies the GTE predicate on the "build_id" field.
func BuildIDGTE(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldBuildID, v))
}

// BuildIDLT applies the LT predicate on the "build_id" field.
func BuildIDLT(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldBuildID, v))
}

// BuildIDLTE applies the LTE predicate on the "build_id" field.
func BuildIDLTE(v uuid.UUID) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldBuildID, v))
}

// BuildCountEQ applies the EQ predicate on the "build_count" field.
func BuildCountEQ(v int32) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldBuildCount, v))
}

// BuildCountNEQ applies the NEQ predicate on the "build_count" field.
func BuildCountNEQ(v int32) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldBuildCount, v))
}

// BuildCountIn applies the In predicate on the "build_count" field.
func BuildCountIn(vs ...int32) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldBuildCount, vs...))
}

// BuildCountNotIn applies the NotIn predicate on the "build_count" field.
func BuildCountNotIn(vs ...int32) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldBuildCount, vs...))
}

// BuildCountGT applies the GT predicate on the "build_count" field.
func BuildCountGT(v int32) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldBuildCount, v))
}

// BuildCountGTE applies the GTE predicate on the "build_count" field.
func BuildCountGTE(v int32) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldBuildCount, v))
}

// BuildCountLT applies the LT predicate on the "build_count" field.
func BuildCountLT(v int32) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldBuildCount, v))
}

// BuildCountLTE applies the LTE predicate on the "build_count" field.
func BuildCountLTE(v int32) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldBuildCount, v))
}

// SpawnCountEQ applies the EQ predicate on the "spawn_count" field.
func SpawnCountEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldSpawnCount, v))
}

// SpawnCountNEQ applies the NEQ predicate on the "spawn_count" field.
func SpawnCountNEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldSpawnCount, v))
}

// SpawnCountIn applies the In predicate on the "spawn_count" field.
func SpawnCountIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldSpawnCount, vs...))
}

// SpawnCountNotIn applies the NotIn predicate on the "spawn_count" field.
func SpawnCountNotIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldSpawnCount, vs...))
}

// SpawnCountGT applies the GT predicate on the "spawn_count" field.
func SpawnCountGT(v int64) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldSpawnCount, v))
}

// SpawnCountGTE applies the GTE predicate on the "spawn_count" field.
func SpawnCountGTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldSpawnCount, v))
}

// SpawnCountLT applies the LT predicate on the "spawn_count" field.
func SpawnCountLT(v int64) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldSpawnCount, v))
}

// SpawnCountLTE applies the LTE predicate on the "spawn_count" field.
func SpawnCountLTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldSpawnCount, v))
}

// LastSpawnedAtEQ applies the EQ predicate on the "last_spawned_at" field.
func LastSpawnedAtEQ(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldLastSpawnedAt, v))
}

// LastSpawnedAtNEQ applies the NEQ predicate on the "last_spawned_at" field.
func LastSpawnedAtNEQ(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldLastSpawnedAt, v))
}

// LastSpawnedAtIn applies the In predicate on the "last_spawned_at" field.
func LastSpawnedAtIn(vs ...time.Time) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldLastSpawnedAt, vs...))
}

// LastSpawnedAtNotIn applies the NotIn predicate on the "last_spawned_at" field.
func LastSpawnedAtNotIn(vs ...time.Time) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldLastSpawnedAt, vs...))
}

// LastSpawnedAtGT applies the GT predicate on the "last_spawned_at" field.
func LastSpawnedAtGT(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldLastSpawnedAt, v))
}

// LastSpawnedAtGTE applies the GTE predicate on the "last_spawned_at" field.
func LastSpawnedAtGTE(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldLastSpawnedAt, v))
}

// LastSpawnedAtLT applies the LT predicate on the "last_spawned_at" field.
func LastSpawnedAtLT(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldLastSpawnedAt, v))
}

// LastSpawnedAtLTE applies the LTE predicate on the "last_spawned_at" field.
func LastSpawnedAtLTE(v time.Time) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldLastSpawnedAt, v))
}

// LastSpawnedAtIsNil applies the IsNil predicate on the "last_spawned_at" field.
func LastSpawnedAtIsNil() predicate.Env {
	return predicate.Env(sql.FieldIsNull(FieldLastSpawnedAt))
}

// LastSpawnedAtNotNil applies the NotNil predicate on the "last_spawned_at" field.
func LastSpawnedAtNotNil() predicate.Env {
	return predicate.Env(sql.FieldNotNull(FieldLastSpawnedAt))
}

// VcpuEQ applies the EQ predicate on the "vcpu" field.
func VcpuEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldVcpu, v))
}

// VcpuNEQ applies the NEQ predicate on the "vcpu" field.
func VcpuNEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldVcpu, v))
}

// VcpuIn applies the In predicate on the "vcpu" field.
func VcpuIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldVcpu, vs...))
}

// VcpuNotIn applies the NotIn predicate on the "vcpu" field.
func VcpuNotIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldVcpu, vs...))
}

// VcpuGT applies the GT predicate on the "vcpu" field.
func VcpuGT(v int64) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldVcpu, v))
}

// VcpuGTE applies the GTE predicate on the "vcpu" field.
func VcpuGTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldVcpu, v))
}

// VcpuLT applies the LT predicate on the "vcpu" field.
func VcpuLT(v int64) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldVcpu, v))
}

// VcpuLTE applies the LTE predicate on the "vcpu" field.
func VcpuLTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldVcpu, v))
}

// RAMMBEQ applies the EQ predicate on the "ram_mb" field.
func RAMMBEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldRAMMB, v))
}

// RAMMBNEQ applies the NEQ predicate on the "ram_mb" field.
func RAMMBNEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldRAMMB, v))
}

// RAMMBIn applies the In predicate on the "ram_mb" field.
func RAMMBIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldRAMMB, vs...))
}

// RAMMBNotIn applies the NotIn predicate on the "ram_mb" field.
func RAMMBNotIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldRAMMB, vs...))
}

// RAMMBGT applies the GT predicate on the "ram_mb" field.
func RAMMBGT(v int64) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldRAMMB, v))
}

// RAMMBGTE applies the GTE predicate on the "ram_mb" field.
func RAMMBGTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldRAMMB, v))
}

// RAMMBLT applies the LT predicate on the "ram_mb" field.
func RAMMBLT(v int64) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldRAMMB, v))
}

// RAMMBLTE applies the LTE predicate on the "ram_mb" field.
func RAMMBLTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldRAMMB, v))
}

// FreeDiskSizeMBEQ applies the EQ predicate on the "free_disk_size_mb" field.
func FreeDiskSizeMBEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldFreeDiskSizeMB, v))
}

// FreeDiskSizeMBNEQ applies the NEQ predicate on the "free_disk_size_mb" field.
func FreeDiskSizeMBNEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldFreeDiskSizeMB, v))
}

// FreeDiskSizeMBIn applies the In predicate on the "free_disk_size_mb" field.
func FreeDiskSizeMBIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldFreeDiskSizeMB, vs...))
}

// FreeDiskSizeMBNotIn applies the NotIn predicate on the "free_disk_size_mb" field.
func FreeDiskSizeMBNotIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldFreeDiskSizeMB, vs...))
}

// FreeDiskSizeMBGT applies the GT predicate on the "free_disk_size_mb" field.
func FreeDiskSizeMBGT(v int64) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldFreeDiskSizeMB, v))
}

// FreeDiskSizeMBGTE applies the GTE predicate on the "free_disk_size_mb" field.
func FreeDiskSizeMBGTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldFreeDiskSizeMB, v))
}

// FreeDiskSizeMBLT applies the LT predicate on the "free_disk_size_mb" field.
func FreeDiskSizeMBLT(v int64) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldFreeDiskSizeMB, v))
}

// FreeDiskSizeMBLTE applies the LTE predicate on the "free_disk_size_mb" field.
func FreeDiskSizeMBLTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldFreeDiskSizeMB, v))
}

// TotalDiskSizeMBEQ applies the EQ predicate on the "total_disk_size_mb" field.
func TotalDiskSizeMBEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldTotalDiskSizeMB, v))
}

// TotalDiskSizeMBNEQ applies the NEQ predicate on the "total_disk_size_mb" field.
func TotalDiskSizeMBNEQ(v int64) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldTotalDiskSizeMB, v))
}

// TotalDiskSizeMBIn applies the In predicate on the "total_disk_size_mb" field.
func TotalDiskSizeMBIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldTotalDiskSizeMB, vs...))
}

// TotalDiskSizeMBNotIn applies the NotIn predicate on the "total_disk_size_mb" field.
func TotalDiskSizeMBNotIn(vs ...int64) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldTotalDiskSizeMB, vs...))
}

// TotalDiskSizeMBGT applies the GT predicate on the "total_disk_size_mb" field.
func TotalDiskSizeMBGT(v int64) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldTotalDiskSizeMB, v))
}

// TotalDiskSizeMBGTE applies the GTE predicate on the "total_disk_size_mb" field.
func TotalDiskSizeMBGTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldTotalDiskSizeMB, v))
}

// TotalDiskSizeMBLT applies the LT predicate on the "total_disk_size_mb" field.
func TotalDiskSizeMBLT(v int64) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldTotalDiskSizeMB, v))
}

// TotalDiskSizeMBLTE applies the LTE predicate on the "total_disk_size_mb" field.
func TotalDiskSizeMBLTE(v int64) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldTotalDiskSizeMB, v))
}

// KernelVersionEQ applies the EQ predicate on the "kernel_version" field.
func KernelVersionEQ(v string) predicate.Env {
	return predicate.Env(sql.FieldEQ(FieldKernelVersion, v))
}

// KernelVersionNEQ applies the NEQ predicate on the "kernel_version" field.
func KernelVersionNEQ(v string) predicate.Env {
	return predicate.Env(sql.FieldNEQ(FieldKernelVersion, v))
}

// KernelVersionIn applies the In predicate on the "kernel_version" field.
func KernelVersionIn(vs ...string) predicate.Env {
	return predicate.Env(sql.FieldIn(FieldKernelVersion, vs...))
}

// KernelVersionNotIn applies the NotIn predicate on the "kernel_version" field.
func KernelVersionNotIn(vs ...string) predicate.Env {
	return predicate.Env(sql.FieldNotIn(FieldKernelVersion, vs...))
}

// KernelVersionGT applies the GT predicate on the "kernel_version" field.
func KernelVersionGT(v string) predicate.Env {
	return predicate.Env(sql.FieldGT(FieldKernelVersion, v))
}

// KernelVersionGTE applies the GTE predicate on the "kernel_version" field.
func KernelVersionGTE(v string) predicate.Env {
	return predicate.Env(sql.FieldGTE(FieldKernelVersion, v))
}

// KernelVersionLT applies the LT predicate on the "kernel_version" field.
func KernelVersionLT(v string) predicate.Env {
	return predicate.Env(sql.FieldLT(FieldKernelVersion, v))
}

// KernelVersionLTE applies the LTE predicate on the "kernel_version" field.
func KernelVersionLTE(v string) predicate.Env {
	return predicate.Env(sql.FieldLTE(FieldKernelVersion, v))
}

// KernelVersionContains applies the Contains predicate on the "kernel_version" field.
func KernelVersionContains(v string) predicate.Env {
	return predicate.Env(sql.FieldContains(FieldKernelVersion, v))
}

// KernelVersionHasPrefix applies the HasPrefix predicate on the "kernel_version" field.
func KernelVersionHasPrefix(v string) predicate.Env {
	return predicate.Env(sql.FieldHasPrefix(FieldKernelVersion, v))
}

// KernelVersionHasSuffix applies the HasSuffix predicate on the "kernel_version" field.
func KernelVersionHasSuffix(v string) predicate.Env {
	return predicate.Env(sql.FieldHasSuffix(FieldKernelVersion, v))
}

// KernelVersionEqualFold applies the EqualFold predicate on the "kernel_version" field.
func KernelVersionEqualFold(v string) predicate.Env {
	return predicate.Env(sql.FieldEqualFold(FieldKernelVersion, v))
}

// KernelVersionContainsFold applies the ContainsFold predicate on the "kernel_version" field.
func KernelVersionContainsFold(v string) predicate.Env {
	return predicate.Env(sql.FieldContainsFold(FieldKernelVersion, v))
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.Env {
	return predicate.Env(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeamTable, TeamColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Team
		step.Edge.Schema = schemaConfig.Env
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.Team) predicate.Env {
	return predicate.Env(func(s *sql.Selector) {
		step := newTeamStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Team
		step.Edge.Schema = schemaConfig.Env
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEnvAliases applies the HasEdge predicate on the "env_aliases" edge.
func HasEnvAliases() predicate.Env {
	return predicate.Env(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EnvAliasesTable, EnvAliasesColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.EnvAlias
		step.Edge.Schema = schemaConfig.EnvAlias
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEnvAliasesWith applies the HasEdge predicate on the "env_aliases" edge with a given conditions (other predicates).
func HasEnvAliasesWith(preds ...predicate.EnvAlias) predicate.Env {
	return predicate.Env(func(s *sql.Selector) {
		step := newEnvAliasesStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.EnvAlias
		step.Edge.Schema = schemaConfig.EnvAlias
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Env) predicate.Env {
	return predicate.Env(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Env) predicate.Env {
	return predicate.Env(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Env) predicate.Env {
	return predicate.Env(sql.NotPredicates(p))
}
