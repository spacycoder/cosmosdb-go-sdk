package cosmos

import (
	"testing"
)

func TestQuery(t *testing.T) {
	queryTests := []struct {
		input               *SqlQuerySpec
		expectedParamLength int
		expectedQueryString string
	}{
		{Q("SELECT * FROM root"), 0, "SELECT * FROM root"},
		{Q("SELECT * FROM root WHERE root.length < @LENGTH AND  root.age > @AGE",
			P{Name: "@LENGTH", Value: 180},
			P{Name: "@AGE", Value: 30}), 2, "SELECT * FROM root WHERE root.length < @LENGTH AND  root.age > @AGE"},
	}

	for _, queryTest := range queryTests {
		if len(queryTest.input.Parameters) != queryTest.expectedParamLength {
			t.Fatalf("query should contain %d parameters but has: %d", queryTest.expectedParamLength, len(queryTest.input.Parameters))
		}

		if queryTest.input.Query != queryTest.expectedQueryString {
			t.Fatalf("query should be %s but is %s", queryTest.expectedQueryString, queryTest.input.Query)
		}
	}
}

func TestLineString(t *testing.T) {
	ls := NewLineString()
	ls.AddPoint(5.0, 10)
	ls.AddPoint(10, 11)
	ls.AddPoint(5, 10)

	if len(ls.Coordinates) != 3 {
		t.Fatalf("LineString should contain %d points but has: %d", 3, len(ls.Coordinates))
	}

	ls = NewLineString(Coordinate{5.0, 10}, Coordinate{10, 11}, Coordinate{5, 10})

	if len(ls.Coordinates) != 3 {
		t.Fatalf("LineString should contain %d points but has: %d", 3, len(ls.Coordinates))
	}

	ls = NewLineString(Coordinate{5.0, 10}, Coordinate{10, 11})
	ls.AddPoint(5, 10)
}
