package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

const attributionFileName = "Attributions.md"

type config struct {
	scanFolder  string
	scanMatches []string
}

func parseConfig() config {
	i := flag.String("i", ".", "Folder to scan assets for.")
	flag.Parse()

	matches := flag.Args()

	return config{scanFolder: *i, scanMatches: matches}
}

func attribExists() bool {
	_, err := os.Stat(attributionFileName)

	return err == nil
}

func scanAssetsInFolder(folder string, match []string) ([]string, error) {
	var paths []string

	err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, v := range match {
			matched, err := filepath.Match("*."+v, filepath.Base(path))

			if err != nil {
				return err
			}

			if matched {
				paths = append(paths, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func createAttribFile(c *config) {
	fmt.Printf("⬇️  Creating \"%s\"...\n", attributionFileName)
	os.Create(attributionFileName)
}

func populateAttribFile(c *config) {
	var attribBuf bytes.Buffer

	fmt.Printf("🖊️  Populating...\n")

	assets, _ := scanAssetsInFolder(c.scanFolder, c.scanMatches)

	attribBuf.WriteString("# Asset attributions\n")
	for _, v := range assets {
		attribBuf.WriteString(fmt.Sprintf("## %s\nBy [author] from [source]. \n\n[LICENSE CLAUSE]\n\n", v))
	}

	f, _ := os.Create(attributionFileName)
	f.Write(attribBuf.Bytes())
}

func main() {
	c := parseConfig()
	now := time.Now()

	if attribExists() {
		populateAttribFile(&c)
	} else {
		createAttribFile(&c)
		populateAttribFile(&c)
	}

	fmt.Printf("✅ Done in %s!\n", time.Since(now))
}
