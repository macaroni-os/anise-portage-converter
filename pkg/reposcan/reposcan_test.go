/*
Copyright © 2021-2026 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package reposcan_test

import (
	"errors"
	"fmt"

	. "github.com/macaroni-os/anise-portage-converter/pkg/reposcan"
	specs "github.com/macaroni-os/anise-portage-converter/pkg/specs"

	gentoo "github.com/geaaru/pkgs-checker/pkg/gentoo"
	anise_pkg "github.com/macaroni-os/anise/pkg/package"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reposcan", func() {
	Context("Json parser", func() {
		It("Load raw JSON", func() {

			source := `
{
  "cache_data_version": "1.0.5",
  "atoms": {
    "x11-misc/services_notificator-svn-0.1": {
      "relations": [],
      "relations_by_kind": {},
      "category": "x11-misc",
      "revision": "0",
      "package": "services_notificator-svn",
      "catpkg": "x11-misc/services_notificator-svn",
      "atom": "x11-misc/services_notificator-svn-0.1",
      "eclasses": [],
      "kit": "geaaru_overlay",
      "branch": "master",
      "metadata": null,
      "md5": "1750e323fe2cf4a6171da167afde6157",
      "metadata_out": "",
      "manifest_md5": "0a8c57df4540f66abba8f10d048fae38"
    }
  }
}
`
			resolver := NewRepoScanResolver()
			err := resolver.LoadRawJson(source, "test")
			Expect(err).Should(BeNil())

			err = resolver.BuildMap()
			Expect(err).Should(BeNil())

			val, ok := resolver.GetMap()["x11-misc/services_notificator-svn"]
			atom := val[0]
			Expect(ok).Should(Equal(true))
			Expect(atom.Package).Should(Equal("services_notificator-svn"))
			Expect(atom.Category).Should(Equal("x11-misc"))
			Expect(atom.Kit).Should(Equal("geaaru_overlay"))
		})

	})

	Context("Json parser2", func() {

		It("Load Json file with deps1", func() {

			resolver := NewRepoScanResolver()
			err := resolver.LoadJson("../../tests/fixtures/reposcan-01.json")
			Expect(err).Should(BeNil())
			resolver.SetIgnoreMissingDeps(true)
			err = resolver.BuildMap()
			Expect(err).Should(BeNil())

			val, ok := resolver.GetMap()["net-dns/noip-updater"]
			atom := val[0]
			Expect(ok).Should(Equal(true))
			Expect(atom.Package).Should(Equal("noip-updater"))
			Expect(atom.Category).Should(Equal("net-dns"))
			Expect(atom.Kit).Should(Equal("geaaru_overlay"))

			solution, err := resolver.Resolve(
				"net-dns/noip-updater",
				specs.NewPortageResolverOpts())
			Expect(err).Should(BeNil())

			Expect(solution.Package.GetPackageName()).Should(Equal("net-dns/noip-updater"))
			Expect(solution.Package.Slot).Should(Equal("0"))
			Expect(solution.Package.License).Should(Equal("GPL-2"))
			Expect(len(solution.BuildDeps)).Should(Equal(1))
			Expect(solution.BuildDeps[0]).Should(Equal(gentoo.GentooPackage{
				Name:          "gcc",
				Category:      "sys-devel",
				Version:       "10.2.0",
				VersionSuffix: "-r3",
				Slot:          "10",
				License:       "GPL-3+ LGPL-3+ || ( GPL-3+ libgcc libstdc++ gcc-runtime-library-exception-3.1 ) FDL-1.3+",
				Repository:    "gentoo",
				Condition:     gentoo.PkgCondEqual,
			}))
			Expect(len(solution.RuntimeDeps)).Should(Equal(0))

			Expect(solution.ToPack(false)).Should(Equal(
				&anise_pkg.DefaultPackage{
					Name:        "noip-updater",
					Category:    "net-dns",
					Version:     "2.1.9",
					Description: "no-ip.com dynamic DNS updater",
					Uri: []string{
						"http://www.no-ip.com",
					},
					License: "GPL-2",
					PackageRequires: []*anise_pkg.DefaultPackage{
						&anise_pkg.DefaultPackage{
							Name:     "gcc",
							Category: "sys-devel-10",
							Version:  ">=0",
						},
					},
					PackageConflicts: nil,
					Labels: map[string]string{
						"emerge.packages":          "net-dns/noip-updater",
						"kit":                      "geaaru_overlay",
						"DEPEND":                   "sys-devel/gcc virtual/pkgconfig",
						"original.package.name":    "net-dns/noip-updater",
						"original.package.version": "2.1.9-r1",
						"original.package.slot":    "0",
					},
				},
			))
		})

		It("Load Json file with deps1 without slot on category", func() {

			resolver := NewRepoScanResolver()
			err := resolver.LoadJson("../../tests/fixtures/reposcan-01.json")
			Expect(err).Should(BeNil())
			resolver.SetIgnoreMissingDeps(true)
			resolver.SetDepsWithSlot(false)
			err = resolver.BuildMap()
			Expect(err).Should(BeNil())

			val, ok := resolver.GetMap()["net-dns/noip-updater"]
			atom := val[0]
			Expect(ok).Should(Equal(true))
			Expect(atom.Package).Should(Equal("noip-updater"))
			Expect(atom.Category).Should(Equal("net-dns"))
			Expect(atom.Kit).Should(Equal("geaaru_overlay"))

			solution, err := resolver.Resolve(
				"net-dns/noip-updater",
				specs.NewPortageResolverOpts())
			Expect(err).Should(BeNil())

			Expect(solution.Package.GetPackageName()).Should(Equal("net-dns/noip-updater"))
			Expect(solution.Package.Slot).Should(Equal("0"))
			Expect(solution.Package.License).Should(Equal("GPL-2"))
			Expect(len(solution.BuildDeps)).Should(Equal(1))
			Expect(solution.BuildDeps[0]).Should(Equal(gentoo.GentooPackage{
				Name:          "gcc",
				Category:      "sys-devel",
				Version:       "10.2.0",
				VersionSuffix: "-r3",
				Slot:          "",
				License:       "GPL-3+ LGPL-3+ || ( GPL-3+ libgcc libstdc++ gcc-runtime-library-exception-3.1 ) FDL-1.3+",
				Repository:    "gentoo",
				Condition:     gentoo.PkgCondEqual,
			}))
			Expect(len(solution.RuntimeDeps)).Should(Equal(0))

			Expect(solution.ToPack(false)).Should(Equal(
				&anise_pkg.DefaultPackage{
					Name:        "noip-updater",
					Category:    "net-dns",
					Version:     "2.1.9",
					Description: "no-ip.com dynamic DNS updater",
					Uri: []string{
						"http://www.no-ip.com",
					},
					License: "GPL-2",
					PackageRequires: []*anise_pkg.DefaultPackage{
						&anise_pkg.DefaultPackage{
							Name:     "gcc",
							Category: "sys-devel",
							Version:  ">=0",
						},
					},
					PackageConflicts: nil,
					Labels: map[string]string{
						"emerge.packages":          "net-dns/noip-updater",
						"kit":                      "geaaru_overlay",
						"DEPEND":                   "sys-devel/gcc virtual/pkgconfig",
						"original.package.name":    "net-dns/noip-updater",
						"original.package.version": "2.1.9-r1",
						"original.package.slot":    "0",
					},
				},
			))
		})

		It("Load Json file with deps2", func() {

			resolver := NewRepoScanResolver()
			err := resolver.LoadJson("../../tests/fixtures/reposcan-01.json")
			Expect(err).Should(BeNil())
			resolver.SetIgnoreMissingDeps(true)
			err = resolver.BuildMap()
			Expect(err).Should(BeNil())

			val, ok := resolver.GetMap()["net-dns/noip-updater"]
			atom := val[0]
			Expect(ok).Should(Equal(true))
			Expect(atom.Package).Should(Equal("noip-updater"))
			Expect(atom.Category).Should(Equal("net-dns"))
			Expect(atom.Kit).Should(Equal("geaaru_overlay"))

			_, err = resolver.Resolve(
				">=net-dns/noip-updater-2.2.0",
				specs.NewPortageResolverOpts(),
			)
			Expect(err).Should(Equal(errors.New("No package found matching >=net-dns/noip-updater-2.2.0")))

		})

		It("Load Json file with deps3", func() {

			resolver := NewRepoScanResolver()
			err := resolver.LoadJson("../../tests/fixtures/reposcan-01.json")
			Expect(err).Should(BeNil())
			resolver.SetIgnoreMissingDeps(true)
			err = resolver.BuildMap()
			Expect(err).Should(BeNil())

			val, ok := resolver.GetMap()["net-dns/noip-updater"]
			atom := val[0]
			Expect(ok).Should(Equal(true))
			Expect(atom.Package).Should(Equal("noip-updater"))
			Expect(atom.Category).Should(Equal("net-dns"))
			Expect(atom.Kit).Should(Equal("geaaru_overlay"))

			_, err = resolver.Resolve(
				"net-dns/noip-updater:2",
				specs.NewPortageResolverOpts(),
			)
			Expect(err).Should(Equal(errors.New("No package found matching net-dns/noip-updater:2")))

		})

		It("Load Json file with deps4 (with runtime deps)", func() {

			resolver := NewRepoScanResolver()
			err := resolver.LoadJson("../../tests/fixtures/reposcan-01.json")
			Expect(err).Should(BeNil())
			resolver.SetIgnoreMissingDeps(true)
			err = resolver.BuildMap()
			Expect(err).Should(BeNil())

			val, ok := resolver.GetMap()["sys-devel/gcc"]
			atom := val[0]
			Expect(ok).Should(Equal(true))
			Expect(atom.Package).Should(Equal("gcc"))
			Expect(atom.Category).Should(Equal("sys-devel"))
			Expect(atom.Kit).Should(Equal("gentoo"))
			rdeps, err := atom.GetRuntimeDeps()
			Expect(err).Should(BeNil())
			Expect(len(rdeps)).Should(Equal(8))

			solution, err := resolver.Resolve(
				"sys-devel/gcc:9.3.0",
				specs.NewPortageResolverOpts(),
			)
			Expect(err).Should(BeNil())
			// We have only one runtime deps because atom aren't
			// available on json file.
			Expect(len(solution.RuntimeDeps)).Should(Equal(1))
			fmt.Println(solution)
		})
	})

})
