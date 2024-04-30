package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"
)

type testResult struct {
	Package string
	License string
	Score   int
	Extra   int
	Missing int
	Err     string
}

var mainModuleName string

func TestMain(m *testing.M) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		os.Exit(m.Run())
	}
	mainModuleName = bi.Main.Path
	m.Run()
}

func listTestLicenses(pkgs []string) ([]testResult, error) {
	gopath, err := filepath.Abs("testdata")
	if err != nil {
		return nil, err
	}
	licenses, err := listLicenses(gopath, pkgs)
	if err != nil {
		return nil, &MissingError{Err: err.Error()}
	}
	var res []testResult
	for _, l := range licenses {
		r := testResult{
			Package: l.Package,
		}
		if l.Template != nil {
			r.License = l.Template.Title
			r.Score = int(100 * l.Score)
		}
		if l.Err != "" {
			r.Err = "some error"
		}
		r.Extra = len(l.ExtraWords)
		r.Missing = len(l.MissingWords)
		res = append(res, r)
	}
	return res, nil
}

func compareTestLicenses(pkgs []string, wanted []testResult) error {
	stringify := func(res []testResult) string {
		var parts []string
		for _, r := range res {
			s := fmt.Sprintf("%s \"%s\" %d%%", r.Package, r.License, r.Score)
			if r.Err != "" {
				s += " " + r.Err
			}
			if r.Extra > 0 {
				s += fmt.Sprintf(" +%d", r.Extra)
			}
			if r.Missing > 0 {
				s += fmt.Sprintf(" -%d", r.Missing)
			}
			parts = append(parts, s)
		}
		return strings.Join(parts, "\n")
	}

	licenses, err := listTestLicenses(pkgs)
	if err != nil {
		return err
	}
	got := stringify(licenses)
	expected := stringify(wanted)
	if got != expected {
		return fmt.Errorf("licenses do not match:\n%s\n!=\n%s", got, expected)
	}
	return nil
}

func TestNoDependencies(t *testing.T) {
	err := compareTestLicenses([]string{"./testdata/src/colors/red"}, []testResult{
		{Package: mainModuleName + "/testdata/src/colors/red", License: "MIT License", Score: 98, Missing: 2},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipleLicenses(t *testing.T) {
	err := compareTestLicenses([]string{"./testdata/src/colors/blue"}, []testResult{
		{Package: mainModuleName + "/testdata/src/colors/blue", License: "Apache License 2.0", Score: 100},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoLicense(t *testing.T) {
	err := compareTestLicenses([]string{"./testdata/src/colors/green"}, []testResult{
		{Package: mainModuleName + "/testdata/src/colors/green", License: "", Score: 0},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMainWithDependencies(t *testing.T) {
	// It also tests license retrieval in parent directory.
	err := compareTestLicenses([]string{"./testdata/src/colors/cmd/paint"}, []testResult{
		{Package: mainModuleName + "/testdata/src/colors/cmd/paint", License: "Academic Free License v3.0", Score: 100},
		{Package: mainModuleName + "/testdata/src/colors/red", License: "MIT License", Score: 98, Missing: 2},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMainWithAliasedDependencies(t *testing.T) {
	err := compareTestLicenses([]string{"./testdata/src/colors/cmd/mix"}, []testResult{
		{Package: mainModuleName + "/testdata/src/colors/cmd/mix", License: "Academic Free License v3.0", Score: 100},
		{Package: mainModuleName + "/testdata/src/colors/red", License: "MIT License", Score: 98, Missing: 2},
		{Package: mainModuleName + "/testdata/src/couleurs/red", License: "GNU Lesser General Public License v2.1",
			Score: 100},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMissingPackage(t *testing.T) {
	_, err := listTestLicenses([]string{"./testdata/src/colors/missing"})
	if err == nil {
		t.Fatal("no error on missing package")
	}
	var missingError *MissingError
	if !errors.As(err, &missingError) {
		t.Fatalf("MissingError expected")
	}
}

func TestMismatch(t *testing.T) {
	err := compareTestLicenses([]string{"./testdata/src/colors/yellow"}, []testResult{
		{Package: mainModuleName + "/testdata/src/colors/yellow", License: "Microsoft Reciprocal License", Score: 25,
			Extra: 106, Missing: 131},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoBuildableGoSourceFiles(t *testing.T) {
	_, err := listTestLicenses([]string{"./testdata/src/colors/cmd"})
	if err == nil {
		t.Fatal("no error on missing package")
	}
}

func TestCleanLicenseData(t *testing.T) {
	data := `The MIT License (MIT)

	Copyright (c) 2013 Ben Johnson

	Some other lines.
	And more.
	`
	cleaned := string(cleanLicenseData([]byte(data)))
	wanted := "the mit license (mit)\n\n\tsome other lines.\n\tand more.\n\t"
	if wanted != cleaned {
		t.Fatalf("license data mismatch: %q\n!=\n%q", cleaned, wanted)
	}
}
