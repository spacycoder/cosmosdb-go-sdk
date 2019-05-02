package cosmos

import "testing"

func TestCosmosDatabase(t *testing.T) {
	client := getDummyClient()
	db := client.Database("dbtest")
	if db.client.rType != "dbs" {
		t.Errorf("%+v", db.client)
	}

	if db.client.rLink != "dbs/dbtest" {
		t.Errorf("%+v", db.client)
	}

	if db.client.path != "dbs/dbtest" {
		t.Errorf("%+v", db.client)
	}

	dbs := client.Databases()
	if dbs.client.rType != "dbs" {
		t.Errorf("Wrong rType %s", dbs.client.rType)
	}
	if dbs.client.rLink != "" {
		t.Errorf("Wrong rLink %s", dbs.client.rLink)
	}

	if dbs.client.path != "dbs" {
		t.Errorf("Wrong path %s", dbs.client.path)
	}
}
