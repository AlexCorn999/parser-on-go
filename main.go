package main

import "parser/actions"

func main() {
	// мониторинг необходимых тэгов
	habrTags := []string{"go", "docker", "kubernetes"}

	for _, i := range habrTags {
		actions.HabrGo(i)
	}
}
