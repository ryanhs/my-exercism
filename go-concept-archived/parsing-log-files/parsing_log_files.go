package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\](.*)`).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~|*|=|-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	c := 0
	r := regexp.MustCompile(`"(.*)((?i)password)(.*)"`)

	for _, v := range lines {
		if r.MatchString(strings.ToLower(v)) {
			c++
		}
	}

	return c
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line(\d*)`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	r := regexp.MustCompile(`User\s+(\S+)`)

	for i, v := range lines {
		finds := r.FindStringSubmatch(v)
		if len(finds) > 0 {
			lines[i] = fmt.Sprintf("[USR] %s %s", finds[1], v)
		}
	}

	return lines
}
