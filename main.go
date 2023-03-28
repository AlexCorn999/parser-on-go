package main

import "parser/actions"

func main() {
	// tags monitoring
	habrTags := []string{"go", "kubernetes"}

	for _, i := range habrTags {
		actions.HabrGo(i)
	}
}
