package toolbox

import (
	"log"
	"math"
	"os"

	"github.com/joho/godotenv"
)

// AlmostEqual compares two floating-point numbers within a small epsilon range.
// It returns true if the absolute difference between 'a' and 'b' is less than or equal to epsilon.
func AlmostEqual(a, b float64) bool {
	epsilon := 0.00000001
	return math.Abs(a-b) <= epsilon
}

// GetEnv retrieves the value of an environment variable identified by 'key'.
// If the environment variable is not set, it attempts to load it from a .env file,
// and if loading fails, it returns the specified 'fallback' value.
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		if err := godotenv.Load(); err == nil {
			value = os.Getenv(key)
		} else {
			log.Printf("Error loading .env file: %s\n", err)
		}
	}

	if value == "" {
		return fallback
	}

	return value
}
