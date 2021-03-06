package main

import (
	"io"
	"os"
	"os/exec"
	"syscall"
)

type ReportingWriter struct {
	w io.Writer
	r func () string
}

var _ io.Writer = (*ReportingWriter) (nil)

func NewReportingWriter(w io.Writer, r func () string) *ReportingWriter {
	return &ReportingWriter{w, r}
}

func (w *ReportingWriter) Write(d []byte) (int, error) {
	var r = w.r();
	if r != "" {
		w.w.Write([]byte(r + ":\n"))
	}
    return w.w.Write(d)
}

func Exec(command string, status *Status) (err error) {
	oldUmask := syscall.Umask(Args.Umask)
	defer syscall.Umask(oldUmask)

	var report = NewReport(command, Args.Report, Args.Verbose);
	var cmd = exec.Command(command, Args.Arg...);
	cmd.Stderr = NewReportingWriter(os.Stderr, report.ErrReport)
	cmd.Stdout = NewReportingWriter(os.Stdout, report.OutReport)
	err = cmd.Start()
	if err != nil {
		return
	}
	err = cmd.Wait()
	status.ExitCode = cmd.ProcessState.ExitCode()
	return
}
