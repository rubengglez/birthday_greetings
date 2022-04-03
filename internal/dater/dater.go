package dater

import "time"

type Date interface {
	Month() int
	Day() int
}

type date struct{}

func (d *date) Month() int {
	return int(time.Now().UTC().Month())
}

func (d *date) Day() int {
	return int(time.Now().UTC().Day())
}

func NewDefaultDater() Date {
	return &date{}
}
