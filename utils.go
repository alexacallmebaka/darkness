package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/thecsw/darkness/emilia"
	"github.com/thecsw/darkness/html"
	"github.com/thecsw/darkness/internals"
	"github.com/thecsw/darkness/orgmode"
)

// bundle is a struct that hold filename and contents -- used for
// reading files and passing context or writing them too.
type bundle struct {
	File string
	Data string
}

// findFilesByExt finds all files with a given extension
func findFilesByExt(dir, ext string) ([]string, error) {
	files := make([]string, 0, 32)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ext {
			return nil
		}
		// Check if it is not excluded
		isExcluded := false
		for _, excludedPath := range emilia.Config.Project.Exclude {
			if strings.HasPrefix(path, excludedPath) {
				isExcluded = true
				break
			}
		}
		// Ignore hidden files
		if !isExcluded && !strings.HasPrefix(filepath.Base(path), ".") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// getTarget returns the target file name
func getTarget(file string) string {
	htmlFilename := strings.Replace(filepath.Base(file),
		emilia.Config.Project.Input, emilia.Config.Project.Output, 1)
	return filepath.Join(filepath.Dir(file), htmlFilename)
}

// orgToHTML converts an org file to html
func orgToHTML(file string) string {
	page := orgmode.ParseFile(workDir, file)
	htmlPage := exportAndEnrich(page)
	// Usually, each page only needs 1 holoscene replacement.
	// For the fortunes page, we need to replace all of them
	htmlPage = emilia.AddHolosceneTitles(htmlPage, func() int {
		if strings.HasSuffix(page.URL, "quotes") {
			return -1
		}
		return 1
	}())
	return htmlPage
}

// exportAndEnrich automatically applies all the emilia enhancements
// and converts Page into an html document.
func exportAndEnrich(page *internals.Page) string {
	emiliaStuff(page)
	result := html.ExportPage(page)
	result = emilia.AddHolosceneTitles(result, func() int {
		if strings.HasSuffix(page.URL, "quotes") {
			return -1
		}
		return 1
	}())
	return result
}

// emiliaStuff applies common emilia enhancements.
func emiliaStuff(page *internals.Page) {
	emilia.ResolveComments(page)
	emilia.EnrichHeadings(page)
	emilia.ResolveFootnotes(page)
	emilia.AddMathSupport(page)
	emilia.SourceCodeTrimLeftWhitespace(page)
	emilia.AddSyntaxHighlighting(page)
}