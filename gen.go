package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func gentemps() *cobra.Command {
	return &cobra.Command{
		Use:   "gen",
		Short: "Generate templates. args: round-number",
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
	// err := ioutil.WriteFile("/tmp/dat1", url, 0644)
	problems := "ABCDEF"
	// dirPath, _ := os.Getwd()
	_, err := os.Stat(str)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(str, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
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

	filename, err := ioutil.ReadFile("~/.config/codef-gen/config")
	check(err)
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
