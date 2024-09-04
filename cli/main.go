package main

import (
	"github.com/spf13/cobra"
)

func main() {
	(&cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
		},
	}).Execute()
}
