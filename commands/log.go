package commands

import (
    "os"
    "fmt"
    "strings"
    "time"
    "github.com/jordanknott/chrono/chronolib"
    "github.com/spf13/cobra"
)

var forCurrentWeek bool
var forCurrentMonth bool
var forCurrentYear bool

func ContainsMoreThanOneBooleanFlag(flags... bool) bool {
    count := 0
    for _, flag := range flags {
        if flag {
            count += 1
        }
        if count == 2 {
            return true
        }
    }
    return false
}

func IsFrameInTimespan(frame chronolib.Frame, start time.Time, end time.Time) bool {
    if !chronolib.IsTimeInTimespan(frame.StartedAt, start, end) {
        return false
    }
    if !chronolib.IsTimeInTimespan(frame.EndedAt, start, end) {
        return false
    }
    return true
}

func PrettyDate(t *time.Time) string {
    return t.Format("_2 January 2006 15:02")
}

func FilterFramesByTimespan(start time.Time, end time.Time, frames *[]chronolib.Frame) map[time.Time][]chronolib.Frame {
    filteredFrames := make(map[time.Time][]chronolib.Frame)
    for _, frame := range *frames {
        if IsFrameInTimespan(frame, start, end) {
            date := chronolib.NormalizeDate(frame.StartedAt)
            filteredFrames[date] = append(filteredFrames[date], frame)
        }
    }
    return filteredFrames
}

func newLogCmd() *cobra.Command {
    cmd :=  &cobra.Command{
        Use: "log",
        Short: "Get the frames logged for a single day",
        Long: "Get the frames logged for a single day",
        Run: func(cmd *cobra.Command, args []string) {
            framesPath := chronolib.GetAppFilePath("frames", "")
            data := chronolib.LoadFrames(framesPath)

            if ContainsMoreThanOneBooleanFlag(forCurrentWeek, forCurrentMonth, forCurrentYear) {
                fmt.Println("Error: the folllowing flags are mutually exclusive: ['--week', '--year', '--month']")
                os.Exit(0)
            }

            var tsStart, tsEnd time.Time

            if forCurrentWeek {
                tsStart, tsEnd = chronolib.GetTimespanForWeek()

            } else if forCurrentMonth {
                tsStart, tsEnd = chronolib.GetTimespanForMonth()
            } else if forCurrentYear {
                tsStart, tsEnd = chronolib.GetTimespanForYear()
            } else {
                tsStart, tsEnd = chronolib.GetTimespanForToday()
            }
            filteredFrames := FilterFramesByTimespan(tsStart, tsEnd, &data.Frames)
            dates := chronolib.SortTimeMapKeys(&filteredFrames)
            for _, date := range dates {
                fmt.Println(date.Format("Monday _2 January 2006"))
                for _, frame := range filteredFrames[date] {
                    start := frame.StartedAt
                    end := frame.EndedAt
                    hours, minutes, seconds := chronolib.GetTimeElapsed(start, end)
                    if len(frame.Tags) == 0 {
                        fmt.Printf("\t(ID: %s) %s to %s %4dh %02dm %02ds  %-12s\n", chronolib.GetShortHex(frame.UUID), start.Format("15:04"), end.Format("15:04"), hours, minutes, seconds, frame.Project)
                    } else {
                        fmt.Printf("\t(ID: %s) %s to %s %4dh %02dm %02ds  %-12s [%s]\n",
                        chronolib.GetShortHex(frame.UUID), start.Format("15:04"), end.Format("15:04"),
                        hours, minutes, seconds, frame.Project, strings.Join(frame.Tags, ", "))
                    }
                }
            }
        },
    }
    cmd.Flags().BoolVarP(&forCurrentWeek, "week", "w", false, "show frames for entire week")
    cmd.Flags().BoolVarP(&forCurrentMonth, "month", "m", false, "show frames for entire month")
    cmd.Flags().BoolVarP(&forCurrentYear, "year", "y", false, "show frames for entire year")
    return cmd
}
