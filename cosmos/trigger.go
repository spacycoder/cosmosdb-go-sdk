package cosmos

import "context"

type Trigger struct {
	client    Client
	coll      Collection
	triggerID string
}

type Triggers struct {
	client Client
	coll   Collection
}

func newTrigger(coll Collection, triggerID string) *Trigger {
	coll.client.path += "/triggers/" + triggerID
	coll.client.rType = "triggers"
	coll.client.rLink = coll.client.path
	trigger := &Trigger{
		client:    coll.client,
		coll:      coll,
		triggerID: triggerID,
	}

	return trigger
}

func newTriggers(coll Collection) *Triggers {
	coll.client.path += "/triggers"
	coll.client.rType = "triggers"
	triggers := &Triggers{
		client: coll.client,
		coll:   coll,
	}

	return triggers
}

func (s *Triggers) Create(ctx context.Context, newTrigger *TriggerDefintion, opts ...CallOption) (*TriggerDefintion, error) {
	createdTrigger := &TriggerDefintion{}
	_, err := s.client.create(ctx, newTrigger, &createdTrigger, opts...)
	if err != nil {
		return nil, err
	}

	return createdTrigger, err
}

func (s *Trigger) Replace(ctx context.Context, trigger *TriggerDefintion, opts ...CallOption) (*TriggerDefintion, error) {
	updatedTrigger := &TriggerDefintion{}
	_, err := s.client.replace(ctx, trigger, &updatedTrigger, opts...)
	if err != nil {
		return nil, err
	}

	return updatedTrigger, err
}

func (s *Triggers) ReadAll(ctx context.Context, opts ...CallOption) ([]TriggerDefintion, error) {
	data := struct {
		Triggers []TriggerDefintion `json:"triggers,omitempty"`
		Count    int                `json:"_count,omitempty"`
	}{}

	_, err := s.client.read(ctx, &data, opts...)
	if err != nil {
		return nil, err
	}
	return data.Triggers, err
}

func (s *Trigger) Delete(ctx context.Context, opts ...CallOption) (*Response, error) {
	return s.client.delete(ctx, opts...)
}

// Execute stored procedure
func (s *Trigger) Execute(ctx context.Context, params, body interface{}, opts ...CallOption) (*Response, error) {
	return s.client.execute(ctx, params, &body, opts...)
}
