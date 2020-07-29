package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:   "codef-gen",
		Short: "Generated files for Codeforces from the template and copies the test cases",
	}
	app.AddCommand(printabout())
	app.AddCommand(gentemps())
	app.AddCommand(settemp())
	err := app.Execute()
	if err != nil {
		os.Exit(1)
	}
}
