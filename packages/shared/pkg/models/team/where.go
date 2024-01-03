// Code generated by ent, DO NOT EDIT.

package team

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/internal"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Team {
	return predicate.Team(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Team {
	return predicate.Team(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Team {
	return predicate.Team(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Team {
	return predicate.Team(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Team {
	return predicate.Team(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Team {
	return predicate.Team(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldCreatedAt, v))
}

// IsDefault applies equality check predicate on the "is_default" field. It's identical to IsDefaultEQ.
func IsDefault(v bool) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldIsDefault, v))
}

// IsBlocked applies equality check predicate on the "is_blocked" field. It's identical to IsBlockedEQ.
func IsBlocked(v bool) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldIsBlocked, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldName, v))
}

// Tier applies equality check predicate on the "tier" field. It's identical to TierEQ.
func Tier(v string) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldTier, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldEmail, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Team {
	return predicate.Team(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Team {
	return predicate.Team(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Team {
	return predicate.Team(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Team {
	return predicate.Team(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Team {
	return predicate.Team(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Team {
	return predicate.Team(sql.FieldLTE(FieldCreatedAt, v))
}

// IsDefaultEQ applies the EQ predicate on the "is_default" field.
func IsDefaultEQ(v bool) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldIsDefault, v))
}

// IsDefaultNEQ applies the NEQ predicate on the "is_default" field.
func IsDefaultNEQ(v bool) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldIsDefault, v))
}

// IsBlockedEQ applies the EQ predicate on the "is_blocked" field.
func IsBlockedEQ(v bool) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldIsBlocked, v))
}

// IsBlockedNEQ applies the NEQ predicate on the "is_blocked" field.
func IsBlockedNEQ(v bool) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldIsBlocked, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Team {
	return predicate.Team(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Team {
	return predicate.Team(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Team {
	return predicate.Team(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Team {
	return predicate.Team(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Team {
	return predicate.Team(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Team {
	return predicate.Team(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Team {
	return predicate.Team(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Team {
	return predicate.Team(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Team {
	return predicate.Team(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Team {
	return predicate.Team(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Team {
	return predicate.Team(sql.FieldContainsFold(FieldName, v))
}

// TierEQ applies the EQ predicate on the "tier" field.
func TierEQ(v string) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldTier, v))
}

// TierNEQ applies the NEQ predicate on the "tier" field.
func TierNEQ(v string) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldTier, v))
}

// TierIn applies the In predicate on the "tier" field.
func TierIn(vs ...string) predicate.Team {
	return predicate.Team(sql.FieldIn(FieldTier, vs...))
}

// TierNotIn applies the NotIn predicate on the "tier" field.
func TierNotIn(vs ...string) predicate.Team {
	return predicate.Team(sql.FieldNotIn(FieldTier, vs...))
}

// TierGT applies the GT predicate on the "tier" field.
func TierGT(v string) predicate.Team {
	return predicate.Team(sql.FieldGT(FieldTier, v))
}

// TierGTE applies the GTE predicate on the "tier" field.
func TierGTE(v string) predicate.Team {
	return predicate.Team(sql.FieldGTE(FieldTier, v))
}

// TierLT applies the LT predicate on the "tier" field.
func TierLT(v string) predicate.Team {
	return predicate.Team(sql.FieldLT(FieldTier, v))
}

// TierLTE applies the LTE predicate on the "tier" field.
func TierLTE(v string) predicate.Team {
	return predicate.Team(sql.FieldLTE(FieldTier, v))
}

// TierContains applies the Contains predicate on the "tier" field.
func TierContains(v string) predicate.Team {
	return predicate.Team(sql.FieldContains(FieldTier, v))
}

// TierHasPrefix applies the HasPrefix predicate on the "tier" field.
func TierHasPrefix(v string) predicate.Team {
	return predicate.Team(sql.FieldHasPrefix(FieldTier, v))
}

// TierHasSuffix applies the HasSuffix predicate on the "tier" field.
func TierHasSuffix(v string) predicate.Team {
	return predicate.Team(sql.FieldHasSuffix(FieldTier, v))
}

// TierEqualFold applies the EqualFold predicate on the "tier" field.
func TierEqualFold(v string) predicate.Team {
	return predicate.Team(sql.FieldEqualFold(FieldTier, v))
}

// TierContainsFold applies the ContainsFold predicate on the "tier" field.
func TierContainsFold(v string) predicate.Team {
	return predicate.Team(sql.FieldContainsFold(FieldTier, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Team {
	return predicate.Team(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Team {
	return predicate.Team(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Team {
	return predicate.Team(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Team {
	return predicate.Team(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Team {
	return predicate.Team(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Team {
	return predicate.Team(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Team {
	return predicate.Team(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Team {
	return predicate.Team(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Team {
	return predicate.Team(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Team {
	return predicate.Team(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Team {
	return predicate.Team(sql.FieldContainsFold(FieldEmail, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UsersTeams
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := newUsersStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UsersTeams
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeamAPIKeys applies the HasEdge predicate on the "team_api_keys" edge.
func HasTeamAPIKeys() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeamAPIKeysTable, TeamAPIKeysColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TeamAPIKey
		step.Edge.Schema = schemaConfig.TeamAPIKey
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamAPIKeysWith applies the HasEdge predicate on the "team_api_keys" edge with a given conditions (other predicates).
func HasTeamAPIKeysWith(preds ...predicate.TeamAPIKey) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := newTeamAPIKeysStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TeamAPIKey
		step.Edge.Schema = schemaConfig.TeamAPIKey
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeamTier applies the HasEdge predicate on the "team_tier" edge.
func HasTeamTier() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeamTierTable, TeamTierColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Tier
		step.Edge.Schema = schemaConfig.Team
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamTierWith applies the HasEdge predicate on the "team_tier" edge with a given conditions (other predicates).
func HasTeamTierWith(preds ...predicate.Tier) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := newTeamTierStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Tier
		step.Edge.Schema = schemaConfig.Team
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEnvs applies the HasEdge predicate on the "envs" edge.
func HasEnvs() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EnvsTable, EnvsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Env
		step.Edge.Schema = schemaConfig.Env
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEnvsWith applies the HasEdge predicate on the "envs" edge with a given conditions (other predicates).
func HasEnvsWith(preds ...predicate.Env) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := newEnvsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Env
		step.Edge.Schema = schemaConfig.Env
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsersTeams applies the HasEdge predicate on the "users_teams" edge.
func HasUsersTeams() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UsersTeamsTable, UsersTeamsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.UsersTeams
		step.Edge.Schema = schemaConfig.UsersTeams
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersTeamsWith applies the HasEdge predicate on the "users_teams" edge with a given conditions (other predicates).
func HasUsersTeamsWith(preds ...predicate.UsersTeams) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := newUsersTeamsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.UsersTeams
		step.Edge.Schema = schemaConfig.UsersTeams
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Team) predicate.Team {
	return predicate.Team(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Team) predicate.Team {
	return predicate.Team(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Team) predicate.Team {
	return predicate.Team(sql.NotPredicates(p))
}
