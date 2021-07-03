// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/bamboooo-dev/meshi-api/ent/like"
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
	TypeLike       = "Like"
	TypeRestaurant = "Restaurant"
)

// LikeMutation represents an operation that mutates the Like nodes in the graph.
type LikeMutation struct {
	config
	op                Op
	typ               string
	id                *int
	user_id           *string
	clearedFields     map[string]struct{}
	restaurant        *int
	clearedrestaurant bool
	done              bool
	oldValue          func(context.Context) (*Like, error)
	predicates        []predicate.Like
}

var _ ent.Mutation = (*LikeMutation)(nil)

// likeOption allows management of the mutation configuration using functional options.
type likeOption func(*LikeMutation)

// newLikeMutation creates new mutation for the Like entity.
func newLikeMutation(c config, op Op, opts ...likeOption) *LikeMutation {
	m := &LikeMutation{
		config:        c,
		op:            op,
		typ:           TypeLike,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLikeID sets the ID field of the mutation.
func withLikeID(id int) likeOption {
	return func(m *LikeMutation) {
		var (
			err   error
			once  sync.Once
			value *Like
		)
		m.oldValue = func(ctx context.Context) (*Like, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Like.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLike sets the old Like of the mutation.
func withLike(node *Like) likeOption {
	return func(m *LikeMutation) {
		m.oldValue = func(context.Context) (*Like, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LikeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LikeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LikeMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUserID sets the "user_id" field.
func (m *LikeMutation) SetUserID(s string) {
	m.user_id = &s
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *LikeMutation) UserID() (r string, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Like entity.
// If the Like object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LikeMutation) OldUserID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID resets all changes to the "user_id" field.
func (m *LikeMutation) ResetUserID() {
	m.user_id = nil
}

// SetRestaurantID sets the "restaurant" edge to the Restaurant entity by id.
func (m *LikeMutation) SetRestaurantID(id int) {
	m.restaurant = &id
}

// ClearRestaurant clears the "restaurant" edge to the Restaurant entity.
func (m *LikeMutation) ClearRestaurant() {
	m.clearedrestaurant = true
}

// RestaurantCleared reports if the "restaurant" edge to the Restaurant entity was cleared.
func (m *LikeMutation) RestaurantCleared() bool {
	return m.clearedrestaurant
}

// RestaurantID returns the "restaurant" edge ID in the mutation.
func (m *LikeMutation) RestaurantID() (id int, exists bool) {
	if m.restaurant != nil {
		return *m.restaurant, true
	}
	return
}

// RestaurantIDs returns the "restaurant" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// RestaurantID instead. It exists only for internal usage by the builders.
func (m *LikeMutation) RestaurantIDs() (ids []int) {
	if id := m.restaurant; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetRestaurant resets all changes to the "restaurant" edge.
func (m *LikeMutation) ResetRestaurant() {
	m.restaurant = nil
	m.clearedrestaurant = false
}

// Op returns the operation name.
func (m *LikeMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Like).
func (m *LikeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LikeMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.user_id != nil {
		fields = append(fields, like.FieldUserID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LikeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case like.FieldUserID:
		return m.UserID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LikeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case like.FieldUserID:
		return m.OldUserID(ctx)
	}
	return nil, fmt.Errorf("unknown Like field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LikeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case like.FieldUserID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	}
	return fmt.Errorf("unknown Like field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LikeMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LikeMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LikeMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Like numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LikeMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LikeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LikeMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Like nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LikeMutation) ResetField(name string) error {
	switch name {
	case like.FieldUserID:
		m.ResetUserID()
		return nil
	}
	return fmt.Errorf("unknown Like field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LikeMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.restaurant != nil {
		edges = append(edges, like.EdgeRestaurant)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LikeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case like.EdgeRestaurant:
		if id := m.restaurant; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LikeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LikeMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LikeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedrestaurant {
		edges = append(edges, like.EdgeRestaurant)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LikeMutation) EdgeCleared(name string) bool {
	switch name {
	case like.EdgeRestaurant:
		return m.clearedrestaurant
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LikeMutation) ClearEdge(name string) error {
	switch name {
	case like.EdgeRestaurant:
		m.ClearRestaurant()
		return nil
	}
	return fmt.Errorf("unknown Like unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LikeMutation) ResetEdge(name string) error {
	switch name {
	case like.EdgeRestaurant:
		m.ResetRestaurant()
		return nil
	}
	return fmt.Errorf("unknown Like edge %s", name)
}

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
	likes         map[int]struct{}
	removedlikes  map[int]struct{}
	clearedlikes  bool
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

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
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

// AddLikeIDs adds the "likes" edge to the Like entity by ids.
func (m *RestaurantMutation) AddLikeIDs(ids ...int) {
	if m.likes == nil {
		m.likes = make(map[int]struct{})
	}
	for i := range ids {
		m.likes[ids[i]] = struct{}{}
	}
}

// ClearLikes clears the "likes" edge to the Like entity.
func (m *RestaurantMutation) ClearLikes() {
	m.clearedlikes = true
}

// LikesCleared reports if the "likes" edge to the Like entity was cleared.
func (m *RestaurantMutation) LikesCleared() bool {
	return m.clearedlikes
}

// RemoveLikeIDs removes the "likes" edge to the Like entity by IDs.
func (m *RestaurantMutation) RemoveLikeIDs(ids ...int) {
	if m.removedlikes == nil {
		m.removedlikes = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.likes, ids[i])
		m.removedlikes[ids[i]] = struct{}{}
	}
}

// RemovedLikes returns the removed IDs of the "likes" edge to the Like entity.
func (m *RestaurantMutation) RemovedLikesIDs() (ids []int) {
	for id := range m.removedlikes {
		ids = append(ids, id)
	}
	return
}

// LikesIDs returns the "likes" edge IDs in the mutation.
func (m *RestaurantMutation) LikesIDs() (ids []int) {
	for id := range m.likes {
		ids = append(ids, id)
	}
	return
}

// ResetLikes resets all changes to the "likes" edge.
func (m *RestaurantMutation) ResetLikes() {
	m.likes = nil
	m.clearedlikes = false
	m.removedlikes = nil
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
	edges := make([]string, 0, 1)
	if m.likes != nil {
		edges = append(edges, restaurant.EdgeLikes)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RestaurantMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case restaurant.EdgeLikes:
		ids := make([]ent.Value, 0, len(m.likes))
		for id := range m.likes {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RestaurantMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedlikes != nil {
		edges = append(edges, restaurant.EdgeLikes)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RestaurantMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case restaurant.EdgeLikes:
		ids := make([]ent.Value, 0, len(m.removedlikes))
		for id := range m.removedlikes {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RestaurantMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedlikes {
		edges = append(edges, restaurant.EdgeLikes)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RestaurantMutation) EdgeCleared(name string) bool {
	switch name {
	case restaurant.EdgeLikes:
		return m.clearedlikes
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RestaurantMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Restaurant unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RestaurantMutation) ResetEdge(name string) error {
	switch name {
	case restaurant.EdgeLikes:
		m.ResetLikes()
		return nil
	}
	return fmt.Errorf("unknown Restaurant edge %s", name)
}
