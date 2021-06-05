package task

import "log"

func parseSteps(config Configuration) map[string]Step {
	stepMap := map[string]Step{}
	for _, step := range config.Config.Step {
		if _, ok := stepMap[step.Name]; ok {
			log.Fatalf("Duplicate step %s", step.Name)
		}
		stepMap[step.Name] = step
	}
	return stepMap
}
