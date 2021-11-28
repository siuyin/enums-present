// Package compass provides direction finding routines.
package compass

// Dir is a compass direction
type Dir int

//go:generate stringer -type=Dir
// install stringer with: go install golang.org/x/tools/cmd/stringer@latest

// Compass directions.
const (
	North Dir = iota
	East
	South
	West
)
