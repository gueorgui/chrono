package commands

import (
	"fmt"
	"github.com/gochrono/chrono/chronolib"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func newNotesCmd() *cobra.Command {
	notes := &cobra.Command{
		Use:   "notes",
		Short: "Manage notes for the current frame",
		Long: `Manage notes for the current frame. You can add new notes with the 'add' command
		or remove them with the 'delete' command. You can list all added notes with the
		'show' command.`,
	}
	notes.AddCommand(newNotesAddCmd(), newNotesShowCmd(), newNotesDeleteCmd())
	return notes
}

func newNotesAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [message]",
		Short: "Add a new note to current frame",
		Long:  "Add a new note to current frame",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			state, _ := chronolib.GetState(config)
			if state.IsEmpty() {
				fmt.Println(chronolib.FormatNoProjectMessage())
			} else {
				currentFrame := state.Get()
				currentFrame.Notes = append(currentFrame.Notes, args[0])
				state.Update(currentFrame)
				chronolib.SaveState(config, state)
			}
		},
	}
}
func newNotesShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show all notes for the current frame",
		Long:  "Show all notes for the current frame",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			state, _ := chronolib.GetState(config)
			if state.IsEmpty() {
				fmt.Println(chronolib.FormatNoProjectMessage())
			} else {
				notes := state.Get().Notes
				if len(notes) != 0 {
					for index, note := range notes {
						fmt.Println(chronolib.FormatNoteShowLine(index, note))
					}
				} else {
					fmt.Println(chronolib.FormatNoNotesMessage())
				}
			}
		},
	}
}

func newNotesDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a note from the current frame",
		Long:  "Delete a note from the current frame",
		Run: func(cmd *cobra.Command, args []string) {
			configDir := chronolib.GetCorrectConfigDirectory("")
			config := chronolib.GetConfig(configDir)
			state, _ := chronolib.GetState(config)
			if state.IsEmpty() {
				fmt.Println(chronolib.FormatNoProjectMessage())
			} else {
				currentFrame := state.Get()
				index, err := strconv.Atoi(args[0])
				if err != nil {
					fmt.Println("Index must be a number!")
					os.Exit(0)
				} else if index > len(currentFrame.Notes) || index < 0 {
					fmt.Println("Index must be a number!")
					os.Exit(0)
				}
				fmt.Printf("Deleting note '%s'\n", currentFrame.Notes[index])
				currentFrame.Notes = append(currentFrame.Notes[:index], currentFrame.Notes[index+1:]...)
				state.Update(currentFrame)
				_ = chronolib.SaveState(config, state)
			}
		},
	}
}
