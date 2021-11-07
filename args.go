package main

import (
	"log"

	flag "github.com/spf13/pflag"
)

type RunPartsArgs struct {
	Arg []string
	Dir string
	ExitOnError bool
	List bool
	LSBSysInit bool
	Regex string
	Report bool
	Reverse bool
	Test bool
	Umask string
	Verbose bool
}

var Args = &RunPartsArgs{}

func init() {
	flag.StringArrayVarP(&Args.Arg, "arg", "a", []string{},
`pass argument to the scripts.  Use --arg once for each argument you want passed.`)
	flag.BoolVar(&Args.ExitOnError, "exit-on-error", false,
`exit as soon as a script returns with a non-zero exit code.`)
	flag.BoolVar(&Args.List, "list", false,
`print the names of the all matching files (not limited to executables), but don't
actually run them. This option cannot be used with --test.`)
	flag.BoolVar(&Args.LSBSysInit, "lsbsysinit", false,
`filename must be in one or more of either the LANANA-assigned namespace, the LSB
namespaces - either hierarchical or reserved - or the Debian cron script namespace.`)
	flag.StringVar(&Args.Regex, "regex", "",
`validate filenames against custom extended regular expression REGEX.`)
	flag.BoolVar(&Args.Report, "report", false,
`similar to --verbose, but only prints the name of scripts which produce output.
The script's name is printed to whichever of stdout or stderr the script produces
output on. The script's name is not printed to stderr if --verbose also specified.`)
	flag.BoolVar(&Args.Reverse, "reverse", false,
`reverse the scripts' execution order.`)
	flag.BoolVar(&Args.Test, "test", false,
`print the names of the scripts which would be run, but don't actually run them.`)
	flag.StringVar(&Args.Umask, "umask", "022",
`sets the umask to umask before running the scripts. umask should be specified in
octal. By default the umask is set to 022.`)
	flag.BoolVarP(&Args.Verbose, "verbose", "v", false,
`print the name of each script to stderr before running.`)

	flag.Parse()
	if Args.Test && Args.List {
		log.Fatalln("--list and --test cannot be used together")
	}
	if flag.NArg() > 1 {
		flag.Usage()
		log.Fatalln("only one DIRECTORY is expected")
	}
	if flag.NArg() == 1 {
		Args.Dir = flag.Arg(0)
	} else {
		Args.Dir = "."
	}
}
