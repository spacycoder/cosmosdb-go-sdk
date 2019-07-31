package cosmos

import "testing"

func TestLineString(t *testing.T) {
	coords := Coordinates{{5.0, 10.0}, {10.0, 11.0}, {5.0, 10.0}}
	ls := NewLineString()
	for _, p := range coords {
		ls.AddPoint(p[0], p[1])
	}

	if len(ls.Coordinates) != 3 {
		t.Fatalf("expected %d coords, got: %d", 3, len(ls.Coordinates))
	}

	for i, c := range *ls.Coords() {
		if c[0] != coords[i][0] {
			t.Fatalf("Invalid longitude. Expected: %f, got: %f", coords[i][0], c[0])
		}
		if c[1] != coords[i][1] {
			t.Fatalf("Invalid latitude. Expected: %f, got: %f", coords[i][1], c[1])
		}
	}

	ls2 := NewLineString(Coordinate{5.0, 10}, Coordinate{10, 11}, Coordinate{5, 10})

	if len(ls2.Coordinates) != 3 {
		t.Fatalf("expected %d coords got: %d", 3, len(ls.Coordinates))
	}

	if ls2.GeoType() != "LineString" {
		t.Fatalf("expected: LineString, got: %s", ls2.GeoType())
	}
}

func TestPolygon(t *testing.T) {
	coords := []Coordinates{{{5.0, 10.0}, {10.0, 11.0}, {5.0, 10.0}}}
	poly := NewPolygon(coords...)

	if len(poly.Coords()) != len(coords) {
		t.Fatalf("expected %d coords, got: %d", len(coords), len(poly.Coords()))
	}

	for i, cc := range poly.Coords() {
		for j, c := range cc {
			coord := coords[i][j]
			if c[0] != coord[0] {
				t.Fatalf("Invalid longitude. Expected: %f, got: %f", coord[0], c[0])
			}
			if c[1] != coord[1] {
				t.Fatalf("Invalid latitude. Expected: %f, got: %f", coord[1], c[1])
			}
		}

	}

	if poly.GeoType() != "Polygon" {
		t.Fatalf("expected: Polygon, got: %s", poly.GeoType())
	}
}

func TestPoint(t *testing.T) {
	coord := Coordinate{5.0, 10.0}
	point := NewPoint(coord)

	if point.Coords()[0] != coord[0] {
		t.Fatalf("Invalid longitude. Expected: %f, got: %f", coord[0], point.Coords()[0])
	}
	if point.Coords()[1] != coord[1] {
		t.Fatalf("Invalid latitude. Expected: %f, got: %f", coord[1], point.Coords()[1])
	}
	if point.GeoType() != "Point" {
		t.Fatalf("expected: Point, got: %s", point.GeoType())
	}
}
