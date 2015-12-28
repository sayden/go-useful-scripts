package main
import (
	"os/exec"
	"fmt"
	"os"
)

func main(){
	args := os.Args[1:]
	if len(args) != 2 {
		panic("You must provide 2 arguments")
	}

	cmdName := "git"
	cmdArgs := []string{"config", "--local", "--add", "user.name", args[0]}

	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		panic(err)
	}

	cmdArgs[3] = "user.email"
	cmdArgs[4] = args[1]

	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		panic(err)
	}

	fmt.Println("User and email set correctly")
}