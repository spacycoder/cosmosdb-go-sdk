package cosmos

import "testing"

func TestPermission(t *testing.T) {
	client := getDummyClient()
	db := client.Database("dbtest")
	user := db.User("myUser")
	perm := user.Permission("myPerm")

	if perm.client.rType != "permissions" {
		t.Errorf("%+v", perm.client)
	}

	if perm.client.rLink != "dbs/dbtest/users/myUser/permissions/myPerm" {
		t.Errorf("%+v", perm.client)
	}

	if perm.client.path != "dbs/dbtest/users/myUser/permissions/myPerm" {
		t.Errorf("%+v", perm.client)
	}

	perms := user.Permissions()
	if perms.client.rType != "permissions" {
		t.Errorf("Wrong rType %s", perms.client.rType)
	}

	if perms.client.rLink != "dbs/dbtest/users/myUser" {
		t.Errorf("Wrong rLink %s", perms.client.rLink)
	}

	if perms.client.path != "dbs/dbtest/users/myUser/permissions" {
		t.Errorf("Wrong path, expexted: %s, got: %s", "dbs/dbtest/users/myUser/permissions", perms.client.path)
	}
}
