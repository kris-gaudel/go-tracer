package interval

import "go-tracer/src/utils"

type Interval struct {
	Min float64
	Max float64
}

func (i *Interval) Contains(value float64) bool {
	return i.Min <= value && value <= i.Max
}

func (i *Interval) Surrounds(value float64) bool {
	return i.Min < value && value < i.Max
}

var EmptyInterval Interval = Interval{Min: utils.INFINITY, Max: -utils.INFINITY}
var UniverseInterval Interval = Interval{Min: -utils.INFINITY, Max: utils.INFINITY}
