package main

import (
	"log"

	flag "github.com/spf13/pflag"
)

var arg = flag.StringArrayP("arg", "a", []string{},
	`pass argument to the scripts.  Use --arg once for each argument you want passed.`)
var exit_on_error = flag.Bool("exit-on-error", false,
	`exit as soon as a script returns with a non-zero exit code.`)
var list = flag.Bool("list", false,
	`print the names of the all matching files (not limited to executables), but don't
actually run them. This option cannot be used with --test.`)
var lsbsysinit = flag.Bool("lsbsysinit", false,
	`filename must be in one or more of either the LANANA-assigned namespace, the LSB
namespaces - either hierarchical or reserved - or the Debian cron script namespace.`)
var regex = flag.String("regex", "",
	`validate filenames against custom extended regular expression REGEX.`)
var report = flag.Bool("report", false,
	`similar to --verbose, but only prints the name of scripts which produce output. 
The script's name is printed to whichever of stdout or stderr the script produces
output on. The script's name is not printed to stderr if --verbose also specified.`)
var reverse = flag.Bool("reverse", false,
	`reverse the scripts' execution order.`)
var test = flag.Bool("test", false,
	`print the names of the scripts which would be run, but don't actually run them.`)
var umask = flag.String("umask", "022",
	`sets the umask to umask before running the scripts. umask should be specified in
octal. By default the umask is set to 022.`)
var verbose = flag.BoolP("verbose", "v", false,
	`print the name of each script to stderr before running.`)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("run-parts-go: ")
	log.SetFlags(0)
	flag.Parse()
	if *test && *list {
		log.Fatalln("--list and --test cannot be used together")
	}
	log.Println("Hello, World!")
}
