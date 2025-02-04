package common

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "0.1.3"

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show agent version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version)
		},
	}
}
