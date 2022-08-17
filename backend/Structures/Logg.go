package structures

import (
	"fmt"
	"time"
)

type Logg struct {
	Func string
	Time string
}

func NewLog(f string) *Logg {
	l := new(Logg)

	t := time.Now()

	l.Func = f
	l.Time = string(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()))

	return l
}
