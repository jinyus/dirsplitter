package commands

import "regexp"

var partRe = regexp.MustCompile("part\\d+$")

func isPartDir(path string) bool {
	return partRe.MatchString(path)
}
