package cosmos

import "testing"

func TestUDF(t *testing.T) {
	client := getDummyClient()
	coll := client.Database("dbtest").Collection("colltest")
	udf := coll.UDF("myudf")

	if udf.client.rType != "udfs" {
		t.Errorf("%+v", udf.client)
	}

	if udf.client.rLink != "dbs/dbtest/colls/colltest/udfs/myudf" {
		t.Errorf("%+v", udf.client)
	}

	if udf.client.path != "dbs/dbtest/colls/colltest/udfs/myudf" {
		t.Errorf("%+v", udf.client)
	}

	udfs := coll.UDFs()
	if udfs.client.rType != "udfs" {
		t.Errorf("Wrong rType %s", udfs.client.rType)
	}

	if udfs.client.rLink != "dbs/dbtest/colls/colltest" {
		t.Errorf("Wrong rLink %s", udfs.client.rLink)
	}

	if udfs.client.path != "dbs/dbtest/colls/colltest/udfs" {
		t.Errorf("Wrong path %s", udfs.client.path)
	}
}
