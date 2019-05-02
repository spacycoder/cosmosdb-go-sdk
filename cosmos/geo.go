package cosmos

// Coordinate = [lon, lat]
type Coordinate [2]float64

type Coordinates []Coordinate

type Geometry interface {
	GeoType() string
}

// LineString struct defines a line string
type LineString struct {
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates"`
}

// NewLineString creates a new LineString struct.
func NewLineString(coords ...Coordinate) *LineString {
	line := &LineString{Type: "LineString", Coordinates: coords}
	return line
}

// AddPoint is a helper method for adding point to a LineString
func (l *LineString) AddPoint(lon, lat float64) {
	l.Coordinates = append(l.Coordinates, Coordinate{lon, lat})
}

func (l *LineString) Coords() *Coordinates {
	return &l.Coordinates
}

func (l *LineString) GeoType() string {
	return l.Type
}

type Polygon struct {
	Type        string        `json:"type"`
	Coordinates []Coordinates `json:"coordinates"`
}

// NewPolygon creates a new Polygon struct.
func NewPolygon(coords ...Coordinates) *Polygon {
	polygon := &Polygon{Type: "Polygon", Coordinates: coords}
	return polygon
}

func (p *Polygon) Coords() []Coordinates {
	return p.Coordinates
}

func (p *Polygon) GeoType() string {
	return p.Type
}

type Point struct {
	Type        string     `json:"type"`
	Coordinates Coordinate `json:"coordinates"`
}

// NewPoint creates a new point struct and returns it.
func NewPoint(coords Coordinate) *Point {
	line := &Point{Type: "Point", Coordinates: coords}
	return line
}

func (p *Point) Coords() *Coordinate {
	return &p.Coordinates
}

func (p *Point) GeoType() string {
	return p.Type
}
