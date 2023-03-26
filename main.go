package main

import "parser/actions"

func main() {
	habrTags := []string{"go", "docker", "kubernetes"}
	actions.Gophers()

	for _, i := range habrTags {
		actions.HabrGo(i)
	}
}
