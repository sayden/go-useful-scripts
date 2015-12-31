package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// AddLocalUser will add in local git config the provided user name and email
// It must be executed in the same directory of the git repo
func AddLocalUser(user, email, path *string) {
	cmdName := "git"
	cmdArgs := []string{"config", "--local", "--add", "user.name", *user}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Error getting current directory: %s", err)
	}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Path = fmt.Sprintf("%s/%s", dir, *path)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error adding username: %s", err)
	}

	cmdArgs[3] = "user.email"
	cmdArgs[4] = *email

	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Println("Error adding email info")
		panic(err)
	}

	fmt.Println("User and email set correctly")
}

//Clone clones the specified repo in the executing directory
func Clone(repo *string) {
	cmdName := "git"
	cmdArgs := []string{"clone", *repo}

	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		panic(err)
	}

	fmt.Printf("Repo %s cloned succesfully\n", *repo)
}
