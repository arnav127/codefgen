package main

import (
	"github.com/spf13/cobra"
)

func printabout() *cobra.Command {
	return &cobra.Command{
		Use:   "about",
		Short: "Print details about the author",
		RunE: func(app *cobra.Command, args []string) error {
			app.Println("This software is developed by Arnav Dixit \nhttps://github.com/arnav127")
			return nil
		},
	}
}
