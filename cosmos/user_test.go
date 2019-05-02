package cosmos

import "testing"

func TestUser(t *testing.T) {
	client := getDummyClient()
	db := client.Database("dbtest")
	user := db.User("myUser")

	if user.client.rType != "users" {
		t.Errorf("%+v", user.client)
	}

	if user.client.rLink != "dbs/dbtest/users/myUser" {
		t.Errorf("%+v", user.client)
	}

	if user.client.path != "dbs/dbtest/users/myUser" {
		t.Errorf("%+v", user.client)
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
