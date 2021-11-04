package main

import (
	"os"
	"regexp"
	"strings"
)

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
