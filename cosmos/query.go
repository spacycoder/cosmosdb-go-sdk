package cosmos

type QueryParam struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type SqlQuerySpec struct {
	Query      string       `json:"query"`
	Parameters []QueryParam `json:"parameters,omitempty"`
}

func Q(query string, queryParams ...QueryParam) *SqlQuerySpec {
	return &SqlQuerySpec{Query: query, Parameters: queryParams}
}

type P = QueryParam
