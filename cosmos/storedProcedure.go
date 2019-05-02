package cosmos

type StoredProcedure struct {
	client            Client
	coll              Collection
	storedProcedureID string
}

type StoredProcedures struct {
	client Client
	coll   Collection
}

func newStoredProcedure(coll Collection, storedProcedureID string) *StoredProcedure {
	coll.client.path += "/sprocs/" + storedProcedureID
	coll.client.rType = "sprocs"
	coll.client.rLink = coll.client.path
	udf := &StoredProcedure{
		client:            coll.client,
		coll:              coll,
		storedProcedureID: storedProcedureID,
	}

	return udf
}

func newStoredProcedures(coll Collection) *StoredProcedures {
	coll.client.path += "/sprocs"
	coll.client.rType = "sprocs"
	udfs := &StoredProcedures{
		client: coll.client,
		coll:   coll,
	}

	return udfs
}

func (s *StoredProcedures) Create(newStoredProcedure *StoredProcedureDefinition, opts ...CallOption) (*StoredProcedureDefinition, error) {
	storedProcedureResp := &StoredProcedureDefinition{}
	_, err := s.client.create(newStoredProcedure, &storedProcedureResp, opts...)
	if err != nil {
		return nil, err
	}

	return storedProcedureResp, err
}

func (s *StoredProcedure) Replace(newStoredProcedure *StoredProcedureDefinition, opts ...CallOption) (*StoredProcedureDefinition, error) {
	storedProcedureResp := &StoredProcedureDefinition{}
	_, err := s.client.replace(newStoredProcedure, &storedProcedureResp, opts...)
	if err != nil {
		return nil, err
	}

	return storedProcedureResp, err
}

func (s *StoredProcedures) ReadAll(opts ...CallOption) ([]StoredProcedureDefinition, error) {
	data := struct {
		StoredProcedures []StoredProcedureDefinition `json:"StoredProcedures,omitempty"`
		Count            int                         `json:"_count,omitempty"`
	}{}

	_, err := s.client.read(&data, opts...)
	if err != nil {
		return nil, err
	}

	return data.StoredProcedures, err
}

func (s *StoredProcedure) Delete(opts ...CallOption) (*Response, error) {
	return s.client.delete(opts...)
}

// Execute stored procedure
func (s *StoredProcedure) Execute(params, body interface{}, opts ...CallOption) (*Response, error) {
	return s.client.execute(params, &body, opts...)
}
