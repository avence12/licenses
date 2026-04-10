package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/avence12/licenses"
)

func printLicenses() error {
	msg := `Usage: licenses IMPORTPATH...

	licenses lists all dependencies of specified packages or commands, excluding
	standard library packages, and prints their licenses. Licenses are detected by
	looking for files named like LICENSE, COPYING, COPYRIGHT and other variants in
	the package directory, and its parent directories until one is found. Files
	content is matched against a set of well-known licenses and the best match is
	displayed along with its score.

	With -a, all individual packages are displayed instead of grouping them by
	license files.
	With -w, words in package license file not found in the template license are
	displayed. It helps assessing the changes importance.
	`
	flag.Usage = func() {
		fmt.Println(msg)
		os.Exit(1)
	}
	all := flag.Bool("a", false, "display all individual packages")
	words := flag.Bool("w", false, "display words not matching license template")
	flag.Parse()
	if flag.NArg() < 1 {
		return fmt.Errorf("expect at least one package argument")
	}
	pkgs := flag.Args()

	confidence := 0.9
	results, err := licenses.ListLicenses("", pkgs)
	if err != nil {
		return err
	}
	if !*all {
		results, err = licenses.GroupLicenses(results)
		if err != nil {
			return err
		}
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 4, 2, ' ', 0)
	for _, l := range results {
		license := "?"
		if l.Template != nil {
			if l.Score > .99 {
				license = fmt.Sprintf("%s", l.Template.Title)
			} else if l.Score >= confidence {
				license = fmt.Sprintf("%s (%2d%%)", l.Template.Title, int(100*l.Score))
				if *words && len(l.ExtraWords) > 0 {
					license += "\n\t+words: " + strings.Join(l.ExtraWords, ", ")
				}
				if *words && len(l.MissingWords) > 0 {
					license += "\n\t-words: " + strings.Join(l.MissingWords, ", ")
				}
			} else {
				license = fmt.Sprintf("? (%s, %2d%%)", l.Template.Title, int(100*l.Score))
			}
		} else if l.Err != "" {
			license = strings.Replace(l.Err, "\n", " ", -1)
		}
		_, err = w.Write([]byte(l.Package + "\t" + license + "\n"))
		if err != nil {
			return err
		}
	}
	return w.Flush()
}

func main() {
	err := printLicenses()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
