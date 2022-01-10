// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"
)

// PasswordTokenCreate is the builder for creating a PasswordToken entity.
type PasswordTokenCreate struct {
	config
	mutation *PasswordTokenMutation
	hooks    []Hook
}

// SetHash sets the "hash" field.
func (ptc *PasswordTokenCreate) SetHash(s string) *PasswordTokenCreate {
	ptc.mutation.SetHash(s)
	return ptc
}

// SetCreatedAt sets the "created_at" field.
func (ptc *PasswordTokenCreate) SetCreatedAt(t time.Time) *PasswordTokenCreate {
	ptc.mutation.SetCreatedAt(t)
	return ptc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptc *PasswordTokenCreate) SetNillableCreatedAt(t *time.Time) *PasswordTokenCreate {
	if t != nil {
		ptc.SetCreatedAt(*t)
	}
	return ptc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ptc *PasswordTokenCreate) SetUserID(id int) *PasswordTokenCreate {
	ptc.mutation.SetUserID(id)
	return ptc
}

// SetUser sets the "user" edge to the User entity.
func (ptc *PasswordTokenCreate) SetUser(u *User) *PasswordTokenCreate {
	return ptc.SetUserID(u.ID)
}

// Mutation returns the PasswordTokenMutation object of the builder.
func (ptc *PasswordTokenCreate) Mutation() *PasswordTokenMutation {
	return ptc.mutation
}

// Save creates the PasswordToken in the database.
func (ptc *PasswordTokenCreate) Save(ctx context.Context) (*PasswordToken, error) {
	var (
		err  error
		node *PasswordToken
	)
	ptc.defaults()
	if len(ptc.hooks) == 0 {
		if err = ptc.check(); err != nil {
			return nil, err
		}
		node, err = ptc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasswordTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptc.check(); err != nil {
				return nil, err
			}
			ptc.mutation = mutation
			if node, err = ptc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ptc.hooks) - 1; i >= 0; i-- {
			if ptc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *PasswordTokenCreate) SaveX(ctx context.Context) *PasswordToken {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptc *PasswordTokenCreate) Exec(ctx context.Context) error {
	_, err := ptc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptc *PasswordTokenCreate) ExecX(ctx context.Context) {
	if err := ptc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptc *PasswordTokenCreate) defaults() {
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		v := passwordtoken.DefaultCreatedAt()
		ptc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptc *PasswordTokenCreate) check() error {
	if _, ok := ptc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "hash"`)}
	}
	if v, ok := ptc.mutation.Hash(); ok {
		if err := passwordtoken.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "hash": %w`, err)}
		}
	}
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ptc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (ptc *PasswordTokenCreate) sqlSave(ctx context.Context) (*PasswordToken, error) {
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ptc *PasswordTokenCreate) createSpec() (*PasswordToken, *sqlgraph.CreateSpec) {
	var (
		_node = &PasswordToken{config: ptc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: passwordtoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: passwordtoken.FieldID,
			},
		}
	)
	if value, ok := ptc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: passwordtoken.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := ptc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: passwordtoken.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := ptc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.password_token_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PasswordTokenCreateBulk is the builder for creating many PasswordToken entities in bulk.
type PasswordTokenCreateBulk struct {
	config
	builders []*PasswordTokenCreate
}

// Save creates the PasswordToken entities in the database.
func (ptcb *PasswordTokenCreateBulk) Save(ctx context.Context) ([]*PasswordToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*PasswordToken, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PasswordTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *PasswordTokenCreateBulk) SaveX(ctx context.Context) []*PasswordToken {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptcb *PasswordTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := ptcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcb *PasswordTokenCreateBulk) ExecX(ctx context.Context) {
	if err := ptcb.Exec(ctx); err != nil {
		panic(err)
	}
}
