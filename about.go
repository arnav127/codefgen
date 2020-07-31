package main

import (
	"github.com/spf13/cobra"
)

func printabout() *cobra.Command {
	return &cobra.Command{
		Use:   "about",
		Short: "Print details about the author",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("This software is developed by Arnav Dixit \nhttps://github.com/arnav127")
			return nil
		},
	}
}
