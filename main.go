package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:   "codefgen",
		Short: "Generated files for Codeforces from given template",
	}
	app.AddCommand(printabout())
	app.AddCommand(gentemps())
	app.AddCommand(settemp())
	err := app.Execute()
	if err != nil {
		os.Exit(1)
	}
}
