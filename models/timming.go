package models

import (
	"time"
)

var t = time.Now()

// Gets now's date and save it to Struct
func Now(Data *Model) *Model {
	t2 := t.Format("2006.01.02")
	Data.NewPath = "Backup-" + t2
	return Data
}

// Define old dates, up to 9 days ago and save it to Struct
func Past(Data *Model) *Model {
	for i := range Data.OldPath {
		t2 := t.Add(time.Duration(i+1) * -24 * time.Hour)
		t3 := t2.Format("2006.01.02")
		Data.OldPath[i] = "Backup-" + t3
	}
	return Data
}
