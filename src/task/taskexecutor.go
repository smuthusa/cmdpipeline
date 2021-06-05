package task

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

type Executor func(task string, description string, level int) error

func Execute(configFile string, executor Executor) {
	config := readConfiguration(configFile)

	for _, step := range config.Config.Step {
		for _, stepDef := range step.Steps {
			wg := &sync.WaitGroup{}
			wg.Add(1)
			ExecuteStep(stepDef, wg, executor, 0)
			wg.Wait()
		}
	}
}

func readConfiguration(configFile string) Configuration {
	var config Configuration
	data, err := ioutil.ReadFile(configFile)
	handleError(err)

	err = yaml.Unmarshal(data, &config)
	handleError(err)
	return config
}

func ExecuteStep(step Step, parentWg *sync.WaitGroup, executor Executor, level int) {
	hyphen := ""
	for i := 0; i < level; i++ {
		hyphen = hyphen + "====="
	}
	fmt.Println(fmt.Sprintf("+%s==== [ %s ] Execution of step %s started!", hyphen, time.Now().Format("15:04:05"), step.Name))
	defer parentWg.Done()
	executeCommandsAndWaitForCompletion(step.Commands, level, executor)
	executeStepsAndWaitForCompletion(step.Steps, level, executor)
	fmt.Println(fmt.Sprintf("+%s==== [ %s ] Execution of step %s completed!", hyphen, time.Now().Format("15:04:05"), step.Name))
}

func executeStepsAndWaitForCompletion(steps []Step, level int, executor Executor) {
	stepWg := &sync.WaitGroup{}
	for _, stepG := range steps {
		stepWg.Add(1)
		groupLevel := level + 1
		go func(stepLocal Step, pwg *sync.WaitGroup, lvl int) {
			pwg.Add(1)
			ExecuteStep(stepLocal, pwg, executor, lvl)
			pwg.Done()
		}(stepG, stepWg, groupLevel)
	}
	stepWg.Wait()
}

func executeCommandsAndWaitForCompletion(commands []Command, level int, executor Executor) {
	wg := &sync.WaitGroup{}
	for _, command := range commands {
		wg.Add(1)
		go func(cmd string, desc string, taskExecutor Executor) {
			taskExecutor(cmd, desc, level)
			wg.Done()
		}(command.Name, command.Command, executor)
	}
	wg.Wait()
}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
