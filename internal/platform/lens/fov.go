package lens

import "math"

// Calculates lens field of view in degrees based on sensor or film plane size and lens focal length.
// Returns 0 field of view angle for 0 sized plane.
func FieldOfView(plane PlaneSize, focalLength float64) FieldOfViewAngle {
	if plane.Height == 0 || plane.Width == 0 || focalLength == 0 {
		return FieldOfViewAngle{}
	}

	return FieldOfViewAngle{
		fovAngle(plane.Height, focalLength),
		fovAngle(plane.Width, focalLength),
	}
}

func fovAngle(sensorSideSize float64, focalLength float64) float64 {
	return 2 * math.Atan(sensorSideSize/(2*focalLength)) * (180 / math.Pi)
}
