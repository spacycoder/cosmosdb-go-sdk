package cosmos

type User struct {
	client Client
	db     Database
	userID string
}

type UserDefinition struct {
	Resource
	_persmissions string `json:"_persmissions,omitempty"`
}

type Users struct {
	client Client
	db     Database
}

func (u User) Permission(id string) *Permission {
	return newPermission(u, id)
}

func (u User) Permissions() *Permissions {
	return newPermissions(u)
}

func newUser(db Database, userID string) *User {
	db.client.path += "/users/" + userID
	db.client.rType = "users"
	db.client.rLink = db.client.path
	user := &User{
		client: db.client,
		db:     db,
		userID: userID,
	}

	return user
}

func newUsers(db Database) *Users {
	db.client.path += "/users"
	db.client.rType = "users"
	users := &Users{
		client: db.client,
		db:     db,
	}

	return users
}

// Create a new user
func (u *Users) Create(user *UserDefinition, opts ...CallOption) (*UserDefinition, error) {
	createdUser := &UserDefinition{}
	_, err := u.client.create(user, &createdUser, opts...)
	if err != nil {
		return nil, err
	}

	return createdUser, err
}

// Replace an existing user with a new one.
func (u *User) Replace(user *UserDefinition, opts ...CallOption) (*UserDefinition, error) {
	updatedUser := &UserDefinition{}
	_, err := u.client.replace(user, &updatedUser, opts...)
	if err != nil {
		return nil, err
	}

	return updatedUser, err
}

// ReadAll users in a collection
func (u *Users) ReadAll(opts ...CallOption) ([]UserDefinition, error) {
	data := struct {
		Users []UserDefinition `json:"users,omitempty"`
		Count int              `json:"_count,omitempty"`
	}{}

	_, err := u.client.read(&data, opts...)
	if err != nil {
		return nil, err
	}
	return data.Users, err
}

// Delete existing user
func (u *User) Delete(opts ...CallOption) (*Response, error) {
	return u.client.delete(opts...)
}

// Read a single user from collection
func (u *User) Read(opts ...CallOption) (*UserDefinition, error) {
	user := &UserDefinition{}
	_, err := u.client.read(user, opts...)
	return user, err
}
