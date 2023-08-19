package color

import (
	"fmt"
	"go-tracer/src/vec3"
)

const (
	COLOR_MAX = 255.999
)

type Color struct {
	ColorX, ColorY, ColorZ int
}

func (c Color) ConvertToRGB(rawVec vec3.Vec3) Color {
	r := int(COLOR_MAX * rawVec.GetX())
	g := int(COLOR_MAX * rawVec.GetY())
	b := int(COLOR_MAX * rawVec.GetZ())
	return Color{r, g, b}
}

func (c Color) String() string {
	return fmt.Sprintf("%d %d %d", c.ColorX, c.ColorY, c.ColorZ)
}
