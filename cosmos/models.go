package cosmos

// Resource every cosmos resource have these.
type Resource struct {
	ID   string `json:"id,omitempty"`
	Self string `json:"_self,omitempty"`
	Etag string `json:"_etag,omitempty"`
	Rid  string `json:"_rid,omitempty"`
	Ts   int    `json:"_ts,omitempty"`
}

// PolicyIndex
type PolicyIndex struct {
	DataType  string `json:"dataType,omitempty"`
	Precision int    `json:"precision,omitempty"`
	Kind      string `json:"kind,omitempty"`
}

// IndexingPolicyPath model
type IndexingPolicyPath struct {
	Path    string        `json:"path,omitempty"`
	Indexes []PolicyIndex `json:"indexes,omitempty"`
}

// IndexingPolicy model
type IndexingPolicy struct {
	IndexingMode  string               `json:"indexingMode,omitempty"`
	Automatic     bool                 `json:"automatic,omitempty"`
	IncludedPaths []IndexingPolicyPath `json:"includedPaths,omitempty"`
	ExcludedPaths []IndexingPolicyPath `json:"excludedPaths,omitempty"`
}

// DatabaseDefinition defines the structure of a database data query
type DatabaseDefinition struct {
	Resource
	Colls string `json:"_colls,omitempty"`
	Users string `json:"_users,omitempty"`
}

// DatabaseDefinitions slice of Database elements
type DatabaseDefinitions []DatabaseDefinition

// Length returns the number of databases.
func (d *DatabaseDefinitions) Length() int {
	return len(*d)
}

// CollectionDefinition defiens the structure of a Collection
type CollectionDefinition struct {
	Resource
	IndexingPolicy IndexingPolicy         `json:"indexingPolicy,omitempty"`
	PartitionKey   PartitionKeyDefinition `json:"partitionKey,omitempty"`
	Docs           string                 `json:"_docs,omitempty"`
	Udf            string                 `json:"_udfs,omitempty"`
	Sporcs         string                 `json:"_sporcs,omitempty"`
	Triggers       string                 `json:"_triggers,omitempty"`
	Conflicts      string                 `json:"_conflicts,omitempty"`
}

// PartitionKeyDefinition is used to define a partitionkey.
type PartitionKeyDefinition struct {
	Paths []string `json:"paths"`
	Kind  string   `json:"kind"`
}

// CollectionDefinitions is a slice of CollectionDefinition elements
type CollectionDefinitions []CollectionDefinition

// Length returns the number of collections.
func (c *CollectionDefinitions) Length() int {
	return len(*c)
}

// At return collection at the specified index
func (c CollectionDefinitions) At(index int) *CollectionDefinition {
	return &c[index]
}

// DocumentDefinition defines all the default document properties.
type DocumentDefinition struct {
	Resource
	Attachments string `json:"attachments,omitempty"`
}

// StoredProcedureDefinition stored procedure model
type StoredProcedureDefinition struct {
	Resource
	Body string `json:"body,omitempty"`
}

// UDFDefinition (User Defined Function) definition
type UDFDefinition struct {
	Resource
	Body string `json:"body,omitempty"`
}

// PartitionKeyRange partition key range model
type PartitionKeyRange struct {
	Resource
	PartitionKeyRangeID string `json:"id,omitempty"`
	MinInclusive        string `json:"minInclusive,omitempty"`
	MaxInclusive        string `json:"maxExclusive,omitempty"`
}

// TriggerDefinition defines the model of a cosmos trigger.
type TriggerDefintion struct {
	Resource
	Body             string `json:"body"`
	TriggerOperation string `json:"triggerOperation"`
	TriggerType      string `json:"triggerType"`
}

type MsToken struct {
	Token string `json:"token"`
	Range struct {
		Min string `json:"min"`
		Max string `json:"max"`
	} `json:"range"`
}
