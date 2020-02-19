// Quick and dirty script for generating JSON files from EAD files using zek
// https://jira.nyu.edu/jira/browse/FADESIGN-27
// For the JSON file output: https://jira.nyu.edu/jira/browse/FADESIGN-31

package main

import (
	zek "./zek"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type RepositoryEADXMLFile struct {
	repository string
	path       string
}

func usage() {
	fmt.Println("Usage: main EADXMLDirectory JSONDirectory")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	EADXMLDirectoryPath := os.Args[1]
	JSONDirectoryPath := os.Args[2]

	EADXMLDirectory, err := os.Stat(EADXMLDirectoryPath)
	if err != nil {
		panic(err)
	}

	JSONDirectory, err := os.Stat(JSONDirectoryPath)
	if err != nil {
		panic(err)
	}

	if ! os.FileInfo.IsDir(EADXMLDirectory) || ! os.FileInfo.IsDir(JSONDirectory) {
		usage()
		os.Exit(1)
	}

	EADXMLDirectoryPath, _ = filepath.Abs(EADXMLDirectoryPath)
	JSONDirectoryPath, _ = filepath.Abs(JSONDirectoryPath)

	os.RemoveAll(JSONDirectoryPath)
	os.MkdirAll(JSONDirectoryPath, 0755)

	var repositoryEADXMLFiles []RepositoryEADXMLFile
	repositories := make(map[string]bool)
	err = filepath.Walk(EADXMLDirectoryPath, func(path string, info os.FileInfo, err error) error {
		// Exclude hidden files and directories
		if strings.Contains(path, "/.") {
			return nil
		}

		if filepath.Ext(path) == ".xml" {
			repository := filepath.Base(filepath.Dir(path))
			repositories[repository] = true
			repositoryEADXMLFile := RepositoryEADXMLFile{
				repository: repository,
				path:       path,
			}

			repositoryEADXMLFiles = append(repositoryEADXMLFiles, repositoryEADXMLFile)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	for repository := range repositories {
		err = os.Mkdir(filepath.Join(JSONDirectoryPath, repository), 0755)
		if err != nil {
			panic(err)
		}
	}

	for _, EADXMLFile := range repositoryEADXMLFiles {
		eadid, jsonData := parseEADIDAndJSONFromEADXMLFile(EADXMLFile)

		// eadid must consist of digits, uppercase and lowercase letters, and
		// leading and/or trailing spaces
		re := regexp.MustCompile("^\\s*[0-9A-Za-z_\\-]+\\s*$")
		if re.MatchString(eadid) != true {
			fmt.Fprintf(os.Stderr, "ERROR: skipping %s: invalid <eadid> \"%s\"\n", EADXMLFile, eadid)
			continue
		}

		jsonFile := filepath.Join(JSONDirectoryPath,
			EADXMLFile.repository,
			strings.ToLower(strings.TrimSpace(eadid)) + ".json")

		fmt.Println(EADXMLFile.path + " -> " + jsonFile)

		err := ioutil.WriteFile(jsonFile, []byte(jsonData), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func parseEADIDAndJSONFromEADXMLFile(EADXMLFile RepositoryEADXMLFile) (string, string) {
	EADXML, err := ioutil.ReadFile(EADXMLFile.path)
	if err != nil {
		panic(err)
	}

	var ead zek.EAD
	err = xml.Unmarshal([]byte(EADXML), &ead)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse %s, %s", EADXMLFile, err))
	}

	jsonData, err := json.MarshalIndent(ead, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return ead.Eadheader.Eadid.Text, string(jsonData)
}
