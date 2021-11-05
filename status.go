package main

type Status struct {
	ExitCode int;
}

func NewStatus() *Status {
	return &Status{}
}

func (s *Status) Reset() {
	s.ExitCode = 0
}
