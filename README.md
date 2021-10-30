# run-parts-go
Go implementation of `run-parts` - concept taken from [Debian](http://www.unix.com/man-page/linux/8/run-parts/) with extensions from [Ubuntu](http://manpages.ubuntu.com/manpages/trusty/man8/run-parts.8.html)

# Usage

    NAME
           run-parts - run scripts or programs in a directory

    SYNOPSIS
           run-parts  [--test]  [--verbose]  [--report]  [--umask=umask] [--lsbsysinit] [--regex=REGEX]
           [--arg=argument] [--exit-on-error] [--help] [--list] [--reverse]  [--]  DIRECTORY

    DESCRIPTION
           run-parts runs all the executable files named within constraints described below, found in
           directory directory.  Other files and directories are silently ignored.

           If neither the --lsbsysinit option nor the --regex option is given  then  the  names  must
           consist  entirely of ASCII upper- and lower-case letters, ASCII digits, ASCII underscores,
           and ASCII minus-hyphens.

           If the --lsbsysinit option is given,  then  the  names  must  not  end  in  .dpkg-old   or
           .dpkg-dist  or  .dpkg-new  or  .dpkg-tmp,  and must belong to one or more of the following
           namespaces: the LANANA-assigned namespace (^[a-z0-9]+$); the LSB hierarchical and reserved
           namespaces  (^_?([a-z0-9_.]+-)+[a-z0-9]+$);  and the Debian cron script namespace (^[a-zA-
           Z0-9_-]+$).

           If the --regex option  is  given,  the  names  must  match  the  custom  extended  regular
           expression specified as that option's argument.

           Files  are  run	in  the  lexical sort order of their names unless the --reverse option is
           given, in which case they are run in the opposite order.

    OPTIONS
           --test print the names of the scripts which would be run, but don't actually run them.

           --list print the names of the all matching files (not limited to executables),  but  don't
              actually run them. This option cannot be used with --test.

           -v, --verbose
              print the name of each script to stderr before running.

           --report
              similar  to  --verbose,  but  only prints the name of scripts which produce output.
              The script's name is printed to whichever of stdout or stderr the script produces
              output on. The script's name is not printed to stderr if --verbose also specified.

           --reverse
              reverse the scripts' execution order.

           --exit-on-error
              exit as soon as a script returns with a non-zero exit code.

           --umask=umask
              sets  the  umask to umask before running the scripts.  umask should be specified in
              octal.  By default the umask is set to 022.

           --lsbsysinit
              filename must be in one or more of either the LANANA-assigned namespace, the LSB
              namespaces - either hierarchical or reserved - or the Debian cron script namespace.

           --regex=REGEX
              validate filenames against custom extended regular expression REGEX

           -a, --arg=argument
              pass argument to the scripts.  Use --arg once for each argument you want passed.

           --     specifies that this is the end of the options.  Any filename after -- will  be  not
              be interpreted as an option even if it starts with a hyphen.

           -h, --help
              display usage information and exit.

# Useful Links

* [Building simple command-line (CLI) applications in Go using Commando](https://medium.com/sysf/building-simple-command-line-cli-applications-in-go-using-commando-8a8e0edbd48a)
* [How to write fast, fun command-line applications with Golang](https://www.freecodecamp.org/news/writing-command-line-applications-in-go-2bc8c0ace79d/)
* [pflag is a drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.](https://github.com/spf13/pflags)