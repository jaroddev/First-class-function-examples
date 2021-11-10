package oop

import "fmt"

type Config struct{}

// Same idea than in terrain-first but with interfaces
type Option interface {
	Apply(*Config)
}

type Terrain struct {
	config Config
}

func NewTerrain(options ...Option) *Terrain {
	var config Config
	for _, option := range options {
		option.Apply(&config)
	}

	t := Terrain{
		config,
	}

	return &t
}

type cities struct {
	cities int
}

func WithCities(n int) Option {
	return &cities{
		cities: n,
	}
}

func (cities *cities) Apply(config *Config) {
	fmt.Println("Cities: ", cities.cities)
}

func main() {
	t := NewTerrain(
		WithCities(0),
		WithCities(6),
	)

	// random line to avoid the compile error
	fmt.Println(t)
}
