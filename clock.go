package main

import "time"

var clock Clock

func init() {
	clock = RealClock{}
}

type Clock interface {
	Now() time.Time
}

type RealClock struct {}

func (clock RealClock) Now() time.Time {
	return time.Now()
}
