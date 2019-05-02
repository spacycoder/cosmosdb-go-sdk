package cosmos

import "testing"

func TestCosmosCollection(t *testing.T) {
	client := getDummyClient()
	db := client.Database("dbtest")
	coll := db.Collection("colltest")
	if coll.client.rType != "colls" {
		t.Errorf("%+v", coll.client)
	}

	if coll.client.rLink != "dbs/dbtest/colls/colltest" {
		t.Errorf("%+v", coll.client)
	}

	if coll.client.path != "dbs/dbtest/colls/colltest" {
		t.Errorf("%+v", coll.client)
	}

	colls := db.Collections()
	if coll.client.rType != "colls" {
		t.Errorf("Wrong rType %s", colls.client.rType)
	}
	if colls.client.rLink != "dbs/dbtest" {
		t.Errorf("Wrong rLink %s", colls.client.rLink)
	}

	if colls.client.path != "dbs/dbtest/colls" {
		t.Errorf("Wrong path %s", colls.client.path)
	}
}
