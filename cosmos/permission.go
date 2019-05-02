package cosmos

// Permission model
type Permission struct {
	client       Client
	user         User
	permissionID string
}

// PermissionDefinition model
type PermissionDefinition struct {
	ID             string `json:"id"`
	PermissionMode string `json:"permissionMode,omitempty"`
	Resource       string `json:"resource"`
}

// Permissions model
type Permissions struct {
	client Client
	user   User
}

func newPermission(user User, permissionID string) *Permission {
	user.client.path += "/permissions/" + permissionID
	user.client.rType = "permissions"
	user.client.rLink = user.client.path
	permission := &Permission{
		client:       user.client,
		user:         user,
		permissionID: permissionID,
	}
	return permission
}

func newPermissions(user User) *Permissions {
	user.client.path += "/permissions"
	user.client.rType = "permissions"
	permissions := &Permissions{
		client: user.client,
		user:   user,
	}
	return permissions
}

// Replace permission.
func (u *Permission) Replace(permission *PermissionDefinition, opts ...CallOption) (*PermissionDefinition, error) {
	updatedPermission := &PermissionDefinition{}
	_, err := u.client.replace(permission, &updatedPermission, opts...)
	if err != nil {
		return nil, err
	}
	return updatedPermission, err
}

// Delete permission.
func (u *Permission) Delete(opts ...CallOption) (*Response, error) {
	return u.client.delete(opts...)
}

// Read permission.
func (u *Permission) Read(opts ...CallOption) (*PermissionDefinition, error) {
	permission := &PermissionDefinition{}
	_, err := u.client.read(permission, opts...)
	return permission, err
}

// ReadAll permissions.
func (u *Permissions) ReadAll(opts ...CallOption) ([]PermissionDefinition, error) {
	data := struct {
		Permissions []PermissionDefinition `json:"permissions,omitempty"`
		Count       int                    `json:"_count,omitempty"`
	}{}

	_, err := u.client.read(&data, opts...)
	if err != nil {
		return nil, err
	}
	return data.Permissions, err
}

// Create a new permission
func (u *Permissions) Create(permission *PermissionDefinition, opts ...CallOption) (*PermissionDefinition, error) {
	createdPermission := &PermissionDefinition{}
	_, err := u.client.create(permission, &createdPermission, opts...)
	if err != nil {
		return nil, err
	}
	return createdPermission, err
}
