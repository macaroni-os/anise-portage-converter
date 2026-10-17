/*
Copyright © 2021-2023 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/macaroni-os/anise-portage-converter/pkg/reposcan"
	"github.com/macaroni-os/anise-portage-converter/pkg/specs"

	"github.com/geaaru/pkgs-checker/pkg/gentoo"
	"github.com/macaroni-os/anise/pkg/config"
	fhelpers "github.com/macaroni-os/anise/pkg/helpers/file"
	. "github.com/macaroni-os/anise/pkg/logger"
	anise_pkg "github.com/macaroni-os/anise/pkg/package"
	pkg "github.com/macaroni-os/anise/pkg/package"
	artifact "github.com/macaroni-os/anise/pkg/v2/compiler/types/artifact"
	installer "github.com/macaroni-os/anise/pkg/v2/installer"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type PortageSyncOpts struct {
	OnlyNew bool
	DryRun  bool
	Force   bool
}

type PortageSyncTask struct {
	Opts      *PortageSyncOpts
	DbPkgsDir string
	Packages  map[string][]*gentoo.PortageMetaData
	Manager   *installer.ArtifactsManager
}

func parseRdepend(rdepend string, a *artifact.PackageArtifact,
	task *PortageSyncTask) error {

	if rdepend == "" {
		// Nothing to do
		return nil
	}

	deps, err := reposcan.ParseDependenciesMultiline(rdepend)
	if err != nil {
		return err
	}

	// For every dep I check if is present on local db.
	for _, d := range deps.Dependencies {
		// Check if the dependencies exists

		category := specs.SanitizeCategory(d.Dep.Category, d.Dep.Slot)
		version := ">=0"
		conflict := false
		switch d.Dep.Condition {
		case gentoo.PkgCondNot:
			conflict = true
		case gentoo.PkgCondNotLess:
			version = "<" + d.Dep.Version
			conflict = true
		case gentoo.PkgCondNotGreater:
			version = ">" + d.Dep.Version
			conflict = true
		}

		dp := pkg.NewPackageWithCat(category, d.Dep.Name, version,
			[]*pkg.DefaultPackage{}, []*pkg.DefaultPackage{})

		if conflict {
			a.Runtime.PackageConflicts = append(a.Runtime.PackageConflicts, dp)
			Debug(fmt.Sprintf("[%s] Add conflict with %s.",
				a.Runtime.PackageName(), dp.HumanReadableString()))
		} else {
			// Add the dependencies only if it's available on filesystem
			key := fmt.Sprintf("%s/%s", d.Dep.Category, d.Dep.Name)
			if _, ok := task.Packages[key]; ok {
				a.Runtime.PackageRequires = append(a.Runtime.PackageRequires, dp)
				Debug(fmt.Sprintf("[%s] Depends on %s.",
					a.Runtime.PackageName(), dp.HumanReadableString()))
			}
		}

	}

	return nil
}

func populateArtifact(p *gentoo.PortageMetaData, a *artifact.PackageArtifact,
	task *PortageSyncTask) error {
	files := []string{}

	// TODO: Check if this is correct ignoring directories
	for _, f := range p.CONTENTS {
		if f.Type == "dir" {
			// Ignoring directory
			continue
		}
		// Luet skip the first /
		files = append(files, f.File[1:])
	}

	// Add portage files not inserted in the CONTENTS
	files = append(files, []string{
		fmt.Sprintf("var/db/pkg/%s/%s/BUILD_TIME", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/CATEGORY", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/CONTENTS", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/COUNTER", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/DEFINED_PHASES", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/DESCRIPTION", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/EAPI", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/FEATURES", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/HOMEPAGE", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/INHERITED", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/IUSE", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/IUSE_EFFECTIVE", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/KEYWORDS", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/LICENSE", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/PF", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/SIZE", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/SLOT", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/USE", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/environment.bz2", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/repository", p.Category, p.GetPF()),
		fmt.Sprintf("var/db/pkg/%s/%s/%s.ebuild", p.Category, p.GetPF(), p.GetPF()),
	}...)

	// NOTE: the parsing return empty string also when the files exists but is empty.
	//       We need to check if the file exists.

	f := filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/CBUILD", p.Category, p.GetPF()))
	if p.CBUILD != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/CBUILD", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/CC", p.Category, p.GetPF()))
	if p.CC != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/CC", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/CFLAGS", p.Category, p.GetPF()))
	if p.CFlags != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/CFLAGS", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/CHOST", p.Category, p.GetPF()))
	if p.CHost != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/CHOST", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/CXX", p.Category, p.GetPF()))
	if p.Cxx != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/CXX", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/CXXFLAGS", p.Category, p.GetPF()))
	if p.CxxFlags != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/CXXFLAGS", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/CTARGET", p.Category, p.GetPF()))
	if p.CTARGET != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/CTARGET", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/DEPEND", p.Category, p.GetPF()))
	if p.DEPEND != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/DEPEND", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/DEBUGBUILD", p.Category, p.GetPF()))
	if p.DEBUGBUILD != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/DEBUGBUILD", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/LDFLAGS", p.Category, p.GetPF()))
	if p.LdFlags != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/LDFLAGS", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/NEEDED", p.Category, p.GetPF()))
	if p.NEEDED != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/NEEDED", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/NEEDED.ELF.2", p.Category, p.GetPF()))
	if p.NEEDED_ELF2 != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/NEEDED.ELF.2", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/RDEPEND", p.Category, p.GetPF()))
	if p.RDEPEND != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/RDEPEND", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/REQUIRES", p.Category, p.GetPF()))
	if p.REQUIRES != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/REQUIRES", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/PKGUSE", p.Category, p.GetPF()))
	if p.PKGUSE != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/PKGUSE", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/RESTRICT", p.Category, p.GetPF()))
	if p.RESTRICT != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/RESTRICT", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/PROVIDES", p.Category, p.GetPF()))
	if p.PROVIDES != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/PROVIDES", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/PROPERTIES", p.Category, p.GetPF()))
	if p.PROPERTIES != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/PROPERTIES", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/BDEPEND", p.Category, p.GetPF()))
	if p.BDEPEND != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/BDEPEND", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/PDEPEND", p.Category, p.GetPF()))
	if p.PDEPEND != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/PDEPEND", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/BINPKGMD5", p.Category, p.GetPF()))
	if p.BINPKGMD5 != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/BINPKGMD5", p.Category, p.GetPF()))
	}

	f = filepath.Join(task.DbPkgsDir, fmt.Sprintf("/%s/%s/QA_PREBUILT", p.Category, p.GetPF()))
	if p.QA_PREBUILT != "" || fhelpers.Exists(f) {
		files = append(files, fmt.Sprintf("var/db/pkg/%s/%s/QA_PREBUILT", p.Category, p.GetPF()))
	}

	a.Files = files

	a.Runtime = &anise_pkg.DefaultPackage{
		Name:           p.GetPN(),
		Category:       specs.SanitizeCategory(p.Category, p.Slot),
		Version:        p.Version,
		UseFlags:       p.UseFlags,
		Labels:         make(map[string]string, 0),
		Annotations:    make(map[string]interface{}, 0),
		Description:    p.DESCRIPTION,
		Uri:            []string{p.HOMEPAGE},
		License:        p.License,
		BuildTimestamp: p.BUILD_TIME,
		Hidden:         false,
		Repository:     "scm",
	}

	rules := make(map[string][]string, 0)
	rules["devel"] = []string{"^/usr/include/"}
	rules["portage"] = []string{"^/var/db/pkg/"}

	a.Runtime.Annotations["subsets"] = map[string]interface{}{
		"rules": rules,
	}

	a.Runtime.Labels["original.package.name"] = p.GetPackageName()
	a.Runtime.Labels["original.package.version"] = p.GetPVR()
	a.Runtime.Labels["original.package.slot"] = p.Slot
	a.Runtime.Labels["emerge.packages"] = p.GetPackageNameWithSlot()
	a.Runtime.Labels["kit"] = p.Repository

	return parseRdepend(p.RDEPEND, a, task)
}

func syncPackage(p *gentoo.PortageMetaData, task *PortageSyncTask,
	idx, n int) error {

	// TOSEE: We ignore subslot atm.
	if strings.Contains(p.Slot, "/") {
		Debug(fmt.Sprintf("[%s] Ignoring subslot", p.GetPackageNameWithSlot()))
		p.Slot = p.Slot[0:strings.Index(p.Slot, "/")]
	}

	Debug(fmt.Sprintf("[%4d/%4d] [%s] Analyzing...",
		idx+1, n, p.GetPackageNameWithSlot()))

	aniseP := &anise_pkg.DefaultPackage{
		Name:     p.Name,
		Category: specs.SanitizeCategory(p.Category, p.Slot),
		Version:  ">=0",
	}

	pkgs, err := task.Manager.Database.FindPackages(aniseP)
	if err != nil {
		return err
	}

	notFound := true
	var installedPkg *pkg.DefaultPackage = nil
	for _, ipkg := range pkgs {

		if ipkg.HasLabel("original.package.version") {
			originalPackageVersion := ipkg.GetLabels()["original.package.version"]

			if originalPackageVersion != p.GetPVR() {
				if task.Opts.OnlyNew {
					Debug(fmt.Sprintf(
						"[%4d/%4d] [%s] Found version %s on anise and %s on portage. Ignoring.",
						idx+1, n, p.GetPackageNameWithSlot(),
						originalPackageVersion,
						p.GetPVR()))
					notFound = false
				} else {
					Info(fmt.Sprintf(
						"[%4d/%4d] [%s] Found version %s on anise and %s on portage",
						idx+1, n, p.GetPackageNameWithSlot(),
						originalPackageVersion,
						p.GetPVR()))
					installedPkg = ipkg.(*pkg.DefaultPackage)
				}
			} else if originalPackageVersion == p.GetPVR() {
				notFound = false
			}
		}
	}

	if notFound {

		// Create the Package Artifact of the Portage package.
		art := artifact.NewPackageArtifact(
			fmt.Sprintf("/var/db/pkg/%s", p.GetPackageName()),
		)
		err := populateArtifact(p, art, task)
		if err != nil {
			Warning(fmt.Sprintf(
				"[%4d/%4d] [%s] Error on process data: %s",
				idx+1, n, p.GetPackageNameWithSlot(), err.Error()))
			return err
		}

		if len(pkgs) == 0 {
			// POST: the package is not available in the anise database.

			InfoC(GetAurora().Bold(fmt.Sprintf(
				"[%4d/%4d] [%s] Package with version %s not found on anise database.",
				idx+1, n, p.GetPackageNameWithSlot(),
				p.GetPVR())))

			if task.Opts.DryRun {
				InfoC(GetAurora().Bold(fmt.Sprintf(
					"[%4d/%4d] [%s] %s candidated for sync :heavy_check_mark:",
					idx+1, n, p.GetPackageNameWithSlot(),
					p.GetPVR())))
			} else {
				err := task.Manager.RegisterPackage(art, nil, true)
				if err != nil {
					return err
				}

				InfoC(GetAurora().Bold(fmt.Sprintf(
					"[%4d/%4d] [%s] %s added :heavy_check_mark:",
					idx+1, n, p.GetPackageNameWithSlot(),
					p.GetPVR())))
			}
		} else {
			// POST: Package to sync
			if task.Opts.DryRun {
				InfoC(GetAurora().Bold(fmt.Sprintf(
					"[%4d/%4d] [%s] %s candidated for align version to anise version :heavy_check_mark:",
					idx+1, n, p.GetPackageNameWithSlot(),
					p.GetPVR())))
			} else {

				// Delete existing package without remove filesystem files.
				err := task.Manager.Database.RemovePackageFiles(installedPkg)
				if err != nil {
					return err
				}

				err = task.Manager.Database.RemovePackageFinalizer(installedPkg)
				if err != nil && !task.Opts.Force {
					return errors.New("Failed removing package finalizer from database")
				}
				err = task.Manager.Database.RemovePackage(installedPkg)
				if err != nil && !task.Opts.Force {
					return errors.New("Failed removing package from database")
				}

				err = task.Manager.RegisterPackage(art, nil, true)
				if err != nil {
					return err
				}

				InfoC(GetAurora().Bold(fmt.Sprintf(
					"[%4d/%4d] [%s] %s added :heavy_check_mark:",
					idx+1, n, p.GetPackageNameWithSlot(),
					p.GetPVR())))
			}
		}
	}

	return nil
}

func newSyncCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sync",
		Aliases: []string{"rr"},
		Short:   "Sync portage tree to anise database.",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {

			debug, _ := cmd.Flags().GetBool("debug")
			onlyNew, _ := cmd.Flags().GetBool("only-new")
			force, _ := cmd.Flags().GetBool("force")
			dryRun, _ := cmd.Flags().GetBool("dry-run")
			dbPkgsDir, _ := cmd.Flags().GetString("db-pkgs-dir-path")
			verbose, _ := cmd.Flags().GetBool("verbose")
			pkgs, _ := cmd.Flags().GetStringArray("pkg")

			if debug {
				config.AniseCfg.GetGeneral().Debug = debug
			}

			var opts *gentoo.PortageUseParseOpts

			// TODO: permit to pass options from file.
			opts = &gentoo.PortageUseParseOpts{
				UseFilters: []string{
					"^userland_",
					"^kernel_",
					"^x86",
					"^x64",
					"^ppc",
					"^arm",
					"^amd64",
					"^prefix",
					"^m68k",
					"^ia64",
					"^riscv",
					"^s390",
					"^hppa",
					"^mips",
					"^alpha",
					"^sparc",
					"^elibc_",
				},
			}
			opts.Verbose = verbose
			opts.Packages = pkgs

			syncOpts := &PortageSyncOpts{
				OnlyNew: onlyNew,
				DryRun:  dryRun,
				Force:   force,
			}

			// Create sync task
			task := &PortageSyncTask{
				Opts:      syncOpts,
				DbPkgsDir: dbPkgsDir,
				Packages:  make(map[string][]*gentoo.PortageMetaData, 0),
			}

			// Retrieve the list of packages from portage db dir
			portagePkgs, err := gentoo.ParseMetadataDir(dbPkgsDir, opts)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
				os.Exit(1)
			}

			// Popolate portage map
			for _, p := range portagePkgs {
				key := fmt.Sprintf("%s/%s", p.Category, p.Name)
				if val, present := task.Packages[key]; present {
					task.Packages[key] = append(val, p)
				} else {
					task.Packages[key] = []*gentoo.PortageMetaData{p}
				}
			}

			// Initialize anise artifact manager
			task.Manager = installer.NewArtifactsManager(config.AniseCfg)
			defer task.Manager.Close()

			task.Manager.Setup()

			n := len(portagePkgs)

			for idx, p := range portagePkgs {
				err := syncPackage(p, task, idx, n)
				if err != nil {
					fmt.Println("ERROR ", err.Error())
				}

			} // end for idx, p

		},
	}

	var flags = cmd.Flags()
	flags.StringP("db-pkgs-dir-path", "p", "/var/db/pkg",
		"Path of the portage metadata.")
	flags.StringArray("pkg", []string{},
		"Specify one or more packages to sync.",
	)
	flags.Bool("dry-run", false, "Dry run sync operations.")
	flags.Bool("force", false, "Skip errors.")
	flags.Bool("only-new", false, "Sync only new packages not available on anise.")

	return cmd
}
