package main

import "log"

func init() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("run-parts: ")
	log.SetFlags(0)
}
