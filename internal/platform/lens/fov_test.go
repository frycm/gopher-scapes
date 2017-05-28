package lens_test

import (
	"math"
	"testing"

	"github.com/frycm/gopher-scapes/internal/platform/lens"
)

func TestFieldOfView0Values(t *testing.T) {
	defer func() {
		err := recover()

		if err != nil {
			t.Errorf("FieldOfView should not panic but it does with: %v", err)
		}
	}()

	fov := lens.FieldOfView(lens.PlaneSize{0, 1}, 50)

	if fov.Horizontal != 0 && fov.Vertical != 0 {
		t.Error("lens.PlaneSize with 0 width should return zero value of lens.FieldOfViewAngle")
	}

	fov = lens.FieldOfView(lens.PlaneSize{1, 0}, 50)

	if fov.Horizontal != 0 && fov.Vertical != 0 {
		t.Error("lens.PlaneSize with 0 height should return zero value of lens.FieldOfViewAngle")
	}

	fov = lens.FieldOfView(lens.PlaneSize{1, 1}, 0)

	if fov.Horizontal != 0 && fov.Vertical != 0 {
		t.Error("0 focalLength should return zero value of lens.FieldOfViewAngle")
	}

	fov = lens.FieldOfView(lens.PlaneSize{1, 0}, 0)

	if fov.Horizontal != 0 && fov.Vertical != 0 {
		t.Error("0 lensPlanceSize and 0 focalLength should return zero value of lens.FieldOfViewAngle")
	}
}

func TestFieldOfViewFF50mm(t *testing.T) {
	fov := lens.FieldOfView(lens.PlaneSize35mm, 50)

	horizontalAngle := 2 * math.Atan(float64(36)/(2*50)) * (180 / math.Pi)
	verticalAngle := 2 * math.Atan(float64(24)/(2*50)) * (180 / math.Pi)

	if fov.Horizontal != horizontalAngle {
		t.Errorf("%2f horizontal angle was expected on full frame withm 50mm lens, but %2f was calculated", horizontalAngle, fov.Horizontal)
	}
	if fov.Vertical != verticalAngle {
		t.Errorf("%2f vertical angle was expected on full frame withm 50mm lens, but %2f was calculated", horizontalAngle, fov.Horizontal)
	}
}
