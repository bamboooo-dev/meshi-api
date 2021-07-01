// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/bamboooo-dev/meshi-api/ent/predicate"
	"github.com/bamboooo-dev/meshi-api/ent/restaurant"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeRestaurant = "Restaurant"
)

// RestaurantMutation represents an operation that mutates the Restaurant nodes in the graph.
type RestaurantMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	url           *string
	phone         *string
	price         *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Restaurant, error)
	predicates    []predicate.Restaurant
}

var _ ent.Mutation = (*RestaurantMutation)(nil)

// restaurantOption allows management of the mutation configuration using functional options.
type restaurantOption func(*RestaurantMutation)

// newRestaurantMutation creates new mutation for the Restaurant entity.
func newRestaurantMutation(c config, op Op, opts ...restaurantOption) *RestaurantMutation {
	m := &RestaurantMutation{
		config:        c,
		op:            op,
		typ:           TypeRestaurant,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRestaurantID sets the ID field of the mutation.
func withRestaurantID(id int) restaurantOption {
	return func(m *RestaurantMutation) {
		var (
			err   error
			once  sync.Once
			value *Restaurant
		)
		m.oldValue = func(ctx context.Context) (*Restaurant, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Restaurant.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRestaurant sets the old Restaurant of the mutation.
func withRestaurant(node *Restaurant) restaurantOption {
	return func(m *RestaurantMutation) {
		m.oldValue = func(context.Context) (*Restaurant, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RestaurantMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RestaurantMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *RestaurantMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *RestaurantMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *RestaurantMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Restaurant entity.
// If the Restaurant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaurantMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *RestaurantMutation) ResetName() {
	m.name = nil
}

// SetURL sets the "url" field.
func (m *RestaurantMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *RestaurantMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the Restaurant entity.
// If the Restaurant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaurantMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *RestaurantMutation) ResetURL() {
	m.url = nil
}

// SetPhone sets the "phone" field.
func (m *RestaurantMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *RestaurantMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the Restaurant entity.
// If the Restaurant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaurantMutation) OldPhone(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ResetPhone resets all changes to the "phone" field.
func (m *RestaurantMutation) ResetPhone() {
	m.phone = nil
}

// SetPrice sets the "price" field.
func (m *RestaurantMutation) SetPrice(s string) {
	m.price = &s
}

// Price returns the value of the "price" field in the mutation.
func (m *RestaurantMutation) Price() (r string, exists bool) {
	v := m.price
	if v == nil {
		return
	}
	return *v, true
}

// OldPrice returns the old "price" field's value of the Restaurant entity.
// If the Restaurant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaurantMutation) OldPrice(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPrice is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPrice requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrice: %w", err)
	}
	return oldValue.Price, nil
}

// ResetPrice resets all changes to the "price" field.
func (m *RestaurantMutation) ResetPrice() {
	m.price = nil
}

// Op returns the operation name.
func (m *RestaurantMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Restaurant).
func (m *RestaurantMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RestaurantMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, restaurant.FieldName)
	}
	if m.url != nil {
		fields = append(fields, restaurant.FieldURL)
	}
	if m.phone != nil {
		fields = append(fields, restaurant.FieldPhone)
	}
	if m.price != nil {
		fields = append(fields, restaurant.FieldPrice)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RestaurantMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case restaurant.FieldName:
		return m.Name()
	case restaurant.FieldURL:
		return m.URL()
	case restaurant.FieldPhone:
		return m.Phone()
	case restaurant.FieldPrice:
		return m.Price()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RestaurantMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case restaurant.FieldName:
		return m.OldName(ctx)
	case restaurant.FieldURL:
		return m.OldURL(ctx)
	case restaurant.FieldPhone:
		return m.OldPhone(ctx)
	case restaurant.FieldPrice:
		return m.OldPrice(ctx)
	}
	return nil, fmt.Errorf("unknown Restaurant field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RestaurantMutation) SetField(name string, value ent.Value) error {
	switch name {
	case restaurant.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case restaurant.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case restaurant.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case restaurant.FieldPrice:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrice(v)
		return nil
	}
	return fmt.Errorf("unknown Restaurant field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RestaurantMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RestaurantMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RestaurantMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Restaurant numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RestaurantMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RestaurantMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RestaurantMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Restaurant nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RestaurantMutation) ResetField(name string) error {
	switch name {
	case restaurant.FieldName:
		m.ResetName()
		return nil
	case restaurant.FieldURL:
		m.ResetURL()
		return nil
	case restaurant.FieldPhone:
		m.ResetPhone()
		return nil
	case restaurant.FieldPrice:
		m.ResetPrice()
		return nil
	}
	return fmt.Errorf("unknown Restaurant field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RestaurantMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RestaurantMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RestaurantMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RestaurantMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RestaurantMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RestaurantMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RestaurantMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Restaurant unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RestaurantMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Restaurant edge %s", name)
}
