package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func FindFiles() ([]os.FileInfo, error) {
	var files, err = ioutil.ReadDir(Args.Dir)
	if err != nil {
		return nil, err
	}
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	if Args.Reverse {
		for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
			files[i], files[j] = files[j], files[i]
		}
	}
	return files, nil
}

func ActOnFile(file os.FileInfo, status *Status) (err error) {
	command := Args.Dir + "/" + file.Name()
	if Args.ExitOnError && status.ExitCode != 0 {
		return
	}
	status.Reset()
	if Args.List {
		log.Printf("%v %v", command, Args.Arg)
		return
	}
	if file.Mode() & 0111 == 0 {
		return
	}
	if Args.Test {
		log.Printf("%v %v", command, Args.Arg)
		return
	}
	// TODO - implement random sleep
	if Args.Verbose {
		log.Printf("executing %v %v", command, Args.Arg)
	}
	// TODO - implement umask
	Exec(command, status)
	if (Args.Report || Args.Verbose) && status.ExitCode != 0 {
		log.Printf("%v exited with return code %v", command, status.ExitCode)
	}
	return
}

func Run() (*Status, error) {
	var files, err = FindFiles()
	if err != nil {
		return nil, err
	}
	var filesToProcess = []os.FileInfo{}
	for _, file := range files {
		var include, err = FilterFile(file)
		if err != nil {
			return nil, err
		}
		if include {
			filesToProcess = append(filesToProcess, file)
		}
	}
	var status = NewStatus()
	for _, file := range filesToProcess {
		ActOnFile(file, status)
	}
	return status, nil
}

func main() {
	var status, err = Run()
	if err == nil {
		os.Exit(status.ExitCode)
	} else {
		log.Fatalln(err)
	}
}
