# CosmosDB GO sdk

- [Installation](#installation)
- [Usage](#usage)
  - [Documents](#Documents)
  - [Collections](#Collection)
  - [Databases](#Databases)
  - [StoredProcedures](#StoredProcedures)
  - [UDFs](#UDFs)
  - [Triggers](#Triggers)
  - [Query Builder](#QueryBuilder)

## Installation

- go get github.com/spacycoder/cosmosdb-go-sdk/cosmos

## Usage

To get started import the `cosmos` package and create a client.

```go
import github.com/SpacyCoder/cosmosdb-go-sdk/cosmos

client, err := cosmos.New("YOUR_CONNECTION_STRING")
```

### Documents

#### Create Document

note: If your don't supply an `id` for the document it will be automatically created for you

```GO
type Person struct {
    cosmos.DocumentDefinition
    Name string `json:"name"`
    Age string `json:"age"`
}

func main() {
    client, err := cosmos.New("YOUR_CONNECTION_STRING")
    newPerson := &Person{
        Name: "Todd",
        Age: 99
    }
    res, err := client.Database("dbID").Collection("CollectionID").Documents(newPerson).Create()
}

```

#### Read Document

```GO
type Person struct {
    cosmos.DocumentDefinition
    Name string `json:"name"`
    Age string `json:"age"`
}

func main() {
    client, err := cosmos.New("YOUR_CONNECTION_STRING")
    var person Person
    res, err := client.Database("dbID").Collection("CollectionID").Document("docID").Read(&person, cosmos.PartitionKey(99))
}
```

#### List Documents

```GO
type Person struct {
    cosmos.DocumentDefinition
    Name string `json:"name"`
    Age string `json:"age"`
}

func main() {
    client, err := cosmos.New("YOUR_CONNECTION_STRING")
    var people []Person
    res, err := client.Database("dbID").Collection("CollectionID").Documents().ReadAll(&people, cosmos.CrossPartition())
}
```

#### Query Documents

```GO
type Person struct {
    cosmos.DocumentDefinition
    Name string `json:"name"`
    Age string `json:"age"`
}

func main() {
    client, err := cosmos.New("YOUR_CONNECTION_STRING")
    db := client.Database("dbID")
    coll := db.Collection("CollectionID")
    var people []Person

    params := []cosmos.P{{Name: "@LENGTH", Value: 180},{Name: "@AGE", Value: 30}}
    query := cosmos.Q("SELECT * FROM root WHERE root.length < @LENGTH AND  root.age > @AGE", params...)
    res, err := coll.Documents().Query(query, &people, cosmos.CrossPartition())
}
```

or with query builder

```GO
import ("github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"
 "github.com/SpacyCoder/cosmosdb-go-sdk/qbuilder")

type Person struct {
    cosmos.DocumentDefinition
    Name string `json:"name"`
    Age string `json:"age"`
}

func main() {
    client, err := cosmos.New("YOUR_CONNECTION_STRING")
    db := client.Database("dbID")
    coll := db.Collection("CollectionID")

    var people []Person
    params := []cosmos.P{{Name: "@LENGTH", Value: 180},{Name: "@AGE", Value: 30}}
    qb := qbuilder.New()
    query := qb.Select("*").From("root").And("root.length < @LENGTH").And("root.age > @AGE").Params(params...).Build()
    res, err := coll.Documents().Query(query, &people, cosmos.CrossPartition())
}
```

#### Delete Document

```GO
import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

func main() {
    client, err := cosmos.New("YOUR_CONNECTION_STRING")
    db := client.Database("dbID")
    coll := db.Collection("CollectionID")

    res, err := coll.Document("docID").Delete(cosmos.PartitionKey("partitionKey"))
}
```

### Collection

#### Create collection

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        db := client.Database("dbID")
	    newCollDef := &cosmos.CollectionDefinition{
	    	IndexingPolicy: cosmos.IndexingPolicy{IndexingMode: "consistent"},
	    	Resource:       cosmos.Resource{ID: collID},
            PartitionKey:   cosmos.PartitionKeyDefinition{Kind: "hash", Paths: []string{"/name"}
            }
        }
        _, err := db.Collections().Create(newCollDef)
    }
```

#### Read Collection

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        db := client.Database("dbID")

        collDef, err := db.Collection("collID").Read()
    }
```

#### List Collections

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        db := client.Database("dbID")

        collDefs, err :=  db.Collections().ReadAll()
    }
```

#### Delete Collection

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        db := client.Database("dbID")

        res, err :=  db.Collection("collID").Delete()
    }
```

### Databases

#### Create Database

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        dbDef, err := client.Databases().Create("DATABASE_ID")
    }
```

#### Read Database

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        dbDef, err := client.Database("DATABASE_ID").Read()
    }
```

#### List Databases

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        dbDefs, err = client.Databases().ReadAll()
    }
```

#### Delete Database

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        _, err = client.Database(dbID).Delete()
    }
```

### StoredProcedures

#### Create Stored Procedure

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        spDef := &cosmos.StoredProcedureDefinition{
            Resource: cosmos.Resource{ID: "mySP"},
            Body: "function () {\r\n var context = getContext();\r\n var response = context.getResponse();\r\n\r\n  response.setBody(\"Hello, World\");\r\n}"
        }

        createdSP, err := coll.StoredProcedures().Create(spDef)
    }
```

#### Execute Stored Procedure

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

     func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        var res string
        _, err = coll.StoredProcedure("mySP").Execute("", &res)
    }
```

#### Replace Stored Procedure

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

     func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        newSpDef := &cosmos.StoredProcedureDefinition{
            Resource: cosmos.Resource{ID: "mySP"},
            Body: "function (greet, someone) {\r\n var context = getContext();\r\n var response = context.getResponse();\r\n\r\n response.setBody(greet + \", \"+ someone);\r\n}"
        }

        _, err = coll.StoredProcedure("mySP").Replace(newSpDef)
    }
```

#### List Stored Procedures

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        sprocs, err := coll.StoredProcedures().ReadAll()
    }
```

#### Delete Stored Procedure

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        _, err = coll.StoredProcedure("mySP").Delete()
    }
```

### UDFs

#### Create UDF

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        udfDef := &cosmos.UDFDefinition{
		    Body: "function tax(income) {\r\n if(income == undefined) \r\n throw 'no input';\r\n if (income < 1000) \r\n return income * 0.1;\r\n else if (income < 10000) \r\n return income * 0.2;\r\n else\r\n return income * 0.4;\r\n}",
            Resource: cosmos.Resource{ID: "myUDF"},
        }
        createdUDF, err := coll.UDFs().Create(udfDef)
    }
```

#### Replace UDF

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        newUDF := &cosmos.UDFDefinition{
		    Body: "function tax(income) {\r\n if(income == undefined) \r\n throw 'no input';\r\n if (income     < 2000) \r\n return income * 0.1;\r\n else if (income < 10000) \r\n return income * 0.2;\r\n    else\r\n return income * 0.4;\r\n}",
            Resource: cosmos.Resource{ID: "myUDF"},
        }
        updatedUDF, err := coll.UDF("myUDF").Replace(newUDF)
    }
```

#### List UDFs

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        udfs, err := coll.UDFs().ReadAll()
    }
```

#### Delete UDF

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        _, err = coll.UDF("myUDF").Delete()
    }
```

### Triggers

#### Create Trigger

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        triggerDef := &cosmos.TriggerDefintion{
		    Body: "function updateMetadata() {\r\n  var context = getContext();\r\nvar collection = context.getCollection   ();\r\nvar response = context.getResponse();\r\nvar createdDocument = response.getBody();\r\n\r\n// query for metadata     document\r\nvar filterQuery = 'SELECT * FROM root r WHERE r.id = \"_metadata\"';\r\nvar accept = collection.queryDocuments  (collection.getSelfLink(), filterQuery,\r\n  updateMetadataCallback);\r\n    if(!accept) throw \"Unable to update metadata, abort\";\r\n\r\nfunction updateMetadataCallback(err, documents, responseOptions) {\r\n  if(err) throw new Error (\"Error\" + err.message);\r\n   if(documents.length != 1) throw 'Unable to find metadata document';\r\n   var metadataDocument = documents[0];\r\n\r\n   // update metadata\r\n   metadataDocument.createdDocuments += 1;\r\n     metadataDocument.createdNames += \" \" + createdDocument.id;\r\nvar accept = collection.replaceDocument (metadataDocument._self,\r\n metadataDocument, function(err, docReplaced) {\r\n       if(err) throw \"Unable to update  metadata, abort\";\r\n });\r\nif(!accept) throw \"Unable to update metadata, abort\";\r\nreturn;\r\n    }",
		    Resource:         cosmos.Resource{ID: "myTrigger"},
		    TriggerOperation: "All",
            TriggerType:      "Post",
        }
        _, err := coll.Triggers().Create(triggerDef)
    }
```

#### Replace Trigger

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        newTriggerDef := &cosmos.TriggerDefintion{
		    Body: "function updateMetadata() {\r\n  var context = getContext();\r\nvar collection = context.getCollection   ();\r\nvar response = context.getResponse();\r\nvar createdDocument = response.getBody();\r\n\r\n// query for metadata     document\r\nvar filterQuery = 'SELECT * FROM root r WHERE r.id = \"_metadata\"';\r\nvar accept = collection.queryDocuments  (collection.getSelfLink(), filterQuery,\r\n  updateMetadataCallback);\r\n    if(!accept) throw \"Unable to update     metadata, exit\";\r\n\r\nfunction updateMetadataCallback(err, documents, responseOptions) {\r\n  if(err) throw new Error    (\"Error\" + err.message);\r\n   if(documents.length != 1) throw 'Unable to find metadata document';\r\n   var  metadataDocument = documents[0];\r\n\r\n   // update metadata\r\n   metadataDocument.createdDocuments += 1;\r\n      metadataDocument.createdNames += \" \" + createdDocument.id;\r\nvar accept = collection.replaceDocument    (metadataDocument._self,\r\n    metadataDocument, function(err, docReplaced) {\r\n       if(err) throw \"Unable to update   metadata, abort\";\r\n    });\r\nif(!accept) throw \"Unable to update metadata, abort\";\r\nreturn;\r\n    }",
		    Resource:         cosmos.Resource{ID: "myTrigger"},
		    TriggerOperation: "All",
            TriggerType:      "Post",
        }
        updatedTriggerDef, err := coll.Trigger("myTrigger").Replace(newTriggerDef)
    }
```

#### List Triggers

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        coll := client.Database("dbID").Collection("collID")
        triggers, err := coll.Triggers().ReadAll()
    }
```

#### Delete Trigger

```GO
    import "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"

    func main() {
        client, err := cosmos.New("YOUR_CONNECTION_STRING")
        _, err = coll.Trigger("myTrigger").Delete()
    }
```

### QueryBuilder

CosmosDB sdk for go includes a simple query builder.

supports:
* AND
* OR
* SELECT
* FROM
* ORDER BY

#### Example 1
```GO
import (
    "github.com/SpacyCoder/cosmosdb-go-sdk/qbuilder"
    "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"
)

func main()  {
    client, err := cosmos.New("YOUR_CONNECTION_STRING")

    qb := qbuilder.New()
    query := qb.Select("*").From("root").And("root.age > 10").Build()

    var people []People
    client.Database("mydb").Collection("people").Documents().Query(query, people, Cosmos.CrossPartition())
}
```

#### Example 2
```GO
import (
    "github.com/SpacyCoder/cosmosdb-go-sdk/qbuilder"
    "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"
)

func main()  {
    client, err := cosmos.New("YOUR_CONNECTION_STRING")
import "github.com/SpacyCoder/cosmosdb-go-sdk/qbuilder"
 "github.com/SpacyCoder/cosmosdb-go-sdk/cosmos"
    qb := qbuilder.New()
    q1 := qb.Select("*").From("root").And("root.age > @AGE").And("root.height > @HEIGHT").OrderBy("DESC root.height")
    query := q1.Params(cosmos.P{Name: "@AGE", Value: 20}, cosmos.P{Name: "@HEIGHT", Value: 180}).Build()

    var people []People
    client.Database("mydb").Collection("people").Documents().Query(query, people, Cosmos.CrossPartition())
}
```
