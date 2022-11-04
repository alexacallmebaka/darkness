package orgmode

import (
	"regexp"

	"github.com/thecsw/darkness/internals"
)

const (
	commentPrefix      = "# "
	optionPrefix       = "#+"
	optionDropCap      = "drop_cap"
	optionBeginSource  = "begin_src"
	optionEndSource    = "end_src"
	optionBeginExport  = "begin_export"
	optionEndExport    = "end_export"
	optionBeginQuote   = "begin_quote"
	optionEndQuote     = "end_quote"
	optionBeginCenter  = "begin_center"
	optionEndCenter    = "end_center"
	optionBeginDetails = "begin_details"
	optionEndDetails   = "end_details"
	optionBeginGallery = "begin_gallery"
	optionEndGallery   = "end_gallery"
	optionCaption      = "caption"
	optionDate         = "date"
	horizontalLine     = "---"

	sectionLevelOne   = "* "
	sectionLevelTwo   = "** "
	sectionLevelThree = "*** "
	sectionLevelFour  = "**** "
	sectionLevelFive  = "***** "

	listSeparator    = string(rune(30))
	listSeparatorWS  = " " + listSeparator
	tableSeparator   = string(rune(29))
	tableSeparatorWS = " " + tableSeparator
)

var (
	surroundWithNewlines = []string{
		optionBeginQuote, optionEndQuote,
		optionBeginCenter, optionEndCenter,
		optionBeginDetails, optionEndDetails,
		optionBeginGallery, optionEndGallery,
	}
	// sourceCodeRegexp is the regexp for matching source blocks
	sourceCodeRegexp = regexp.MustCompile(`(?s)#\+begin_src ?([[:print:]]+)?`)
	// detailsRegexp is the regexp for matching details
	detailsRegexp = regexp.MustCompile(`(?s)#\+begin_details ?([[:print:]]+)?`)
	// linkRegexp is the regexp for matching links
	linkRegexp = internals.LinkRegexp
	// attentionBlockRegexp is the regexp for matching attention blocks
	attentionBlockRegexp = regexp.MustCompile(`^(WARNING|NOTE|TIP|IMPORTANT|CAUTION):\s*(.+)`)
	// unorderedListRegexp is the regexp for matching unordered lists
	unorderedListRegexp = regexp.MustCompile(`(?mU)- (.+) ` + listSeparator)
	// headingRegexp is the regexp for matching headlines
	headingRegexp = regexp.MustCompile(`(?m)^(\*\*\*\*\*|\*\*\*\*|\*\*\*|\*\*|\*\s+)`)
)
