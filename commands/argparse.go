package commands

import (
	"errors"
	"github.com/gochrono/chrono/chronolib"
	"github.com/jinzhu/now"
	jww "github.com/spf13/jwalterweatherman"
	"strconv"
	"time"
)

// ParseStartArgs properly handles user arguments for the start command
func ParseStartArgs(args []string, startAt string, startEnded string, startNote string) (chronolib.CurrentFrame, time.Time, error) {
	project, frameTags, err := ParseStartArguments(args)
	if err != nil {
		return chronolib.CurrentFrame{}, time.Time{}, err
	}

	frameStart, err := ParseTime(startAt)
	if err != nil {
		return chronolib.CurrentFrame{}, time.Time{}, NewErrTimeStringNotValid(startAt)
	}
	var frameEnd time.Time
	if chronolib.IsKeyword(startEnded) {
		frameEnd, err = chronolib.CompileKeyword(startEnded)
	} else {
		frameEnd, err = ParseTime(startEnded)
	}
	if err != nil {
		return chronolib.CurrentFrame{}, time.Time{}, NewErrTimeStringNotValid(startEnded)
	}

	notes := []string{}
	if startNote != "" {
		notes = append(notes, startNote)
	}

	return chronolib.CurrentFrame{
		Project:   project,
		StartedAt: frameStart,
		UpdatedAt: time.Now(),
		Tags:      frameTags,
		Notes:     notes,
	}, frameEnd, nil

}

// ParseStartArguments splits the argument string list and validates tags
func ParseStartArguments(args []string) (string, []string, error) {
	project := args[0]
	tags := args[1:]

	if err := chronolib.CheckTags(tags); err != nil {
		return "", []string{}, NewErrTagNotValid("invalid tag")
	}

	return project, chronolib.NormalizeTags(tags), nil
}

// ParseTime converts a properly formated time string into a time.Time struct
func ParseTime(t string) (time.Time, error) {
	if t == "" {
		jww.INFO.Printf("time is empty, using time.Now()")
		return time.Now(), nil
	}
	parsedTime, err := now.Parse(t)
	if err != nil {
		jww.ERROR.Printf("error parsing time: %s", t)
		return time.Time{}, errors.New("invalid time format: " + t)
	}
	jww.DEBUG.Printf("using custom time %v", parsedTime)
	return parsedTime, nil
}

// TimespanFlags is a struct containing the four different options for timespans
type TimespanFlags struct {
	AllTime      bool
	Day          string
	Week         string
	Month        string
	Year         string
}

// ParseTimespanFlags gets the correct start and end time for filtering frames based
// on the flags given
func ParseTimespanFlags(timespanFlags TimespanFlags) chronolib.TimespanFilterOptions {
	var tsStart, tsEnd time.Time
	if timespanFlags.AllTime {
		return chronolib.TimespanFilterOptions{}
	} else if timespanFlags.Week != "" {
		tsStart, tsEnd = chronolib.GetTimespanForWeek(timespanFlags.Week)
	} else if timespanFlags.Month != "" {
		tsStart, tsEnd = chronolib.GetTimespanForMonth(timespanFlags.Month)
	} else if timespanFlags.Year != "" {
		tsStart, tsEnd = chronolib.GetTimespanForYear(timespanFlags.Year)
	} else if timespanFlags.Day != "" {
		tsStart, tsEnd = chronolib.GetTimespanForDay(timespanFlags.Day)
	} else {
		tsStart, tsEnd = chronolib.GetTimespanForToday()
	}
	return chronolib.TimespanFilterOptions{Start: tsStart, End: tsEnd}
}

// GetFrame is a helper method for getting a frame by either index or UUID
func GetFrame(frames chronolib.Frames, target string) (chronolib.Frame, bool) {
	index, err := strconv.Atoi(target)
	if err == nil {
		frame, ok := frames.GetByIndex(index)
		if ok {
			return frame, true
		}
	}
	return frames.GetByUUID(target)
}
