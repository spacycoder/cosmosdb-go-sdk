package cosmos

import "testing"

func TestPermission(t *testing.T) {
	client := getDummyClient()
	db := client.Database("dbtest")
	sp := db.User("myUser")

	if sp.client.rType != "users" {
		t.Errorf("%+v", sp.client)
	}

	if sp.client.rLink != "dbs/dbtest/users/myUser" {
		t.Errorf("%+v", sp.client)
	}

	if sp.client.path != "dbs/dbtest/users/myUser" {
		t.Errorf("%+v", sp.client)
	}

	usrs := db.Users()
	if usrs.client.rType != "users" {
		t.Errorf("Wrong rType %s", usrs.client.rType)
	}

	if usrs.client.rLink != "dbs/dbtest" {
		t.Errorf("Wrong rLink %s", usrs.client.rLink)
	}

	if usrs.client.path != "dbs/dbtest/users" {
		t.Errorf("Wrong path %s", usrs.client.path)
	}
}
