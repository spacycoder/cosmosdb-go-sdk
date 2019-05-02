package cosmos

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

func (s *Triggers) Create(newTrigger *TriggerDefintion, opts ...CallOption) (*TriggerDefintion, error) {
	createdTrigger := &TriggerDefintion{}
	_, err := s.client.create(newTrigger, &createdTrigger, opts...)
	if err != nil {
		return nil, err
	}

	return createdTrigger, err
}

func (s *Trigger) Replace(trigger *TriggerDefintion, opts ...CallOption) (*TriggerDefintion, error) {
	updatedTrigger := &TriggerDefintion{}
	_, err := s.client.replace(trigger, &updatedTrigger, opts...)
	if err != nil {
		return nil, err
	}

	return updatedTrigger, err
}

func (s *Triggers) ReadAll(opts ...CallOption) ([]TriggerDefintion, error) {
	data := struct {
		Triggers []TriggerDefintion `json:"triggers,omitempty"`
		Count    int                `json:"_count,omitempty"`
	}{}

	_, err := s.client.read(&data, opts...)
	if err != nil {
		return nil, err
	}
	return data.Triggers, err
}

func (s *Trigger) Delete(opts ...CallOption) (*Response, error) {
	return s.client.delete(opts...)
}

// Execute stored procedure
func (s *Trigger) Execute(params, body interface{}, opts ...CallOption) (*Response, error) {
	return s.client.execute(params, &body, opts...)
}
