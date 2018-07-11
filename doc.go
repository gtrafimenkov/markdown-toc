package tocenize

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// Style describes the style of the TOC we want to produce
type Style int

const (
	// GitHub style of the TOC
	GitHub Style = iota

	// GitLab style of the TOC
	GitLab
)

type Job struct {
	MinDepth     int
	MaxDepth     int
	ExistingOnly bool
	Style        Style
}

type Document struct {
	Path  string
	Lines []string
	eol   string
}

func NewDocFromPath(path string) (Document, error) {
	b, err := ioutil.ReadFile(path)
	return NewDocFromString(string(b), path), err
}

func NewDocFromString(s string, path string) Document {
	doc := Document{Path: path}
	doc.Lines, doc.eol = lines(s)
	return doc
}

func lines(s string) (lines []string, eol string) {
	lines = strings.Split(s, "\n")
	eol = "\n"
	for i, l := range lines {
		if len(l) > 0 && l[len(l)-1] == '\r' {
			eol = "\r\n"
			lines[i] = strings.TrimRight(l, "\r")
		}
	}
	return
}

func (d Document) String() string {
	return strings.Join(d.Lines, d.eol)
}

func (d Document) Update(toc TOC, existingOnly bool, style Style) (Document, error) {
	s, e := d.FindTOC()
	if s == -1 {
		if existingOnly {
			return d, nil
		}
		s, e = d.SuggestTOC(toc)
	}
	nd := Document{Path: d.Path, eol: d.eol}
	if s > -1 && e > -1 {
		nd.Lines = make([]string, s)
		copy(nd.Lines, d.Lines[:s])
		for _, tocLine := range toc.Headings {
			nd.Lines = append(nd.Lines, tocLine.GetMarkdown(toc.MinDepth(), style))
		}
		// Add a new line after inserting a new TOC
		if s == e && len(toc.Headings) > 0 {
			nd.Lines = append(nd.Lines, "")
		}
		nd.Lines = append(nd.Lines, d.Lines[e:]...)
	}
	return nd, nil
}

// SuggestTOC looks for the first heading below a root heading.
// A root heading has minimum depth and the depth must only be used once.
// e.g.
//
// 	# Name
// 	--here--
// 	## A
// 	## B
func (d Document) SuggestTOC(toc TOC) (start, end int) {
	minCount := 0
	secCount := 0
	minIndex := 0
	for _, tocLine := range toc.Headings {
		if tocLine.Depth == toc.MinDepth() {
			minCount++
			if minCount > 1 {
				// too many root headings
				break
			}
			minIndex = tocLine.Index
		}
		if tocLine.Depth > toc.MinDepth() && minCount == 1 && secCount == 0 {
			secCount++
			minIndex = tocLine.Index
		}
	}
	if minCount == 1 {
		return minIndex, minIndex
	}
	if len(toc.Headings) > 0 {
		start = toc.Headings[0].Index
		return start, start
	}
	// in doubt, insert at top
	return 0, 0
}

func (d Document) FindTOC() (start, end int) {
	start = -1
	end = -1
	curStart := -1
	curEnd := -1
	isToc := false
	for i, l := range d.Lines {
		if reIsTOCLine.MatchString(l) {
			if !isToc {
				curStart = i
				isToc = true
			}
			curEnd = i + 1
			continue
		}
		if isToc {
			isToc = false
			if curEnd-curStart > end-start {
				start = curStart
				end = curEnd
			}
		}
	}
	return start, end
}

var reIsTOCLine = regexp.MustCompile(`^\s*- \[.*\]\(#.*\)\s*$`)
