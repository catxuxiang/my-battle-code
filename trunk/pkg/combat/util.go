package combat

import (
	"math"
)

func GetDistance(pos0 point, pos1 point) float64 {
	return math.Sqrt((pos0.x-pos1.x)*(pos0.x-pos1.x) + (pos0.y-pos1.y)*(pos0.y-pos1.y))
}
