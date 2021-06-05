package process

import (
	"fmt"
	"log"
	"os/exec"
)

func Execute(cmd string, description string, args []string) error {
	log.Println(fmt.Sprintf("Executing command %s, %s ", cmd, description))

	command := exec.Command(cmd, args...)
	out, err := command.CombinedOutput()
	log.Println(string(out))
	if err == nil {
		log.Println("Command execution successfully completed")
	} else {
		log.Println(fmt.Sprintf("Command execution failed with error %s", err.Error()))
	}
	return err
}
