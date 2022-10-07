package parsinglogfiles

import (
	"fmt"
	"regexp"
)
var (
	reValidLogLine     *regexp.Regexp
	reStrangeSeparator *regexp.Regexp
	rePassword         *regexp.Regexp
	reEndOfLine        *regexp.Regexp
	reUserWithName     *regexp.Regexp
)
func init() {
	reValidLogLine = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	reStrangeSeparator = regexp.MustCompile(`<[~*=-]*>`)
	rePassword = regexp.MustCompile(`(?i)".*password.*"`)
	reEndOfLine = regexp.MustCompile(`end-of-line\d*`)
	reUserWithName = regexp.MustCompile(`User +(\w+)`)
}

func IsValidLine(text string) bool {
	return reValidLogLine.MatchString(text)
}
func SplitLogLine(text string) []string {
	return reStrangeSeparator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if rePassword.MatchString(line) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return reEndOfLine.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	res := make([]string, len(lines))
	copy(res, lines)
	for i, line := range lines {
		name := reUserWithName.FindStringSubmatch(line)
		if len(name) == 2 {
			res[i] = fmt.Sprintf("[USR] %s %s", name[1], line)
		}
	}
	return res
}