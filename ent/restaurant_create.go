// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bamboooo-dev/meshi-api/ent/like"
	"github.com/bamboooo-dev/meshi-api/ent/restaurant"
)

// RestaurantCreate is the builder for creating a Restaurant entity.
type RestaurantCreate struct {
	config
	mutation *RestaurantMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *RestaurantCreate) SetName(s string) *RestaurantCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetURL sets the "url" field.
func (rc *RestaurantCreate) SetURL(s string) *RestaurantCreate {
	rc.mutation.SetURL(s)
	return rc
}

// SetPhone sets the "phone" field.
func (rc *RestaurantCreate) SetPhone(s string) *RestaurantCreate {
	rc.mutation.SetPhone(s)
	return rc
}

// SetPrice sets the "price" field.
func (rc *RestaurantCreate) SetPrice(s string) *RestaurantCreate {
	rc.mutation.SetPrice(s)
	return rc
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (rc *RestaurantCreate) AddLikeIDs(ids ...int) *RestaurantCreate {
	rc.mutation.AddLikeIDs(ids...)
	return rc
}

// AddLikes adds the "likes" edges to the Like entity.
func (rc *RestaurantCreate) AddLikes(l ...*Like) *RestaurantCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return rc.AddLikeIDs(ids...)
}

// Mutation returns the RestaurantMutation object of the builder.
func (rc *RestaurantCreate) Mutation() *RestaurantMutation {
	return rc.mutation
}

// Save creates the Restaurant in the database.
func (rc *RestaurantCreate) Save(ctx context.Context) (*Restaurant, error) {
	var (
		err  error
		node *Restaurant
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RestaurantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RestaurantCreate) SaveX(ctx context.Context) *Restaurant {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (rc *RestaurantCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := restaurant.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := rc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	if _, ok := rc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New("ent: missing required field \"phone\"")}
	}
	if _, ok := rc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	return nil
}

func (rc *RestaurantCreate) sqlSave(ctx context.Context) (*Restaurant, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RestaurantCreate) createSpec() (*Restaurant, *sqlgraph.CreateSpec) {
	var (
		_node = &Restaurant{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: restaurant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: restaurant.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := rc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := rc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldPrice,
		})
		_node.Price = value
	}
	if nodes := rc.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   restaurant.LikesTable,
			Columns: []string{restaurant.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: like.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RestaurantCreateBulk is the builder for creating many Restaurant entities in bulk.
type RestaurantCreateBulk struct {
	config
	builders []*RestaurantCreate
}

// Save creates the Restaurant entities in the database.
func (rcb *RestaurantCreateBulk) Save(ctx context.Context) ([]*Restaurant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Restaurant, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RestaurantMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RestaurantCreateBulk) SaveX(ctx context.Context) []*Restaurant {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
