package qbuilder

import "github.com/spacycoder/cosmosdb-go-sdk/cosmos"

type Condition struct {
	ConditionType string
	Value         string
}

type QueryBuilder struct {
	selectors  []string
	from       string
	conditions []Condition
	params     []cosmos.QueryParam
	orderBy    string
}

// New creates a new QueryBuilder
func New() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) Select(s ...string) *QueryBuilder {
	qb.selectors = s
	return qb
}

func (qb *QueryBuilder) From(f string) *QueryBuilder {
	qb.from = f
	return qb
}

func (qb *QueryBuilder) And(conditionsValues ...string) *QueryBuilder {
	var conditions []Condition
	for _, c := range conditionsValues {
		condition := Condition{ConditionType: "AND", Value: c}
		conditions = append(conditions, condition)
	}
	qb.conditions = append(qb.conditions, conditions...)
	return qb
}

func (qb *QueryBuilder) Or(conditionsValues ...string) *QueryBuilder {
	var conditions []Condition
	for _, c := range conditionsValues {
		condition := Condition{ConditionType: "OR", Value: c}
		conditions = append(conditions, condition)
	}
	qb.conditions = append(qb.conditions, conditions...)
	return qb
}

func (qb *QueryBuilder) Params(params ...cosmos.QueryParam) *QueryBuilder {
	qb.params = append(qb.params, params...)
	return qb
}

func (qb *QueryBuilder) OrderBy(orderBy string) *QueryBuilder {
	qb.orderBy = " ORDER BY " + orderBy
	return qb
}

func (qb *QueryBuilder) Build() *cosmos.SqlQuerySpec {
	query := "SELECT "
	for i, s := range qb.selectors {
		if i == 0 {
			query += s
		} else {
			query += ", " + s
		}
	}

	query += " FROM " + qb.from
	if len(qb.conditions) == 0 {
		return cosmos.Q(query)
	}

	query += " WHERE"
	for i, c := range qb.conditions {
		if i == 0 {
			query += " " + c.Value
		} else {
			query += " " + c.ConditionType + " " + c.Value
		}
	}

	query += qb.orderBy
	return cosmos.Q(query, qb.params...)
}
