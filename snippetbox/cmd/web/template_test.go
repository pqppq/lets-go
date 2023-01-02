package main

import (
	"testing"
	"time"
)

func TestHumanData(t *testing.T) {
	tm := time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC)
	hd := humanDate(tm)

	if hd != "17 Mar 2022 at 10:15" {
		t.Errorf("expected %q, but got %q instead.", "17 Mar 2022 at 10:15", hd)
	}
}
