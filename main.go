package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

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
var dir = "."

func FindFiles(dir string) ([]os.FileInfo, error) {
	var files, err = ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	if *reverse {
		for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
			files[i], files[j] = files[j], files[i]
		}
	}
	return files, nil
}

func standardSuffixToIgnore() []string {
	return []string{
		"~",
		",",
		".disabled",
		".cfsaved",
		".rpmsave",
		".rpmorig",
		".rpmnew",
		".swp",
		",v",
	}
}

func lsbSysInitSuffixToIgnore() []string {
	return []string{".dpkg-old", ".dpkg-dist", ".dpkg-new", ".dpkg-tmp"}
}

func lsbSysInitRegexToAccept() []regexp.Regexp {
	var r = []string{
		`^[a-z0-9]+$`,                  // LANANA-assigned LSB hierarchical
		`^_?([a-z0-9_.]+-)+[a-z0-9]+$`, // LANANA-assigned LSB reserved
		`^[a-zA-Z0-9_-]+$`,             // Debian cron script namespaces
	}
	var result = []regexp.Regexp{}
	for _, v := range r {
		result = append(result, *regexp.MustCompile(v))
	}
	return result
}

func FilterFileName(fileName string) (bool, error) {
	for _, s := range standardSuffixToIgnore() {
		if strings.HasSuffix(fileName, s) {
			return false, nil
		}
	}
	if *lsbsysinit {
		for _, s := range lsbSysInitSuffixToIgnore() {
			if strings.HasSuffix(fileName, s) {
				return false, nil
			}
		}
		for _, r := range lsbSysInitRegexToAccept() {
			if r.MatchString(fileName) {
				return true, nil
			}
		}
		return false, nil
	}
	if regex != nil {
		var matched, err = regexp.MatchString(*regex, fileName)
		return matched, err
	}
	return true, nil
}

func FilterFile(file os.FileInfo) (bool, error) {
	if file.IsDir() {
		return false, nil
	}
	return FilterFileName(file.Name())
}

func Run() error {
	var files, err = FindFiles(dir)
	if err != nil {
		log.Fatalln(err)
	}
	var filesToProcess = []os.FileInfo{}
	for _, file := range files {
		var include, err = FilterFile(file)
		if err != nil {
			return err
		}
		if include {
			filesToProcess = append(filesToProcess, file)
		}
	}
	for _, file := range filesToProcess {
		log.Printf("Found %s", file.Name())
	}
	return nil
}

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
	if flag.NArg() > 1 {
		flag.Usage()
		log.Fatalln("only one DIRECTORY is expected")
	}
	if flag.NArg() == 1 {
		dir = flag.Arg(0)
	}
	Run()
}
