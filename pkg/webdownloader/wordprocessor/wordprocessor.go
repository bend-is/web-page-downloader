package wordprocessor

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type WordProcessor struct {
	splitter *regexp.Regexp
}

func New() *WordProcessor {
	return &WordProcessor{
		splitter: regexp.MustCompile(`[\s.,!?"';:\[\]()\n\r\t<>/=+#_\-\\{}|^№&$*]`),
	}
}

func (wp *WordProcessor) UniqueWordsCount(reader io.Reader) map[string]int {
	m := make(map[string]int)

	for scanner := bufio.NewScanner(reader); scanner.Scan(); {
		for _, s := range wp.splitter.Split(scanner.Text(), -1) {
			if s == "" || utf8.RuneCountInString(s) == 1 {
				continue
			}
			if _, err := strconv.Atoi(s); err == nil {
				continue
			}
			s = strings.ToUpper(s)
			if _, exist := m[s]; exist {
				m[s]++
				continue
			}
			m[s] = 1
		}
	}

	return m
}
