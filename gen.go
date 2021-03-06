package main

import (
	"bufio"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

func gentemps() *cobra.Command {
	return &cobra.Command{
		Use:   "gen [contest number]",
		Short: "Generate templates. args: contest-number",
		RunE: func(app *cobra.Command, args []string) error {
			generate(args[0])
			writetemp(args[0])
			app.Println("Done!")
			return nil
		},
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generate(str string) {
	problems := "ABCDEF"

	_, err := os.Stat(str)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(str, 0755)
		check(errDir)
	}

	for i := 0; i < 6; i++ {
		file, err := os.Create(str + "/" + string(problems[i]) + ".cpp")
		check(err)
		defer file.Close()

		url := []byte("//https://codeforces.com/contest/" + str + "/problem/" + string(problems[i]) + "\n")

		file.WriteString(string(url))
	}
}

func writetemp(str string) {
	problems := "ABCDEF"

	usr, err := user.Current()
	check(err)

	filename := usr.HomeDir + "/.config/codefgen/config"

	for i := 0; i < 6; i++ {
		file, err := os.Create(str + "/" + string(problems[i]) + ".cpp")
		check(err)
		defer file.Close()

		tempfile, err := os.Open(string(filename))
		check(err)
		defer tempfile.Close()

		scanner := bufio.NewScanner(tempfile)
		for scanner.Scan() {
			file.WriteString(scanner.Text() + "\n")
		}
	}
}
