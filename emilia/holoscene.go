package emilia

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	RFC_EMILY = "Mon, 02 Jan 2006"
)

var (
	HEregex          = regexp.MustCompile(`(\d+);\s*(\d+)\s*H.E.`)
	HEParagraphRegex = regexp.MustCompile(`>(\d+;\s*\d+\s*H.E.)`)
)

// ConvertHoloscene takes a Holoscene time (127; 12022 H.E.) to a time struct.
func ConvertHoloscene(HEtime string) *time.Time {
	matches := HEregex.FindAllStringSubmatch(HEtime, 1)
	// Not a good match, nothing found
	if len(matches) < 1 {
		return nil
	}
	// By the regex, we are guaranteed to have good numbers
	day, _ := strconv.Atoi(matches[0][1])
	year, _ := strconv.Atoi(matches[0][2])
	// Subtract the 10k holoscene years
	year -= 10000

	tt := time.Date(year, time.January, 0, 0, 0, 0, 0, time.Local)
	tt = tt.Add(time.Duration(day) * 24 * time.Hour)
	return &tt
}

func AddHolosceneTitles(data string) string {
	// Match all paragraphs with holoscene time
	matches := HEParagraphRegex.FindAllStringSubmatch(data, -1)
	// No matches found, skip this file
	if len(matches) < 1 {
		return data
	}
	for _, match := range matches {
		HEtime := match[1]
		tt := ConvertHoloscene(HEtime)
		// Add the title to the paragraph
		data = strings.Replace(data,
			`>`+HEtime,
			` title="`+tt.Format(RFC_EMILY)+`">`+HEtime, 1,
		)
	}
	return data
}
