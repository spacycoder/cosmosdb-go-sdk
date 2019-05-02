package cosmos

import "testing"

func TestStoredProcedure(t *testing.T) {
	client := getDummyClient()
	coll := client.Database("dbtest").Collection("colltest")
	sp := coll.StoredProcedure("myStoredProcedure")

	if sp.client.rType != "sprocs" {
		t.Errorf("%+v", sp.client)
	}

	if sp.client.rLink != "dbs/dbtest/colls/colltest/sprocs/myStoredProcedure" {
		t.Errorf("%+v", sp.client)
	}

	if sp.client.path != "dbs/dbtest/colls/colltest/sprocs/myStoredProcedure" {
		t.Errorf("%+v", sp.client)
	}

	sps := coll.StoredProcedures()
	if sps.client.rType != "sprocs" {
		t.Errorf("Wrong rType %s", sps.client.rType)
	}

	if sps.client.rLink != "dbs/dbtest/colls/colltest" {
		t.Errorf("Wrong rLink %s", sps.client.rLink)
	}

	if sps.client.path != "dbs/dbtest/colls/colltest/sprocs" {
		t.Errorf("Wrong path %s", sps.client.path)
	}
}
