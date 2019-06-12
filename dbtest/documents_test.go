package dbtest

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/spacycoder/cosmosdb-go-sdk/cosmos"
)

func getClient() (*cosmos.Client, error) {
	testDbURL := os.Getenv("TEST_COSMOS_URL")
	return cosmos.New(testDbURL)
}

var dbID = "db-test"
var collID = "coll-test"

type TestDoc struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestCosmos(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatalf("Creating client caused error: %s", err.Error())
	}

	testDatabaseOperations(t, client)
	testCollections(t, client)
	testDocuments(t, client)
	testStoredProcedure(t, client)
	testUDF(t, client)
	testTrigger(t, client)
	testUser(t, client)
	testOffers(t, client)
	cleanup(t, client)
}

func cleanup(t *testing.T, client *cosmos.Client) {
	// Delete user
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Database(dbID).User("myUpdatedUser").Delete(ctx)
	if err != nil {
		t.Fatalf("Deleting user caused error: %s", err.Error())
	}

	// Delete collection
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
	coll := client.Database(dbID).Collection(collID)
	_, err = coll.Delete(ctx2)
	if err != nil {
		t.Fatalf("Deleting collection caused error: %s", err.Error())
	}

	// Delete database
	ctx3, cancel3 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel3()
	_, err = client.Database(dbID).Delete(ctx3)
	if err != nil {
		t.Fatalf("Deleting collection caused error: %s", err.Error())
	}
}

func testDatabaseOperations(t *testing.T, client *cosmos.Client) {
	// Create db
	ctx := context.Background()
	newDbRef, err := client.Databases().Create(ctx, dbID)
	if err != nil {
		t.Fatalf("Creating database caused error: %s", err.Error())
	}
	if newDbRef.ID != dbID {
		t.Fatalf("Wrong ID: %s should be: test", newDbRef.ID)
	}

	// Read db
	db := client.Database(dbID)
	testDb, err := db.Read(ctx)
	if err != nil {
		t.Fatalf("Reading database caused error: %s", err.Error())
	}
	if testDb.ID != dbID {
		t.Fatalf("Wrong ID: %s should be: %s", testDb.ID, dbID)
	}

	// List databases
	dbs := client.Databases()
	_, err = dbs.ReadAll(ctx)
	if err != nil {
		t.Fatalf("Listing databases caused error: %s", err.Error())
	}

}

func testCollections(t *testing.T, client *cosmos.Client) {
	db := client.Database(dbID)
	colls := db.Collections()
	ctx := context.Background()
	newCollDef := &cosmos.CollectionDefinition{
		IndexingPolicy: cosmos.IndexingPolicy{IndexingMode: "consistent"},
		Resource:       cosmos.Resource{ID: collID},
		PartitionKey:   cosmos.PartitionKeyDefinition{Kind: "hash", Paths: []string{"/name"}}}
	// Create collection
	_, err := colls.Create(ctx, newCollDef)
	if err != nil {
		t.Fatalf("Creating collection caused error: %s", err.Error())
	}

	coll := db.Collection(collID)
	// Read collection
	collDef, err := coll.Read(ctx)
	if err != nil {
		t.Fatalf("Reading collection caused error: %s", err.Error())
	}
	if collDef.ID != collID {
		t.Fatalf("Wrong ID: %s should be: %s", collDef.ID, collID)
	}
	if collDef.IndexingPolicy.IndexingMode != "consistent" {
		t.Fatalf("Wrong IndexingMode: %s should be: %s", collDef.IndexingPolicy.IndexingMode, "consistent")
	}

	/*  @TODO: Replace collection causes ERROR!!!! UNKOWN REASON
	newPolicy := &cosmos.IndexingPolicy{
		Automatic:    true,
		IndexingMode: "lazy",
		IncludedPaths: []cosmos.IndexingPolicyPath{
			{Path: "/*", Indexes: []cosmos.PolicyIndex{{DataType: "Number", Precision: -1, Kind: "Range"}, {DataType: "String", Precision: -1, Kind: "Hash"}}},
		},
		ExcludedPaths: []cosmos.IndexingPolicyPath{},
	}

		var updatedCollection cosmos.CollectionDefinition
		_, err = coll.Replace(newPolicy, &updatedCollection)
		if err != nil {
			t.Fatalf("Listing collections caused error: %s", err.Error())
		}
		if updatedCollection.IndexingPolicy.IndexingMode != "lazy" {
			t.Fatalf("Wrong IndexingMode: %s should be: %s", collDef.IndexingPolicy.IndexingMode, "lazy")
		}
	*/

	// List collections
	collDefs, err := colls.ReadAll(ctx)
	if err != nil {
		t.Fatalf("Listing collections caused error: %s", err.Error())
	}

	if len(*collDefs) != 1 {
		t.Fatalf("Number of collections are wrong: %d, %+v", len(*collDefs), collDefs)
	}
}

func testDocuments(t *testing.T, client *cosmos.Client) {
	db := client.Database(dbID)
	coll := db.Collection(collID)
	docs := coll.Documents()
	ctx := context.Background()

	user1 := &TestDoc{
		ID:   "user1",
		Name: "Lars",
		Age:  150,
	}
	// Create document
	_, err := docs.Create(ctx, user1, cosmos.PartitionKey(user1.Name))
	if err != nil {
		t.Fatalf("Creating doc caused error: %s", err.Error())
	}

	retUser := &TestDoc{}
	_, err = coll.Document("user1").Read(ctx, retUser, cosmos.PartitionKey("Lars"))
	if err != nil {
		t.Fatalf("Read doc caused error: %s", err.Error())
	}
	if retUser.Name != user1.Name {
		t.Fatalf("Wrong name: %s", retUser.Name)
	}

	// Replace document
	newUser := &TestDoc{
		ID:   "user1",
		Name: "Lars",
		Age:  20,
	}
	var updatedUser TestDoc
	_, err = coll.Document(retUser.ID).Replace(ctx, newUser, &updatedUser, cosmos.PartitionKey(retUser.Name))
	if err != nil {
		t.Fatalf("Replace doc caused error: %s", err.Error())
	}
	if updatedUser.Age != newUser.Age {
		t.Fatalf("User age not updated: %d", updatedUser.Age)
	}
	user2 := &TestDoc{
		ID:   "user2",
		Name: "Trygve",
		Age:  150,
	}

	// Create doc 2
	_, err = docs.Create(ctx, user2, cosmos.PartitionKey(user2.Name))
	if err != nil {
		t.Fatalf("Creating doc caused error: %s", err.Error())
	}

	// List all docs
	users := []TestDoc{}
	_, err = docs.ReadAll(ctx, &users)
	if err != nil {
		t.Fatalf("Listing docs caused error: %s", err.Error())
	}
	if len(users) != 2 {
		t.Fatalf("Should be 2 but is: %d", len(users))
	}

	// Query docs
	qUsers := []TestDoc{}
	query := cosmos.Q("SELECT * FROM root WHERE root.name = @NAME", cosmos.P{Name: "@NAME", Value: user1.Name})
	_, err = docs.Query(ctx, query, &qUsers, cosmos.CrossPartition())
	if err != nil {
		t.Fatalf("Querying docs caused error: %s", err.Error())
	}
	if len(qUsers) != 1 {
		t.Fatalf("Should be 1 but is: %d", len(qUsers))
	}

	_, err = coll.Document("user1").Delete(ctx, cosmos.PartitionKey(user1.Name))
	if err != nil {
		t.Fatalf("Deleting user1 caused error: %s", err.Error())
	}
	_, err = coll.Document("user2").Delete(ctx, cosmos.PartitionKey(user2.Name))
	if err != nil {
		t.Fatalf("Deleting user2 caused error: %s", err.Error())
	}
}

func testStoredProcedure(t *testing.T, client *cosmos.Client) {
	coll := client.Database(dbID).Collection(collID)

	// Create Stored Procedure
	ctx := context.Background()
	spDef := &cosmos.StoredProcedureDefinition{Resource: cosmos.Resource{ID: "mySP"}, Body: "function () {\r\n    var context = getContext();\r\n    var response = context.getResponse();\r\n\r\n    response.setBody(\"Hello, World\");\r\n}"}
	createdSP, err := coll.StoredProcedures().Create(ctx, spDef)
	if err != nil {
		t.Fatalf("Creating stored procedure caused error: %s", err.Error())
	}
	if createdSP.Body != spDef.Body {
		t.Fatalf("Invalid stored procedure body: %s", createdSP.Body)
	}

	// Execute Stored Procedure
	var res string
	_, err = coll.StoredProcedure("mySP").Execute(ctx, "", &res)
	if err != nil {
		t.Fatalf("Executing stored procedure caused error: %s", err.Error())
	}
	if res != "Hello, World" {
		t.Fatalf("Response is: %s", res)
	}

	// Replace Stored Procedure
	newSpDef := &cosmos.StoredProcedureDefinition{Resource: cosmos.Resource{ID: "mySP"}, Body: "function (greet, someone) {\r\n    var context = getContext();\r\n    var response = context.getResponse();\r\n\r\n    response.setBody(greet + \", \"+ someone);\r\n}"}
	_, err = coll.StoredProcedure("mySP").Replace(ctx, newSpDef)
	if err != nil {
		t.Fatalf("Replacing stored procedure caused error: %s", err.Error())
	}

	var res2 string
	_, err = coll.StoredProcedure("mySP").Execute(ctx, []string{"Hello", "Cosmos"}, &res2)
	if err != nil {
		t.Fatalf("Executing stored procedure caused error: %s", err.Error())
	}
	if res2 != "Hello, Cosmos" {
		t.Fatalf("Response is: %s", res2)
	}

	// List all stored procedures
	sprocs, err := coll.StoredProcedures().ReadAll(ctx)
	if err != nil {
		t.Fatalf("Listing stored procedure caused error: %s", err.Error())
	}
	if len(sprocs) != 1 {
		t.Fatalf("Invalid length: %d", len(sprocs))
	}

	// Delete stored procedure
	_, err = coll.StoredProcedure("mySP").Delete(ctx)
	if err != nil {
		t.Fatalf("Deleting stored procedure caused error: %s", err.Error())
	}
}

func testUDF(t *testing.T, client *cosmos.Client) {
	coll := client.Database(dbID).Collection(collID)
	udfDef := &cosmos.UDFDefinition{
		Body:     "function tax(income) {\r\n    if(income == undefined) \r\n        throw 'no input';\r\n    if (income < 1000) \r\n        return income * 0.1;\r\n    else if (income < 10000) \r\n        return income * 0.2;\r\n    else\r\n        return income * 0.4;\r\n}",
		Resource: cosmos.Resource{ID: "myUDF"},
	}
	ctx := context.Background()
	// Create UDF
	createdUDF, err := coll.UDFs().Create(ctx, udfDef)
	if err != nil {
		t.Fatalf("Creating UDF caused error: %s", err.Error())
	}
	if createdUDF.Body != udfDef.Body {
		t.Fatalf("Invalid UDF body: %s", createdUDF.Body)
	}

	// Replace UDF
	newUDF := &cosmos.UDFDefinition{
		Body:     "function tax(income) {\r\n    if(income == undefined) \r\n        throw 'no input';\r\n    if (income < 2000) \r\n        return income * 0.1;\r\n    else if (income < 10000) \r\n        return income * 0.2;\r\n    else\r\n        return income * 0.4;\r\n}",
		Resource: cosmos.Resource{ID: "myUDF"},
	}
	updatedUDF, err := coll.UDF("myUDF").Replace(ctx, newUDF)
	if err != nil {
		t.Fatalf("Replacing UDF caused error: %s", err.Error())
	}
	if updatedUDF.Body != newUDF.Body {
		t.Fatalf("Invalid UDF body: %s", createdUDF.Body)
	}

	// List UDFs
	udfs, err := coll.UDFs().ReadAll(ctx)
	if err != nil {
		t.Fatalf("Listing UDFs caused error: %s", err.Error())
	}
	if len(udfs) != 1 {
		t.Fatalf("Invalid length: %d", len(udfs))
	}

	// Delete UDF
	_, err = coll.UDF("myUDF").Delete(ctx)
	if err != nil {
		t.Fatalf("Deleting UDFs caused error: %s", err.Error())
	}
}

func testTrigger(t *testing.T, client *cosmos.Client) {
	coll := client.Database(dbID).Collection(collID)
	triggerDef := &cosmos.TriggerDefintion{
		Body:             "function updateMetadata() {\r\n  var context = getContext();\r\nvar collection = context.getCollection();\r\nvar response = context.getResponse();\r\nvar createdDocument = response.getBody();\r\n\r\n// query for metadata document\r\nvar filterQuery = 'SELECT * FROM root r WHERE r.id = \"_metadata\"';\r\nvar accept = collection.queryDocuments(collection.getSelfLink(), filterQuery,\r\n  updateMetadataCallback);\r\n    if(!accept) throw \"Unable to update metadata, abort\";\r\n\r\nfunction updateMetadataCallback(err, documents, responseOptions) {\r\n  if(err) throw new Error(\"Error\" + err.message);\r\n   if(documents.length != 1) throw 'Unable to find metadata document';\r\n   var metadataDocument = documents[0];\r\n\r\n   // update metadata\r\n   metadataDocument.createdDocuments += 1;\r\n   metadataDocument.createdNames += \" \" + createdDocument.id;\r\nvar accept = collection.replaceDocument(metadataDocument._self,\r\n    metadataDocument, function(err, docReplaced) {\r\n       if(err) throw \"Unable to update metadata, abort\";\r\n    });\r\nif(!accept) throw \"Unable to update metadata, abort\";\r\nreturn;\r\n    }",
		Resource:         cosmos.Resource{ID: "myTrigger"},
		TriggerOperation: "All",
		TriggerType:      "Post",
	}

	// Create Trigger
	ctx := context.Background()
	_, err := coll.Triggers().Create(ctx, triggerDef)
	if err != nil {
		t.Fatalf("Creating trigger caused error: %s", err.Error())
	}

	// Replace Trigger
	newTriggerDef := &cosmos.TriggerDefintion{
		Body:             "function updateMetadata() {\r\n  var context = getContext();\r\nvar collection = context.getCollection();\r\nvar response = context.getResponse();\r\nvar createdDocument = response.getBody();\r\n\r\n// query for metadata document\r\nvar filterQuery = 'SELECT * FROM root r WHERE r.id = \"_metadata\"';\r\nvar accept = collection.queryDocuments(collection.getSelfLink(), filterQuery,\r\n  updateMetadataCallback);\r\n    if(!accept) throw \"Unable to update metadata, exit\";\r\n\r\nfunction updateMetadataCallback(err, documents, responseOptions) {\r\n  if(err) throw new Error(\"Error\" + err.message);\r\n   if(documents.length != 1) throw 'Unable to find metadata document';\r\n   var metadataDocument = documents[0];\r\n\r\n   // update metadata\r\n   metadataDocument.createdDocuments += 1;\r\n   metadataDocument.createdNames += \" \" + createdDocument.id;\r\nvar accept = collection.replaceDocument(metadataDocument._self,\r\n    metadataDocument, function(err, docReplaced) {\r\n       if(err) throw \"Unable to update metadata, abort\";\r\n    });\r\nif(!accept) throw \"Unable to update metadata, abort\";\r\nreturn;\r\n    }",
		Resource:         cosmos.Resource{ID: "myTrigger"},
		TriggerOperation: "All",
		TriggerType:      "Post",
	}
	updatedTriggerDef, err := coll.Trigger("myTrigger").Replace(ctx, newTriggerDef)
	if err != nil {
		t.Fatalf("Replacing trigger caused error: %s", err.Error())
	}
	if updatedTriggerDef.Body != newTriggerDef.Body {
		t.Fatalf("Invalid trigger body: %s", updatedTriggerDef.Body)
	}

	// List triggers
	triggers, err := coll.Triggers().ReadAll(ctx)
	if err != nil {
		t.Fatalf("Listing triggers caused error: %s", err.Error())
	}
	if len(triggers) != 1 {
		t.Fatalf("Invalid length: %d", len(triggers))
	}

	// Delete trigger
	_, err = coll.Trigger("myTrigger").Delete(ctx)
	if err != nil {
		t.Fatalf("Deleting UDFs caused error: %s", err.Error())
	}
}

func testUser(t *testing.T, client *cosmos.Client) {
	db := client.Database(dbID)
	myUser := &cosmos.UserDefinition{
		Resource: cosmos.Resource{ID: "myUser"},
	}
	ctx := context.Background()
	// Create user
	_, err := db.Users().Create(ctx, myUser)
	if err != nil {
		t.Fatalf("Creating user caused error: %s", err.Error())
	}

	// Read user
	user, err := db.User("myUser").Read(ctx)
	if err != nil {
		t.Fatalf("Reading user caused error: %s", err.Error())
	}
	if user.ID != myUser.ID {
		t.Fatalf("Wrong userID: %s", user.ID)
	}

	newUser := &cosmos.UserDefinition{
		Resource: cosmos.Resource{ID: "myUpdatedUser"},
	}
	// Replace user
	updatedUser, err := db.User("myUser").Replace(ctx, newUser)
	if err != nil {
		t.Fatalf("Replacing user caused error: %s", err.Error())
	}
	if updatedUser.ID != newUser.ID {
		t.Fatalf("Wrong userID: %s", updatedUser.ID)
	}

	// List users
	allUsers, err := db.Users().ReadAll(ctx)
	if err != nil {
		t.Fatalf("Listing user caused error: %s", err.Error())
	}
	if len(allUsers) != 1 {
		t.Fatalf("Wrong amount of users: %d", len(allUsers))
	}
}

func testPermissions(t *testing.T, client *cosmos.Client) {
	db := client.Database(dbID)
	myUser := db.User("myUpdatedUser")
	ctx := context.Background()
	myPermission := &cosmos.PermissionDefinition{ID: "test_permission", PermissionMode: "All", Resource: "letsSEEEEE"}
	// Create permission
	_, err := myUser.Permissions().Create(ctx, myPermission)
	if err != nil {
		t.Fatalf("Creating permission caused error: %s", err.Error())
	}

	newPermission := &cosmos.PermissionDefinition{ID: "test_permission2", PermissionMode: "All", Resource: "letsSEEEEE"}
	_, err = myUser.Permission("test_permission").Replace(ctx, newPermission)
	if err != nil {
		t.Fatalf("Replacing permission caused error: %s", err.Error())
	}

	perm, err := myUser.Permission("test_permission2").Read(ctx)
	if err != nil {
		t.Fatalf("Reading permission caused error: %s", err.Error())
	}
	if perm.PermissionMode != newPermission.PermissionMode {
		t.Fatalf("Wrong permission mode: %s", perm.PermissionMode)
	}

	allPermissions, err := myUser.Permissions().ReadAll(ctx)
	if err != nil {
		t.Fatalf("Listing permission caused error: %s", err.Error())
	}
	if len(allPermissions) != 1 {
		t.Fatalf("Wrong amount of permissions: %d", len(allPermissions))
	}

	_, err = myUser.Permission("test_permission2").Delete(ctx)
	if err != nil {
		t.Fatalf("Deleting permission caused error: %s", err.Error())
	}
}

func testOffers(t *testing.T, client *cosmos.Client) {
	ctx := context.Background()
	coll, err := client.Database(dbID).Collection(collID).Read(ctx)
	if err != nil {
		t.Fatalf("Reading collection caused error: %s", err.Error())
	}

	offers, err := client.Offers().ReadAll(ctx)
	if err != nil {
		t.Fatalf("Listing offers caused error: %s", err.Error())
	}

	/*
		@TODO: reading offer causes error: 401, Unauthorized, The input authorization token can't serve the request.
		unkown reason
		offer, err := client.Offer(offers[0].ID).Read()
		if err != nil {
			t.Fatalf("Reading offer caused error: %s", err.Error())
		}
		if offer.OfferVersion != offers[0].OfferVersion {
			t.Fatalf("Invalid offer version: %s", offer.OfferVersion)
		}
	*/

	newOffer := &cosmos.OfferDefinition{
		Content: struct {
			OfferThroughput int `json:"offerThroughput"`
		}{OfferThroughput: 400},
		OfferVersion:    "V2",
		OfferType:       "Invalid",
		OfferResource:   coll.Self,
		OfferResourceID: coll.Rid,
		Resource:        cosmos.Resource{ID: offers[0].ID, Rid: offers[0].Rid},
	}
	client.Offer(offers[0].ID).Replace(ctx, newOffer)

	offerQuery := cosmos.Q("SELECT * FROM root")
	queryOffers, err := client.Offers().Query(ctx, offerQuery)
	if err != nil {
		t.Fatalf("Querying offers caused error: %s", err.Error())
	}
	if len(queryOffers) != len(offers) {
		t.Fatalf("Invalid count: %d", len(queryOffers))
	}
}
