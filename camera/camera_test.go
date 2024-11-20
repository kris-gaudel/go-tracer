package camera

import (
	"go-tracer/src/vec3"
	"math"
	"testing"
)

const EPSILON = 1e-7

func almostEqual(t *testing.T, got, want vec3.Vec3, msg string) {
	if math.Abs(got.X-want.X) > EPSILON ||
		math.Abs(got.Y-want.Y) > EPSILON ||
		math.Abs(got.Z-want.Z) > EPSILON {
		t.Errorf("%s: got %v, want %v", msg, got, want)
	}
}

func TestCameraInitialization(t *testing.T) {
	// Create a basic camera setup
	cam := Camera{
		AspectRatio:     16.0 / 9.0,
		VFOV:            90.0,
		LookFrom:        vec3.Point3{X: 0, Y: 0, Z: 0},
		LookAt:          vec3.Point3{X: 0, Y: 0, Z: -1},
		ViewUp:          vec3.Vec3{X: 0, Y: 1, Z: 0},
		DefocusAngle:    0.0,
		FocusDistance:   1.0,
		SamplesPerPixel: 10,
		MaxDepth:        50,
	}

	cam.Initalize()

	t.Run("Image dimensions", func(t *testing.T) {
		if cam.ImageWidth != 400 {
			t.Errorf("ImageWidth = %v, want %v", cam.ImageWidth, 400)
		}
		expectedHeight := int(float64(cam.ImageWidth) / cam.AspectRatio)
		if cam.ImageHeight != expectedHeight {
			t.Errorf("ImageHeight = %v, want %v", cam.ImageHeight, expectedHeight)
		}
	})

	t.Run("Camera orientation", func(t *testing.T) {
		// For this setup, W should point in positive Z direction
		wantW := vec3.Vec3{X: 0, Y: 0, Z: 1}
		almostEqual(t, cam.W, wantW, "Camera W vector")

		// U should point in positive X direction
		wantU := vec3.Vec3{X: 1, Y: 0, Z: 0}
		almostEqual(t, cam.U, wantU, "Camera U vector")

		// V should point in positive Y direction
		wantV := vec3.Vec3{X: 0, Y: 1, Z: 0}
		almostEqual(t, cam.V, wantV, "Camera V vector")
	})
}

func TestGetRay(t *testing.T) {
	cam := Camera{
		AspectRatio:     16.0 / 9.0,
		VFOV:            90.0,
		LookFrom:        vec3.Point3{X: 0, Y: 0, Z: 0},
		LookAt:          vec3.Point3{X: 0, Y: 0, Z: -1},
		ViewUp:          vec3.Vec3{X: 0, Y: 1, Z: 0},
		DefocusAngle:    0.0,
		FocusDistance:   1.0,
		SamplesPerPixel: 1,
		MaxDepth:        50,
	}

	cam.Initalize()

	t.Run("Center ray bounds", func(t *testing.T) {
		// Get ray from center of image
		centerI := cam.ImageWidth / 2
		centerJ := cam.ImageHeight / 2
		ray := cam.GetRay(centerI, centerJ)

		// The ray should be roughly pointing forward (-Z direction)
		// but will have some variation due to pixel sampling

		// Check Z component is predominantly negative and close to -1
		if ray.Direction.Z >= 0 || math.Abs(ray.Direction.Z+1) > 0.1 {
			t.Errorf("Ray Z direction out of expected bounds: %v, want close to -1", ray.Direction.Z)
		}

		// X and Y components should be relatively small for center rays
		if math.Abs(ray.Direction.X) > 0.1 {
			t.Errorf("Ray X direction too large: %v, want |x| < 0.1", ray.Direction.X)
		}
		if math.Abs(ray.Direction.Y) > 0.1 {
			t.Errorf("Ray Y direction too large: %v, want |y| < 0.1", ray.Direction.Y)
		}

		// Check origin is at camera center
		almostEqual(t, ray.Origin, cam.Center, "Ray origin")
	})
}

func TestDefocusDiskSample(t *testing.T) {
	cam := Camera{
		AspectRatio:     16.0 / 9.0,
		VFOV:            90.0,
		LookFrom:        vec3.Point3{X: 0, Y: 0, Z: 0},
		LookAt:          vec3.Point3{X: 0, Y: 0, Z: -1},
		ViewUp:          vec3.Vec3{X: 0, Y: 1, Z: 0},
		DefocusAngle:    45.0, // Significant defocus
		FocusDistance:   1.0,
		SamplesPerPixel: 1,
		MaxDepth:        50,
	}

	cam.Initalize()

	t.Run("Defocus disk bounds", func(t *testing.T) {
		// Sample multiple points and ensure they're within expected bounds
		for i := 0; i < 100; i++ {
			sample := cam.DefocusDiskSample()
			distanceFromCenter := sample.Subtract(cam.Center).Length()

			// The sample should be within the defocus disk radius
			defocusRadius := cam.FocusDistance * math.Tan(math.Pi/4) // 45 deg angle
			if distanceFromCenter > defocusRadius {
				t.Errorf("Defocus sample outside disk radius: got %v, max allowed %v",
					distanceFromCenter, defocusRadius)
			}
		}
	})
}

func TestPixelSampleSquare(t *testing.T) {
	cam := Camera{
		AspectRatio:     16.0 / 9.0,
		VFOV:            90.0,
		LookFrom:        vec3.Point3{X: 0, Y: 0, Z: 0},
		LookAt:          vec3.Point3{X: 0, Y: 0, Z: -1},
		ViewUp:          vec3.Vec3{X: 0, Y: 1, Z: 0},
		DefocusAngle:    0.0,
		FocusDistance:   1.0,
		SamplesPerPixel: 1,
		MaxDepth:        50,
	}

	cam.Initalize()

	t.Run("Sample bounds", func(t *testing.T) {
		// Test multiple samples to ensure they're within pixel bounds
		for i := 0; i < 100; i++ {
			sample := cam.PixelSampleSquare()

			// Sample should be within one pixel delta in any direction
			if math.Abs(sample.X) > cam.PixelDeltaU.Length() ||
				math.Abs(sample.Y) > cam.PixelDeltaV.Length() {
				t.Errorf("Pixel sample outside bounds: %v", sample)
			}
		}
	})
}
