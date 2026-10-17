/*
Copyright © 2021-2024 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package portage

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	. "github.com/macaroni-os/anise/pkg/logger"
	gentoo "github.com/geaaru/pkgs-checker/pkg/gentoo"
)

const (
	Worldfile = "/var/lib/portage/world"
)

func CleanEdbCache(path string) error {
	if path == "" {
		path = "/var/cache/edb/"
	}

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf(
			"Error on read dir %s: %s", path, err.Error())
	}

	for _, file := range dirEntries {
		f := filepath.Join(path, file.Name())
		DebugC(fmt.Sprintf(":knife: Removing %s...", f))

		err := os.RemoveAll(f)
		if err != nil {
			return fmt.Errorf(
				"Error on removing %s: %s", f, err.Error())
		}
	}

	Info(":pizza: Portage edb cache purged!")

	return nil
}

func RebuildEdbCache() error {
	cmd := exec.Command("emerge", "--metadata")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("Error on start emerge command: %s", err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("Error on waiting emerge command: %s", err.Error())
	}

	if cmd.ProcessState.ExitCode() != 0 {
		return fmt.Errorf(
			"emerge command exiting with %d",
			cmd.ProcessState.ExitCode())
	}

	return nil
}

func GetWorldPackages(worldfile string) ([]string, error) {
	ans := []string{}

	if worldfile == "" {
		worldfile = Worldfile
	}

	// NOTE: I consider the world file not big and
	//       I read the content directly in memory.
	content, err := os.ReadFile(worldfile)
	if err != nil {
		return ans, fmt.Errorf(
			"Error on read file %s: %s", worldfile, err.Error())
	}

	packages := strings.Split(string(content), "\n")
	for _, p := range packages {
		// Exclude last split that returns an empty string
		if p != "" {
			ans = append(ans, p)
		}
	}

	return ans, nil
}

// Purge duplicates from Portage world list
func CleanupWorldPackages(packages []string) ([]string, error) {
	ans := []string{}

	// Convert the list in GentooPackage

	mPackages := make(map[string]*gentoo.GentooPackage, 0)

	for _, p := range packages {
		gp, err := gentoo.ParsePackageStr(p)
		if err != nil {
			return ans, fmt.Errorf("Error on parse package %s: %s",
				p, err.Error())
		}

		curpkg, ok := mPackages[gp.GetPackageNameWithSlot()]
		if ok {
			presentPkg := curpkg.GetPackageNameWithSlot()
			if curpkg.Repository != "" {
				presentPkg += "::" + curpkg.Repository
			}
			Warning(":haircut: Package %s ignored and already present as %s.",
				p, presentPkg,
			)
		} else {
			mPackages[gp.GetPackageNameWithSlot()] = gp
		}
	}

	for _, p := range mPackages {
		pkg := p.GetPackageNameWithSlot()
		if p.Repository != "" {
			pkg += "::" + p.Repository
		}

		ans = append(ans, pkg)
	}

	sort.Strings(ans)

	return ans, nil
}

func GetWorldPackagesCleaned(packages []string, pkgs2purge []string) ([]string, error) {
	ans := []string{}

	mPackages := make(map[string]*gentoo.GentooPackage, 0)

	// Create map for searching package.
	for _, p := range packages {
		gp, err := gentoo.ParsePackageStr(p)
		if err != nil {
			return ans, fmt.Errorf("Error on parse package %s: %s",
				p, err.Error())
		}

		curpkg, ok := mPackages[gp.GetPackageNameWithSlot()]
		if ok {
			presentPkg := curpkg.GetPackageNameWithSlot()
			if curpkg.Repository != "" {
				presentPkg += "::" + curpkg.Repository
			}
			Warning(":haircut: Package %s ignored and already present as %s.",
				p, presentPkg,
			)
		}
	}

	// Check if the packages to remove is present.
	for _, p := range pkgs2purge {

		gp, err := gentoo.ParsePackageStr(p)
		if err != nil {
			return ans, fmt.Errorf("Error on parse package %s: %s",
				p, err.Error())
		}

		if _, ok := mPackages[gp.GetPackageNameWithSlot()]; ok {
			Debug(":knife: Package %s removed", gp.GetPackageNameWithSlot())
			delete(mPackages, gp.GetPackageNameWithSlot())
		}
	}

	for _, p := range mPackages {
		pkg := p.GetPackageNameWithSlot()
		if p.Repository != "" {
			pkg += "::" + p.Repository
		}

		ans = append(ans, pkg)
	}

	sort.Strings(ans)

	return ans, nil
}

func WriteWorldPackages(worldFile string, packages []string) error {

	if worldFile == "" {
		worldFile = Worldfile
	}

	data := strings.Join(packages, "\n")
	data += "\n"

	err := os.WriteFile(worldFile, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("Error on write file %s: %s",
			worldFile, err.Error())
	}

	cmd := exec.Command(
		"chown", "root:portage", worldFile,
	)
	stdout, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("Error on chown directroy %s: %s",
			worldFile, stdout)
	}

	Info(":party_popper: Portage world file updated!")

	return nil
}
