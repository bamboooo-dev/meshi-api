// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bamboooo-dev/meshi-api/ent/like"
	"github.com/bamboooo-dev/meshi-api/ent/predicate"
	"github.com/bamboooo-dev/meshi-api/ent/restaurant"
)

// RestaurantUpdate is the builder for updating Restaurant entities.
type RestaurantUpdate struct {
	config
	hooks    []Hook
	mutation *RestaurantMutation
}

// Where adds a new predicate for the RestaurantUpdate builder.
func (ru *RestaurantUpdate) Where(ps ...predicate.Restaurant) *RestaurantUpdate {
	ru.mutation.predicates = append(ru.mutation.predicates, ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RestaurantUpdate) SetName(s string) *RestaurantUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetURL sets the "url" field.
func (ru *RestaurantUpdate) SetURL(s string) *RestaurantUpdate {
	ru.mutation.SetURL(s)
	return ru
}

// SetPhone sets the "phone" field.
func (ru *RestaurantUpdate) SetPhone(s string) *RestaurantUpdate {
	ru.mutation.SetPhone(s)
	return ru
}

// SetPrice sets the "price" field.
func (ru *RestaurantUpdate) SetPrice(s string) *RestaurantUpdate {
	ru.mutation.SetPrice(s)
	return ru
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (ru *RestaurantUpdate) AddLikeIDs(ids ...int) *RestaurantUpdate {
	ru.mutation.AddLikeIDs(ids...)
	return ru
}

// AddLikes adds the "likes" edges to the Like entity.
func (ru *RestaurantUpdate) AddLikes(l ...*Like) *RestaurantUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.AddLikeIDs(ids...)
}

// Mutation returns the RestaurantMutation object of the builder.
func (ru *RestaurantUpdate) Mutation() *RestaurantMutation {
	return ru.mutation
}

// ClearLikes clears all "likes" edges to the Like entity.
func (ru *RestaurantUpdate) ClearLikes() *RestaurantUpdate {
	ru.mutation.ClearLikes()
	return ru
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (ru *RestaurantUpdate) RemoveLikeIDs(ids ...int) *RestaurantUpdate {
	ru.mutation.RemoveLikeIDs(ids...)
	return ru
}

// RemoveLikes removes "likes" edges to Like entities.
func (ru *RestaurantUpdate) RemoveLikes(l ...*Like) *RestaurantUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.RemoveLikeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RestaurantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RestaurantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RestaurantUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RestaurantUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RestaurantUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RestaurantUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := restaurant.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (ru *RestaurantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   restaurant.Table,
			Columns: restaurant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: restaurant.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldName,
		})
	}
	if value, ok := ru.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldURL,
		})
	}
	if value, ok := ru.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldPhone,
		})
	}
	if value, ok := ru.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldPrice,
		})
	}
	if ru.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedLikesIDs(); len(nodes) > 0 && !ru.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.LikesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{restaurant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RestaurantUpdateOne is the builder for updating a single Restaurant entity.
type RestaurantUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RestaurantMutation
}

// SetName sets the "name" field.
func (ruo *RestaurantUpdateOne) SetName(s string) *RestaurantUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetURL sets the "url" field.
func (ruo *RestaurantUpdateOne) SetURL(s string) *RestaurantUpdateOne {
	ruo.mutation.SetURL(s)
	return ruo
}

// SetPhone sets the "phone" field.
func (ruo *RestaurantUpdateOne) SetPhone(s string) *RestaurantUpdateOne {
	ruo.mutation.SetPhone(s)
	return ruo
}

// SetPrice sets the "price" field.
func (ruo *RestaurantUpdateOne) SetPrice(s string) *RestaurantUpdateOne {
	ruo.mutation.SetPrice(s)
	return ruo
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (ruo *RestaurantUpdateOne) AddLikeIDs(ids ...int) *RestaurantUpdateOne {
	ruo.mutation.AddLikeIDs(ids...)
	return ruo
}

// AddLikes adds the "likes" edges to the Like entity.
func (ruo *RestaurantUpdateOne) AddLikes(l ...*Like) *RestaurantUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.AddLikeIDs(ids...)
}

// Mutation returns the RestaurantMutation object of the builder.
func (ruo *RestaurantUpdateOne) Mutation() *RestaurantMutation {
	return ruo.mutation
}

// ClearLikes clears all "likes" edges to the Like entity.
func (ruo *RestaurantUpdateOne) ClearLikes() *RestaurantUpdateOne {
	ruo.mutation.ClearLikes()
	return ruo
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (ruo *RestaurantUpdateOne) RemoveLikeIDs(ids ...int) *RestaurantUpdateOne {
	ruo.mutation.RemoveLikeIDs(ids...)
	return ruo
}

// RemoveLikes removes "likes" edges to Like entities.
func (ruo *RestaurantUpdateOne) RemoveLikes(l ...*Like) *RestaurantUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.RemoveLikeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RestaurantUpdateOne) Select(field string, fields ...string) *RestaurantUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Restaurant entity.
func (ruo *RestaurantUpdateOne) Save(ctx context.Context) (*Restaurant, error) {
	var (
		err  error
		node *Restaurant
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RestaurantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RestaurantUpdateOne) SaveX(ctx context.Context) *Restaurant {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RestaurantUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RestaurantUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RestaurantUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := restaurant.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (ruo *RestaurantUpdateOne) sqlSave(ctx context.Context) (_node *Restaurant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   restaurant.Table,
			Columns: restaurant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: restaurant.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Restaurant.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, restaurant.FieldID)
		for _, f := range fields {
			if !restaurant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != restaurant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldName,
		})
	}
	if value, ok := ruo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldURL,
		})
	}
	if value, ok := ruo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldPhone,
		})
	}
	if value, ok := ruo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurant.FieldPrice,
		})
	}
	if ruo.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedLikesIDs(); len(nodes) > 0 && !ruo.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.LikesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Restaurant{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{restaurant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
