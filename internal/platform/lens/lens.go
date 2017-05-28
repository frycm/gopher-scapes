package lens

// Sensor or film plane size in mm
type PlaneSize struct {
	Width  float64
	Height float64
}

// Commonly used plane sizes
var (
	PlaneSize35mm = PlaneSize{36.00, 24.00}
	PlaneSizeApsc = PlaneSize{23.60, 15.70}
	PlaneSizeCanonApsc = PlaneSize{22.20, 14.80}
	PlaneSizeMicro43 = PlaneSize{17.30, 13.00}
	PlaneSizeMavicPhantom = PlaneSize{6.30, 4.70}
	PlaneSizeIPhone12M = PlaneSize{4.80, 3.60}
)

// Lens field of view in degrees
type FieldOfViewAngle struct {
	Vertical   float64
	Horizontal float64
}
