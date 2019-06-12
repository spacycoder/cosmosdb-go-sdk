package cosmos

import "context"

// Iterator used to iterate through documents.
type Iterator struct {
	continuationToken string
	err               error
	response          *Response
	next              bool
	source            IteratorFunc
	docs              *Documents
}

// NewIterator creates iterator instance
func NewIterator(docs *Documents, query *SqlQuerySpec, data interface{}, opts ...CallOption) *Iterator {
	return &Iterator{
		docs: docs,
		next: true,
		source: func(docs *Documents, internalOpts ...CallOption) (*Response, error) {
			return docs.Query(context.Background(), query, data, append(opts, internalOpts...)...)
		},
	}
}

// Response returns *Response object from last call
func (di *Iterator) Response() *Response {
	return di.response
}

// Errror returns error from last call
func (di *Iterator) Error() error {
	return di.err
}

// Next will ask iterator source for results and checks whenever there some more pages left
func (di *Iterator) Next() bool {
	if !di.next {
		return false
	}
	di.response, di.err = di.source(di.docs, Continuation(di.continuationToken))
	if di.err != nil {
		return false
	}
	di.continuationToken = di.response.Continuation()
	next := di.next
	di.next = di.continuationToken != ""
	return next
}

// IteratorFunc is type that describes iterator source
type IteratorFunc func(docs *Documents, internalOpts ...CallOption) (*Response, error)
