package cosmos

// Collection performs operations on a given collection.
type Collection struct {
	client Client
	db     Database
	collID string
}

// Collections struct handles all operations involving mutiple collections
type Collections struct {
	client Client
	db     Database
}

// Document defines possible operations on a single document. e.g. Read, Delete, Replace
func (c Collection) Document(docID string) *Document {
	return newDocument(c, docID)
}

// Documents defines possible operations on multiple documents. e.g. ReadAll, Query
func (c Collection) Documents() *Documents {
	return newDocuments(c)
}

// UDF defines operations on a user defined function
func (c Collection) UDF(id string) *UDF {
	return newUDF(c, id)
}

// UDFs defines operations on multiple user defined functions
func (c Collection) UDFs() *UDFs {
	return newUDFs(c)
}

// StoredProcedure defines operations on a single stored procedure
func (c Collection) StoredProcedure(id string) *StoredProcedure {
	return newStoredProcedure(c, id)
}

// StoredProcedures defines operations on multiple stored procedures
func (c Collection) StoredProcedures() *StoredProcedures {
	return newStoredProcedures(c)
}

// Trigger defines operations on a single trigger
func (c Collection) Trigger(id string) *Trigger {
	return newTrigger(c, id)
}

// Triggers defines operations on multiple triggers
func (c Collection) Triggers() *Triggers {
	return newTriggers(c)
}

func newCollection(db Database, collID string) *Collection {
	db.client.path += "/colls/" + collID
	db.client.rType = "colls"
	db.client.rLink = db.client.path
	coll := &Collection{
		client: db.client,
		db:     db,
		collID: collID,
	}

	return coll
}

func newCollections(db Database) *Collections {
	db.client.path += "/colls"
	db.client.rType = "colls"
	coll := &Collections{
		client: db.client,
		db:     db,
	}

	return coll
}

// Create new collection
func (c *Collections) Create(newColl *CollectionDefinition) (*CollectionDefinition, error) {
	respColl := &CollectionDefinition{}
	_, err := c.client.create(newColl, respColl)
	if err != nil {
		return nil, err
	}

	return respColl, err
}

// ReadAll returns all collections in a database.
func (c *Collections) ReadAll() (*CollectionDefinitions, error) {
	data := struct {
		Collections CollectionDefinitions `json:"DocumentCollections,omitempty"`
		Count       int                   `json:"_count,omitempty"`
	}{}
	_, err := c.client.read(&data)
	return &data.Collections, err
}

// Read returns one collection
func (c *Collection) Read() (*CollectionDefinition, error) {
	coll := &CollectionDefinition{}
	_, err := c.client.read(coll)
	return coll, err
}

// Delete collection
func (c *Collection) Delete() (*Response, error) {
	return c.client.delete()
}

// Replace collection
func (c *Collection) Replace(i *IndexingPolicy, ret interface{}, opts ...CallOption) (*Response, error) {
	body := struct {
		ID             string          `json:"id"`
		IndexingPolicy *IndexingPolicy `json:"indexingPolicy"`
	}{c.collID, i}
	return c.client.replace(&body, ret, opts...)
}
