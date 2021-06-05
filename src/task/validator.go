package task

import (
	"fmt"
	"log"
)

func validate(stages []Stage, stepMap map[string]Step) bool {
	for _, stage := range stages {
		for _, stepName := range stage.Steps {
			if _, ok := stepMap[stepName]; !ok {
				log.Println(fmt.Sprintf("Undefined step %s", stepName))
				return false
			}
		}
	}
	return true
}
