package functional

import "fmt"

type Config struct{}

type ConfigModificators func(*Config)

// Here cities number is fixed
// because we cannot pass the number of cities to our function
// The function signature does not permet it
func WithoutCities(c *Config) {
	citiesNumber := 0
	fmt.Println("Cities: ", citiesNumber)
}

// Here we actually decorate our function by return an inner function
func WithCitiesInner(citiesNumber int) ConfigModificators {
	return func(c *Config) {
		fmt.Println("Cities: ", citiesNumber)
	}
}

type Terrain struct {
	config Config
}

func NewTerrain(options ...ConfigModificators) *Terrain {
	var t Terrain
	for _, option := range options {
		option(&t.config)
	}
	return &t
}

func main() {
	t := NewTerrain(
		WithoutCities,
		WithCitiesInner(0),
		WithCitiesInner(6),
	)

	// random line to avoid the compile error
	fmt.Println(t)
}
