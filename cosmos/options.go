package cosmos

import (
	"encoding/json"
	"strconv"
)

// Consistency type to define consistency levels
type Consistency string

const (
	// Strong consistency level
	Strong Consistency = "strong"

	// Bounded consistency level
	Bounded Consistency = "bounded"

	// Session consistency level
	Session Consistency = "session"

	// Eventual consistency level
	Eventual Consistency = "eventual"
)

// CallOption type defenition.
type CallOption func(r *Request) error

// PartitionKey sets partition for request
func PartitionKey(partitionKey interface{}) CallOption {
	var pk []byte
	var err error

	switch v := partitionKey.(type) {
	case json.Marshaler:
		pk, err = Serialization.Marshal(v)
	default:
		pk, err = Serialization.Marshal([]interface{}{v})
	}

	header := []string{string(pk)}

	return func(r *Request) error {
		if err != nil {
			return err
		}
		r.Header[HeaderPartitionKey] = header
		return nil
	}
}

// Continuation sets continuation token for request
func Continuation(continuation string) CallOption {
	return func(r *Request) error {
		if continuation == "" {
			return nil
		}
		r.Header.Set(HeaderContinuation, continuation)
		return nil
	}
}

// Upsert if set to true, Cosmos DB creates the document with the ID (and partition key value if applicable) if it doesn’t exist, or update the document if it exists.
func Upsert() CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderUpsert, "true")
		return nil
	}
}

// CrossPartition allows query to run on all partitions
func CrossPartition() CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderCrossPartition, "true")
		return nil
	}
}

// ChangeFeedPartitionRangeID used in change feed requests. The partition key range ID for reading data.
func ChangeFeedPartitionRangeID(id string) CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderPartitionKeyRangeID, id)
		return nil
	}
}

// ChangeFeed indicates a change feed request
func ChangeFeed() CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderAIM, "Incremental feed")
		return nil
	}
}

// IfModifiedSince returns etag of resource modified after specified date in RFC 1123 format. Ignored when If-None-Match is specified
// Optional (applicable only on GET)
func IfModifiedSince(date string) CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderIfModifiedSince, date)
		return nil
	}
}

// IfNoneMatch makes operation conditional to only execute if the resource has changed. The value should be the etag of the resource.
// Optional (applicable only on GET)
func IfNoneMatch(etag string) CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderIfNonMatch, etag)
		return nil
	}
}

// IfMatch used to make operation conditional for optimistic concurrency. The value should be the etag value of the resource.
// (applicable only on PUT and DELETE)
func IfMatch(etag string) CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderIfMatch, etag)
		return nil
	}
}

// SessionToken a string token used with session level consistency.
func SessionToken(sessionToken string) CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderSessionToken, sessionToken)
		return nil
	}
}

// ConsistencyLevel override for read options against documents and attachments. The valid values are: Strong, Bounded, Session, or Eventual (in order of strongest to weakest). The override must be the same or weaker than the account�s configured consistency level.
func ConsistencyLevel(consistency Consistency) CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderConsistency, string(consistency))
		return nil
	}
}

// Limit set max item count for response
func Limit(limit int) CallOption {
	header := strconv.Itoa(limit)
	return func(r *Request) error {
		r.Header.Set(HeaderMaxItemCount, header)
		return nil
	}
}

// Sets headers necessary for doing a query
func queryHeaders(len int) CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderContentType, "application/query+json")
		r.Header.Set(HeaderIsQuery, "true")
		r.Header.Set(HeaderContentLength, strconv.Itoa(len))
		return nil
	}
}
