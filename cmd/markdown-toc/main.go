package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/gtrafimenkov/markdown-toc"
	"github.com/pmezard/go-difflib/difflib"
)

var (
	VERSION  = "0.0.1"
	exitCode int
)

// Possible exit status codes.
const (
	ExitOk         = iota // nothing has changed or needs to change
	ExitDiff              // something has changed or needs to change
	ExitWrongUsage        // error in usage
)

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Println("markdown-toc [options] FILE...")
		fmt.Println()
		flag.PrintDefaults()
	}

	var styleStr string

	job := tocenize.Job{}
	flag.IntVar(&job.MinDepth, "min", 1, "minimum depth")
	flag.IntVar(&job.MaxDepth, "max", 99, "maximum depth")
	flag.StringVar(&tocenize.Indent, "indent", "\t", "string used for nesting")
	flag.StringVar(&styleStr, "style", "github", "style of the TOC: github or gitlab")
	doDiff := flag.Bool("d", false, "print full diff to stdout")
	doPrint := flag.Bool("p", false, "print full result to stdout")
	flag.BoolVar(&job.ExistingOnly, "e", false, "update only existing TOC (no insert)")
	showVersion := flag.Bool("v", false, "print version")
	flag.Parse()

	if *showVersion {
		fmt.Println(VERSION)
		os.Exit(ExitOk)
	}
	if flag.NArg() == 0 {
		log.Println("too few arguments")
		flag.Usage()
		os.Exit(ExitWrongUsage)
	}

	switch styleStr {
	case "github":
		job.Style = tocenize.GitHub
	case "gitlab":
		job.Style = tocenize.GitLab
	default:
		fmt.Fprintf(os.Stderr, "Unknown style %q", styleStr)
		os.Exit(1)
	}

	action := update
	if *doDiff {
		action = diff
	}
	if *doPrint {
		action = print
	}

	for _, arg := range flag.Args() {
		paths, err := filepath.Glob(arg)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, path := range paths {
			log.SetPrefix(path + ": ")
			err = runAction(path, job, action)
			if err != nil {
				log.Println(err)
			}
		}
		log.SetPrefix("")
	}
	os.Exit(exitCode)
}

type actionFunc func(job tocenize.Job, a, b tocenize.Document) error

func runAction(path string, job tocenize.Job, action actionFunc) error {
	doc, err := tocenize.NewDocFromPath(path)
	if err != nil {
		return err
	}
	toc := tocenize.NewTOC(doc, job)
	newDoc, err := doc.Update(toc, job.ExistingOnly, job.Style)
	if err != nil {
		return err
	}
	return action(job, doc, newDoc)
}

func diff(job tocenize.Job, a, b tocenize.Document) error {
	as := a.String()
	bs := b.String()
	if as != bs {
		exitCode = ExitDiff
	}
	ud := difflib.UnifiedDiff{
		A:        strings.SplitAfter(as, "\n"),
		B:        strings.SplitAfter(bs, "\n"),
		Context:  3,
		FromFile: a.Path + " (old)",
		ToFile:   a.Path + " (new)",
	}
	diff, err := difflib.GetUnifiedDiffString(ud)
	lines := strings.Split(diff, "\n")
	for i, line := range lines {
		if i+1 == len(lines) {
			break
		}
		switch line[0] {
		case '+':
			line = color.GreenString("%s", line)
		case '-':
			line = color.RedString("%s", line)
		case '@':
			line = color.YellowString("%s", line)
		}
		lines[i] = line
	}
	fmt.Print(strings.Join(lines, "\n"))
	return err
}

func print(job tocenize.Job, a, b tocenize.Document) error {
	fmt.Println(b.String())
	return nil
}

func update(job tocenize.Job, a, b tocenize.Document) error {
	return ioutil.WriteFile(b.Path, []byte(b.String()), 0644)
}
