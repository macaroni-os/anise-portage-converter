/*
Copyright © 2021-2024 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package reposcan

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	helpers "github.com/MottainaiCI/lxd-compose/pkg/helpers"
)

type ManifestFile struct {
	Md5   string         `json:"manifest_md5,omitempty" yaml:"manifest_md5,omitempty"`
	Files []RepoScanFile `json:"files,omitempty" yaml:"files,omitempty"`
}

func NewManifestFile(files []RepoScanFile) *ManifestFile {
	return &ManifestFile{
		Md5:   "",
		Files: files,
	}
}

func (m *ManifestFile) AddFiles(files []RepoScanFile) {
	m.Files = append(m.Files, files...)
}

func (m *ManifestFile) Write(f string) error {
	// Create a map of the files for sort by name
	mFiles := make(map[string]*RepoScanFile, 0)
	filesName := []string{}

	// On SRC_URI could be set the same file name multiple time with
	// failover download URL.
	// The Manifest file in this case must be generate with only one
	// entry for file.Name. I consider the same hash in all download urls.

	for idx := range m.Files {
		if _, present := mFiles[m.Files[idx].Name]; !present {
			mFiles[m.Files[idx].Name] = &m.Files[idx]
			filesName = append(filesName, m.Files[idx].Name)
		}
	}

	// TODO: At the moment we don't support Manifest with EBUILD rows
	sort.Strings(filesName)

	content := ""
	for _, name := range filesName {
		repoFile, _ := mFiles[name]

		blake2Bhash, withBlake2b := repoFile.Hashes["blake2b"]
		sha512hash, withSha512 := repoFile.Hashes["sha512"]
		md5hash, withMd5 := repoFile.Hashes["md5"]

		fields := []string{
			"DIST",
			name, repoFile.Size,
		}

		if withBlake2b {
			fields = append(fields, []string{"BLAKE2B", blake2Bhash}...)
		}
		if withSha512 {
			fields = append(fields, []string{"SHA512", sha512hash}...)
		}
		if withMd5 {
			fields = append(fields, []string{"MD5", md5hash}...)
		}

		content += strings.Join(fields, " ") + "\n"
	}

	return os.WriteFile(f, []byte(content), 0644)
}

func (m *ManifestFile) GetFiles(srcUri string) ([]RepoScanFile, error) {
	ans := []RepoScanFile{}

	srcUri = strings.TrimSpace(srcUri)

	if srcUri == "" {
		// POST: no tarballs and/or files defined.
		return ans, nil
	}

	words := strings.Split(srcUri, " ")

	toParse := len(words)
	idx := 0
	originUri := words[idx]
	for toParse > 0 {

		if words[idx] == "->" {
			idx++
			toParse--

			// Avoid to add two time the same file when the origin is equal
			// to alias
			if words[idx] == filepath.Base(originUri) {
				idx++
				toParse--
				continue
			}
		} else {
			originUri = words[idx]
		}

		baseName := filepath.Base(words[idx])
		// Check if the file is defined in the manifest
		for _, f := range m.Files {
			if f.Name == baseName {
				f.SrcUri = []string{originUri}
				ans = append(ans, f)
				break
			}
		}

		idx++
		toParse--
	}

	return ans, nil
}

func ParseManifestContent(content []byte) *ManifestFile {
	ans := &ManifestFile{
		Files: []RepoScanFile{},
	}

	ans.Md5 = fmt.Sprintf("%x", md5.Sum(content))

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		words := strings.Split(line, " ")
		if len(words) <= 3 || words[0] != "DIST" {
			continue
		}

		// The src_uri is populate later on processing metadata.
		file := &RepoScanFile{
			Size:   words[2],
			Name:   words[1],
			Hashes: make(map[string]string, 0),
		}
		pos := 3
		for pos < len(words) {
			file.Hashes[strings.ToLower(words[pos])] = words[pos+1]
			pos += 2
		}

		ans.Files = append(ans.Files, *file)
	}

	return ans
}

func ParseManifest(f string) (*ManifestFile, error) {
	ans := &ManifestFile{
		Files: []RepoScanFile{},
	}

	if helpers.Exists(f) {
		content, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}

		ans = ParseManifestContent(content)
	}

	return ans, nil
}
