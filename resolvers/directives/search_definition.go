package directives

type BoxOperatorInput struct {
	BottomLeft *CoordinatesInput `json:"bottomLeft,omitempty"`
	UpperRight *CoordinatesInput `json:"upperRight,omitempty"`
}
type CenterOperatorInput struct {
	CoordinatesInput
	Radius         *float64 `json:"radius,omitempty"`
	RadiusOperator *string  `json:"radiusOperator,omitempty"`
}
type GeometryOperatorInput struct {
	Type        string             `json:"type,omitempty"`
	Coordinates []CoordinatesInput `json:"coordinates,omitempty"`
}
type CoordinatesInput struct {
	Lat *float64 `json:"lat,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
}

type GeoSpatialNear struct {
	PointCoordinates *CoordinatesInput `json:"pointCoordinates,omitempty"`
	MaxDistance      *float64          `json:"maxDistance,omitempty"`
}
type GeoSpatialGeoWithinOperators struct {
	Box          BoxOperatorInput      `json:"box,omitempty"`
	Center       CenterOperatorInput   `json:"center,omitempty"`
	CenterSphere CenterOperatorInput   `json:"centerSphere,omitempty"`
	Geometry     GeometryOperatorInput `json:"geometry,omitempty"`
	Polygon      []CoordinatesInput    `json:"polygon,omitempty"`
}

func (o *CenterOperatorInput) calc() (r float32) {
	switch *o.RadiusOperator {
	case "kilometer":
		r = float32(*o.Radius) / 6378.1
		break
	case "miles":
		r = float32(*o.Radius) / 3963.2
		break
	case "meter":
		r = float32(*o.Radius) / (6378.1 * 1000)
		break
	}
	return
}

func (o *CenterOperatorInput) deegreeCalc() (r float32) {
	switch *o.RadiusOperator {
	case "kilometer":
		r = float32(*o.Radius) / 111.32
		break
	case "miles":
		r = float32(*o.Radius) / 69.17
		break
	case "meter":
		r = float32(*o.Radius) / (111.32 * 1000)
		break
	}
	return
}
