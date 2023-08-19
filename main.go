package main

import (
	"fmt"
	"go-tracer/src/color"
	"go-tracer/src/vec3"
	"log"
	"strconv"
)

func main() {
	// Image
	var image_width int = 256
	var image_height int = 256

	// Render
	fmt.Println("P3")
	fmt.Println(strconv.Itoa(image_width) + " " + strconv.Itoa(image_height))
	fmt.Println("255")
	for j := 0; j < image_height; j++ {
		log.Println("Scanlines remaining: " + strconv.Itoa(image_height-j))
		for i := 0; i < image_width; i++ {
			base := vec3.Vec3{X: float64(i) / float64(image_width-1), Y: float64(j) / float64(image_height-1), Z: 0}
			pixel_color := color.Color{ColorX: 0, ColorY: 0, ColorZ: 0}
			pixel_color = pixel_color.ConvertToRGB(base)
			fmt.Println(pixel_color.String())
		}
	}
	log.Println("Done!")
}
