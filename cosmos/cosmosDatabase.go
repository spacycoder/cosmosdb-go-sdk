package cosmos

// Database performs operations on a single database
type Database struct {
	client Client
	dbID   string
}

// Databases performs operations on databases
type Databases struct {
	client *Client
}

// User operations
func (d Database) User(id string) *User {
	return newUser(d, id)
}

// Users operations
func (d Database) Users() *Users {
	return newUsers(d)
}

func newDatabase(client Client, dbID string) *Database {
	client.path = "dbs/" + dbID
	client.rType = "dbs"
	client.rLink = "dbs/" + dbID
	db := &Database{
		client: client,
		dbID:   dbID,
	}
	return db
}

func newDatabases(c *Client) *Databases {
	c.path = "dbs"
	c.rType = "dbs"
	c.rLink = ""

	dbs := &Databases{
		client: c,
	}
	return dbs
}

// Collection gets a handle for a given collection in a database
func (d Database) Collection(collID string) *Collection {
	return newCollection(d, collID)
}

// Collections gets a handle for doing operations on all collections in a database
func (d Database) Collections() *Collections {
	return newCollections(d)
}

// Read database
func (d *Database) Read() (*DatabaseDefinition, error) {
	ret := &DatabaseDefinition{}
	_, err := d.client.read(ret)
	return ret, err
}

// Delete database
func (d *Database) Delete() (*Response, error) {
	return d.client.delete()
}

// Create a new database
func (d *Databases) Create(dbID string, opts ...CallOption) (*DatabaseDefinition, error) {
	dbDef := &DatabaseDefinition{}
	var body struct {
		ID string `json:"id"`
	}
	body.ID = dbID

	_, err := d.client.create(body, dbDef, opts...)
	if err != nil {
		return nil, err
	}
	return dbDef, err
}

// ReadAll databases
func (d *Databases) ReadAll(opts ...CallOption) (*DatabaseDefinitions, error) {
	data := struct {
		Databases DatabaseDefinitions `json:"Databases,omitempty"`
		Count     int                 `json:"_count,omitempty"`
	}{}
	_, err := d.client.read(&data, opts...)
	if err != nil {
		return nil, err
	}
	return &data.Databases, err
}

// Query databases
func (d *Databases) Query(query *SqlQuerySpec, opts ...CallOption) (*DatabaseDefinitions, error) {
	data := struct {
		Databases DatabaseDefinitions `json:"Databases,omitempty"`
		Count     int                 `json:"_count,omitempty"`
	}{}

	_, err := d.client.query(query, &data, opts...)
	if err != nil {
		return nil, err
	}
	return &data.Databases, err
}
