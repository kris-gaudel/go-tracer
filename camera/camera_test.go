package camera

import (
	"go-tracer/src/hittable"
	"go-tracer/src/vec3"
	"log"
	"os"
	"testing"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func TestCameraRender(t *testing.T) {
	log.Println("Running TestCameraRender")
	outputFile, err := os.Create("test_output.ppm")
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	log.Println("Writing to file...")

	oldOutput := os.Stdout
	os.Stdout = outputFile

	var world hittable.HittableList
	ground_material := hittable.Lambertian{Albedo: vec3.Vec3{X: 0.5, Y: 0.5, Z: 0.5}}
	sphereOne := hittable.Sphere{Center: vec3.Vec3{X: 0, Y: -1000, Z: 0}, Radius: 1000, Mat: ground_material}
	sphereTwo := hittable.Sphere{Center: vec3.Vec3{X: 0, Y: 1, Z: 0}, Radius: 1, Mat: ground_material}
	world.Append(&sphereOne)
	world.Append(&sphereTwo)

	// Camera
	var cam Camera
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 1200
	cam.SamplesPerPixel = 500
	cam.MaxDepth = 50
	cam.VFOV = 20.0
	cam.LookFrom = vec3.Point3{X: 13, Y: 2, Z: 3}
	cam.LookAt = vec3.Point3{X: 0, Y: 0, Z: 0}
	cam.ViewUp = vec3.Vec3{X: 0, Y: 1, Z: 0}
	cam.DefocusAngle = 0.6
	cam.FocusDistance = 10.0
	cam.Render(world)

	os.Stdout = oldOutput
	log.Println("Done writing to file!")

	imagick.Initialize()
	defer imagick.Terminate()

	// Load the rendered image and expected image
	output := imagick.NewMagickWand()
	err = output.ReadImage("test_output.ppm")
	if err != nil {
		t.Fatalf("Failed to open rendered image: %v", err)
	}

	expected := imagick.NewMagickWand()
	err = expected.ReadImage("expected.ppm")
	if err != nil {
		t.Fatalf("Failed to open expected image: %v", err)
	}

	_, metric := output.CompareImages(expected, imagick.METRIC_MEAN_SQUARED_ERROR)
	if metric > 0.0001 {
		t.Fatalf("Mean Squared Error: %f", metric)
	}

	output.Destroy()
	expected.Destroy()

	log.Println("Finished running TestCameraRender")
}
