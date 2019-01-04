package common

import (
	"math/rand"

	"github.com/faiface/pixel"
)

type camera struct {
	Position pixel.Vec
	Speed    float64
	Zoom     float64
	ZSpeed   float64
}

// RandFloat produces a random Float64 in the range [low, high].
func RandFloat(low, high float64) float64 {
	return rand.Float64()*(high-low) + low
}

// RandInt produces a random Int in the range [low, high].
func RandInt(low, high int) int {
	return rand.Intn(high-low+1) + low
}

// === DUPLICATES pixel.Clamp(float64) float64 ===
// Clamp forces a number to be in the range [min, max].
// func Clamp(num, min, max float64) float64 {
// 	if num < min {
// 		return min
// 	}
// 	if num > max {
// 		return max
// 	}
// 	return num
// }

// Clamp forces a integer to be in the range [min, max].
func ClampInt(num, min, max int) int {
	if num < min {
		return min
	}
	if num > max {
		return max
	}
	return num
}

func randomMovement(x, y float64) (float64, float64) {
	return x + randFloat(-POINT_SPEED, POINT_SPEED),
		y + randFloat(-POINT_SPEED, POINT_SPEED)
}

func centerDrift(x, y float64) (float64, float64) {
	// go directly towards center 2% of the time
	if rand.Float32() < 0.02 {
		pos := pixel.V(x, y)
		center := pixel.V(WIDTH/2, HEIGHT/2)
		delta := pos.To(center).
			Unit().
			Scaled(POINT_SPEED)
		return x + delta.X, y + delta.Y
	}

	return randomMovement(x, y)
}
