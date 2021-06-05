package handlers

import (
	"fmt"
	"github.com/smuthusa/taskexecutor/src/process"
	"strings"
	"time"
)

func ConsoleLogger(task string, desc string, level int) error {
	hyphen := ""
	for i := 0; i <= level+1; i++ {
		hyphen = hyphen + "+----"
	}
	fmt.Println(fmt.Sprintf("%s [ %s ] Starting execution of command %s, %s", hyphen, time.Now().Format("15:04:05"), task, desc))
	fmt.Println(fmt.Sprintf("%s [ %s ] Execution of command %s, %s completed.", hyphen, time.Now().Format("15:04:05"), task, desc))
	return nil
}

func CommandExecutor(task string, desc string, level int) error {
	hyphen := ""
	for i := 0; i <= level+1; i++ {
		hyphen = hyphen + "+----"
	}
	fmt.Println(fmt.Sprintf("%s [ %s ] Starting execution of command %s, %s", hyphen, time.Now().Format("15:04:05"), task, desc))
	args := strings.Split(task, " ")
	err := process.Execute(args[0], desc, args[1:])
	fmt.Println(fmt.Sprintf("%s [ %s ] Execution of command %s, %s completed.", hyphen, time.Now().Format("15:04:05"), task, desc))
	return err
}
