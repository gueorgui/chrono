package chronolib

import (
	"errors"
	"github.com/jinzhu/now"
	"sort"
	"time"
)

// SimpleFrame is used for converting to and from JSON
type SimpleFrame struct {
	Project   string
	StartedAt string
	EndedAt   string
	Tags      []string
	Notes     []string
}

// Data is a wrapper for frames
type Data struct {
	Frames []Frame
}

// GetFrameByShortHex gets a frame using the short form of its UUID
func GetFrameByShortHex(frames []Frame, hex string) (int, Frame, error) {
	for idx, frame := range frames {
		if GetShortHex(frame.UUID) == hex {
			return idx, frame, nil
		}
	}
	return 0, Frame{}, errors.New("No frame found")
}

// SortFramesByDate sorts frame by its start date
func SortFramesByDate(frames []Frame) {
	sort.Slice(frames, func(i, j int) bool {
		return frames[i].UpdatedAt.Before(frames[j].StartedAt)
	})
}

// SortTimeMapKeys sorts a timemap by its key (a time.Time)
func SortTimeMapKeys(timemap *map[time.Time][]Frame) []time.Time {
	var keys []time.Time
	for k := range *timemap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})
	return keys
}

// ConvertFrameToSimpleFrame converts a frame to a raw frame
func ConvertFrameToSimpleFrame(frame Frame) SimpleFrame {
	return SimpleFrame{
		frame.Project,
		frame.StartedAt.Format("2006-01-02 15:04:05"),
		frame.EndedAt.Format("2006-01-02 15:04:05"),
		frame.Tags,
		frame.Notes,
	}
}

// ConvertSimpleFrameToFrame converts a raw frame back to a frame
func ConvertSimpleFrameToFrame(uuid string, rawFrame SimpleFrame) (Frame, error) {
	started, err := now.Parse(rawFrame.StartedAt)
	if err != nil {
		return Frame{}, err
	}
	ended, err := now.Parse(rawFrame.EndedAt)
	if err != nil {
		return Frame{}, err
	}
	return Frame{
		uuid,
		rawFrame.Project,
		started,
		ended,
		time.Now(),
		rawFrame.Tags,
		rawFrame.Notes,
	}, nil
}

// ContainsMoreThanOneBooleanFlag is a helper method for checking if more than one boolean is true
func ContainsMoreThanOneBooleanFlag(flags ...bool) bool {
	count := 0
	for _, flag := range flags {
		if flag {
			count++
		}
		if count == 2 {
			return true
		}
	}
	return false
}
