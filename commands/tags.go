package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
)

func newTagsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tags",
		Short: "Get a list of all tags used",
		Long:  "Get a list of all tags used",
		Run: func(cmd *cobra.Command, args []string) {
			frameStorage := chronolib.GetFrameStorage()
			tags, err := frameStorage.Tags()
			if err != nil {
				commandError = err
				return
			}

			for _, tag := range tags {
				fmt.Println(tag)
			}
		},
	}
}
