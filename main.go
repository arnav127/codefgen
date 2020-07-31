package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:   "Codeforces File Generator",
		Short: "Generated files for Codeforces from the template and copies the test cases",
	}
	app.AddCommand(printabout())
	err := app.Execute()
	if err != nil {
		os.Exit(1)
	}
}
