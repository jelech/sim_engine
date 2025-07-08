// Package utils provides utility functions and helpers
package utils

import (
	"crypto/rand"
	"fmt"
	"math"
	"reflect"
	"time"
)

// GenerateID generates a unique identifier
func GenerateID() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 4)
	_, _ = rand.Read(randomBytes)
	return fmt.Sprintf("%d-%x", timestamp, randomBytes)
}

// Clamp clamps a value between min and max
func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Lerp performs linear interpolation between a and b
func Lerp(a, b, t float64) float64 {
	return a + t*(b-a)
}

// NormalizeAngle normalizes an angle to the range [0, 2π)
func NormalizeAngle(angle float64) float64 {
	for angle < 0 {
		angle += 2 * math.Pi
	}
	for angle >= 2*math.Pi {
		angle -= 2 * math.Pi
	}
	return angle
}

// Vector2D represents a 2D vector
type Vector2D struct {
	X, Y float64
}

// Add adds two vectors
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{X: v.X + other.X, Y: v.Y + other.Y}
}

// Sub subtracts two vectors
func (v Vector2D) Sub(other Vector2D) Vector2D {
	return Vector2D{X: v.X - other.X, Y: v.Y - other.Y}
}

// Mul multiplies vector by scalar
func (v Vector2D) Mul(scalar float64) Vector2D {
	return Vector2D{X: v.X * scalar, Y: v.Y * scalar}
}

// Magnitude returns the magnitude of the vector
func (v Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns a normalized vector
func (v Vector2D) Normalize() Vector2D {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2D{0, 0}
	}
	return Vector2D{X: v.X / mag, Y: v.Y / mag}
}

// Dot returns the dot product
func (v Vector2D) Dot(other Vector2D) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Vector3D represents a 3D vector
type Vector3D struct {
	X, Y, Z float64
}

// Add adds two vectors
func (v Vector3D) Add(other Vector3D) Vector3D {
	return Vector3D{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

// Sub subtracts two vectors
func (v Vector3D) Sub(other Vector3D) Vector3D {
	return Vector3D{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

// Mul multiplies vector by scalar
func (v Vector3D) Mul(scalar float64) Vector3D {
	return Vector3D{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

// Magnitude returns the magnitude of the vector
func (v Vector3D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns a normalized vector
func (v Vector3D) Normalize() Vector3D {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector3D{0, 0, 0}
	}
	return Vector3D{X: v.X / mag, Y: v.Y / mag, Z: v.Z / mag}
}

// Dot returns the dot product
func (v Vector3D) Dot(other Vector3D) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross returns the cross product
func (v Vector3D) Cross(other Vector3D) Vector3D {
	return Vector3D{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

// Statistics provides statistical functions
type Statistics struct {
	values []float64
}

// NewStatistics creates a new statistics calculator
func NewStatistics() *Statistics {
	return &Statistics{
		values: make([]float64, 0),
	}
}

// Add adds a value to the statistics
func (s *Statistics) Add(value float64) {
	s.values = append(s.values, value)
}

// Clear clears all values
func (s *Statistics) Clear() {
	s.values = s.values[:0]
}

// Count returns the number of values
func (s *Statistics) Count() int {
	return len(s.values)
}

// Mean returns the mean of all values
func (s *Statistics) Mean() float64 {
	if len(s.values) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range s.values {
		sum += v
	}
	return sum / float64(len(s.values))
}

// StandardDeviation returns the standard deviation
func (s *Statistics) StandardDeviation() float64 {
	if len(s.values) <= 1 {
		return 0
	}

	mean := s.Mean()
	sumSquares := 0.0
	for _, v := range s.values {
		diff := v - mean
		sumSquares += diff * diff
	}

	return math.Sqrt(sumSquares / float64(len(s.values)-1))
}

// Min returns the minimum value
func (s *Statistics) Min() float64 {
	if len(s.values) == 0 {
		return 0
	}

	min := s.values[0]
	for _, v := range s.values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value
func (s *Statistics) Max() float64 {
	if len(s.values) == 0 {
		return 0
	}

	max := s.values[0]
	for _, v := range s.values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// DeepCopy performs a deep copy of an interface{}
func DeepCopy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	srcValue := reflect.ValueOf(src)
	return deepCopyValue(srcValue).Interface()
}

func deepCopyValue(src reflect.Value) reflect.Value {
	switch src.Kind() {
	case reflect.Ptr:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.New(src.Type().Elem())
		dst.Elem().Set(deepCopyValue(src.Elem()))
		return dst

	case reflect.Slice:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.MakeSlice(src.Type(), src.Len(), src.Cap())
		for i := 0; i < src.Len(); i++ {
			dst.Index(i).Set(deepCopyValue(src.Index(i)))
		}
		return dst

	case reflect.Map:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.MakeMap(src.Type())
		for _, key := range src.MapKeys() {
			dst.SetMapIndex(key, deepCopyValue(src.MapIndex(key)))
		}
		return dst

	case reflect.Struct:
		dst := reflect.New(src.Type()).Elem()
		for i := 0; i < src.NumField(); i++ {
			if dst.Field(i).CanSet() {
				dst.Field(i).Set(deepCopyValue(src.Field(i)))
			}
		}
		return dst

	default:
		return src
	}
}

// Retry executes a function with retry logic
func Retry(fn func() error, maxAttempts int, delay time.Duration) error {
	var lastErr error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if err := fn(); err != nil {
			lastErr = err
			if attempt < maxAttempts {
				time.Sleep(delay)
				delay *= 2 // Exponential backoff
			}
		} else {
			return nil
		}
	}

	return fmt.Errorf("failed after %d attempts: %w", maxAttempts, lastErr)
}
