package main

import (
	"testing"
	"time"

	"github.com/pqppq/lets-go/snippetbox/internal/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		tm       time.Time
		expected string
	}{
		{
			name:     "UTC",
			tm:       time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			expected: "17 Mar 2022 at 10:15",
		},
		{
			name:     "Empty",
			tm:       time.Time{},
			expected: "",
		},
		{
			name:     "CET",
			tm:       time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			expected: "17 Mar 2022 at 09:15",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hd := humanDate(tc.tm)
			if hd != tc.expected {
				assert.Equal(t, hd, tc.expected)
			}
		})
	}
}
