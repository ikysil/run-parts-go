package main

import (
	"os"
	"sync"
)

type Report struct {
	mu sync.Mutex
	reportString string
	report bool
	verbose bool
	used bool
}

func NewReport(file os.FileInfo, report bool, verbose bool) *Report {
	return &Report{reportString: file.Name(), report: report, verbose: verbose}
}

func (r *Report) getReport(condition bool) (report string) {
	report = ""
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.used {
		return
	}
	if condition {
		report = r.reportString
	}
	return
}

func (r *Report) OutReport() (report string) {
	return r.getReport(r.report)
}

func (r *Report) ErrReport() (report string) {
	return r.getReport(r.report && !r.verbose)
}
