package cosmos

import (
	"context"
	"encoding/json"
	"io"
)

func (q SqlQuerySpec) Read(p []byte) (n int, err error) {
	b, err := json.Marshal(q)
	copy(p, b)
	return len(b), io.EOF
}

// Document performs operations on a given document.
type Document struct {
	client Client
	coll   Collection
	docID  string
}

// Documents performs operations on a multiple documents.
type Documents struct {
	client Client
	coll   Collection
}

func newDocument(coll Collection, docID string) *Document {
	coll.client.path += "/docs/" + docID
	coll.client.rType = "docs"
	coll.client.rLink = coll.client.path
	doc := &Document{
		client: coll.client,
		coll:   coll,
		docID:  docID,
	}

	return doc
}

func newDocuments(coll Collection) *Documents {
	coll.client.path += "/docs"
	coll.client.rType = "docs"
	docs := &Documents{
		client: coll.client,
		coll:   coll,
	}

	return docs
}

// Create new document
func (d *Documents) Create(ctx context.Context, doc interface{}, opts ...CallOption) (*Response, error) {
	d.client.createIDIfNotSet(doc)
	return d.client.create(ctx, doc, &doc, opts...)
}

// Read document
func (d Document) Read(ctx context.Context, ret interface{}, opts ...CallOption) (*Response, error) {
	return d.client.read(ctx, ret, opts...)
}

// Replace existing document
func (d *Document) Replace(ctx context.Context, doc interface{}, ret interface{}, opts ...CallOption) (*Response, error) {
	d.client.createIDIfNotSet(doc)
	return d.client.replace(ctx, doc, ret, opts...)
}

// Delete document
func (d Document) Delete(ctx context.Context, opts ...CallOption) (*Response, error) {
	return d.client.delete(ctx, opts...)
}

// ReadAll returns all documents in collection.
func (d *Documents) ReadAll(ctx context.Context, docs interface{}, opts ...CallOption) (*Response, error) {
	data := struct {
		Documents interface{} `json:"Documents,omitempty"`
		Count     int         `json:"_count,omitempty"`
	}{Documents: docs}
	res, err := d.client.read(ctx, &data, opts...)
	return res, err
}

// Query documents
func (d Documents) Query(ctx context.Context, query *SqlQuerySpec, docs interface{}, opts ...CallOption) (*Response, error) {
	data := struct {
		Documents interface{} `json:"Documents,omitempty"`
		Count     int         `json:"_count,omitempty"`
	}{Documents: docs}
	res, err := d.client.query(ctx, query, &data, opts...)
	return res, err
}
