package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"

	"github.com/spf13/cobra"
)

func settemp() *cobra.Command {
	return &cobra.Command{
		Use:   "stemp",
		Short: "Set template file",
		RunE: func(app *cobra.Command, args []string) error {
			newtemplate()
			return nil
		},
	}
}

func newtemplate() {
	fmt.Println("Enter new file path for template (absolute path): ")
	var path string
	fmt.Scanln(&path)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("The specified file does not exist. Exiting")
		return
	}
	usr, err := user.Current()
	check(err)

	filename := usr.HomeDir + "/.config/codefgen/config"
	script := usr.HomeDir + "/.config/codefgen/stemp.sh"
	cmd := exec.Command("/bin/sh", script, path, filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	errr := cmd.Run()
	if errr != nil {
		fmt.Println(errr)
	}
}
