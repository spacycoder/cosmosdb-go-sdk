package cosmos

import "testing"

func TestTrigger(t *testing.T) {
	client := getDummyClient()
	coll := client.Database("dbtest").Collection("colltest")
	trigger := coll.Trigger("myTrigger")

	if trigger.client.rType != "triggers" {
		t.Errorf("%+v", trigger.client)
	}

	if trigger.client.rLink != "dbs/dbtest/colls/colltest/triggers/myTrigger" {
		t.Errorf("%+v", trigger.client)
	}

	if trigger.client.path != "dbs/dbtest/colls/colltest/triggers/myTrigger" {
		t.Errorf("%+v", trigger.client)
	}

	triggers := coll.Triggers()
	if triggers.client.rType != "triggers" {
		t.Errorf("Wrong rType %s", triggers.client.rType)
	}

	if triggers.client.rLink != "dbs/dbtest/colls/colltest" {
		t.Errorf("Wrong rLink %s", triggers.client.rLink)
	}

	if triggers.client.path != "dbs/dbtest/colls/colltest/triggers" {
		t.Errorf("Wrong path %s", triggers.client.path)
	}
}
