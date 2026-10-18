/*
Copyright © 2021-2026 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package reposcan

import (
	"fmt"

	_gentoo "github.com/geaaru/pkgs-checker/pkg/gentoo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GentooBuilder", func() {
	Context("Parse Dependencies 1", func() {

		rdepend := `
							   	app-crypt/sbsigntools
							   	x11-themes/sabayon-artwork-grub
							   	sys-boot/os-prober
							   	app-arch/xz-utils
							   	>=sys-libs/ncurses-5.2-r5:0=
							   `
		gr, err := ParseDependenciesMultiline(rdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(5))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sbsigntools",
						Category: "app-crypt",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sabayon-artwork-grub",
						Category: "x11-themes",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep5", func() {
			Expect(*gr.Dependencies[4]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:          "ncurses",
						Category:      "sys-libs",
						Slot:          "0",
						Version:       "5.2",
						VersionSuffix: "-r5",
						Condition:     _gentoo.PkgCondGreaterEqual,
					},
				},
			))
		})

	})

	Context("Parse Dependencies 2", func() {

		rdepend := `
							   	app-crypt/sbsigntools
							   	x11-themes/sabayon-artwork-grub
							   	sys-boot/os-prober
							   	app-arch/xz-utils
							   	>=sys-libs/ncurses-5.2-r5:0=
							   	mount? ( sys-fs/fuse )
							   `
		gr, err := ParseDependenciesMultiline(rdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(6))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sbsigntools",
						Category: "app-crypt",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sabayon-artwork-grub",
						Category: "x11-themes",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep5", func() {
			Expect(*gr.Dependencies[4]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:          "ncurses",
						Category:      "sys-libs",
						Slot:          "0",
						Version:       "5.2",
						VersionSuffix: "-r5",
						Condition:     _gentoo.PkgCondGreaterEqual,
					},
				},
			))
		})

		It("Check dep6", func() {
			Expect(*gr.Dependencies[5]).Should(Equal(
				GentooDependency{
					Use:          "mount",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps: []*GentooDependency{
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:     "fuse",
								Category: "sys-fs",
								Slot:     "0",
							},
						},
					},
					Dep: nil,
				},
			))
		})

	})

	Context("Parse Dependencies 3", func() {

		rdepend := `
							   	app-crypt/sbsigntools
							   	x11-themes/sabayon-artwork-grub
							   	sys-boot/os-prober
							   	app-arch/xz-utils
							   	>=sys-libs/ncurses-5.2-r5:0=
							   	mount? ( sys-fs/fuse =sys-apps/pmount-0.9.99_alpha-r5:= )
							   `
		gr, err := ParseDependenciesMultiline(rdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(6))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sbsigntools",
						Category: "app-crypt",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sabayon-artwork-grub",
						Category: "x11-themes",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep5", func() {
			Expect(*gr.Dependencies[4]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:          "ncurses",
						Category:      "sys-libs",
						Slot:          "0",
						Version:       "5.2",
						VersionSuffix: "-r5",
						Condition:     _gentoo.PkgCondGreaterEqual,
					},
				},
			))
		})

		It("Check dep6", func() {
			Expect(*gr.Dependencies[5]).Should(Equal(
				GentooDependency{
					Use:          "mount",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps: []*GentooDependency{
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:     "fuse",
								Category: "sys-fs",
								Slot:     "0",
							},
						},
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:          "pmount",
								Category:      "sys-apps",
								Condition:     _gentoo.PkgCondEqual,
								Version:       "0.9.99",
								VersionSuffix: "_alpha-r5",
								Slot:          "",
							},
						},
					},
					Dep: nil,
				},
			))
		})

	})

	Context("Parse Dependencies", func() {

		rdepend := `
							   	app-crypt/sbsigntools
							   	x11-themes/sabayon-artwork-grub
							   	sys-boot/os-prober
							   	app-arch/xz-utils
							   	>=sys-libs/ncurses-5.2-r5:0=
							   	!mount? ( sys-fs/fuse =sys-apps/pmount-0.9.99_alpha-r5:= )
							   `
		gr, err := ParseDependenciesMultiline(rdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(6))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sbsigntools",
						Category: "app-crypt",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sabayon-artwork-grub",
						Category: "x11-themes",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep5", func() {
			Expect(*gr.Dependencies[4]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:          "ncurses",
						Category:      "sys-libs",
						Slot:          "0",
						Version:       "5.2",
						VersionSuffix: "-r5",
						Condition:     _gentoo.PkgCondGreaterEqual,
					},
				},
			))
		})

		It("Check dep6", func() {
			Expect(*gr.Dependencies[5]).Should(Equal(
				GentooDependency{
					Use:          "mount",
					UseCondition: _gentoo.PkgCondNot,
					SubDeps: []*GentooDependency{
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:     "fuse",
								Category: "sys-fs",
								Slot:     "0",
							},
						},
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:          "pmount",
								Category:      "sys-apps",
								Condition:     _gentoo.PkgCondEqual,
								Version:       "0.9.99",
								VersionSuffix: "_alpha-r5",
								Slot:          "",
							},
						},
					},
					Dep: nil,
				},
			))
		})

	})

	Context("Parse Dependencies 5", func() {

		rdepend := `
							   	app-crypt/sbsigntools
							   	>=sys-libs/ncurses-5.2-r5:0=
							   	mount? (
							   		sys-fs/fuse
							   		=sys-apps/pmount-0.9.99_alpha-r5:=
							   	)
							   `
		gr, err := ParseDependenciesMultiline(rdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(3))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sbsigntools",
						Category: "app-crypt",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:          "ncurses",
						Category:      "sys-libs",
						Slot:          "0",
						Version:       "5.2",
						VersionSuffix: "-r5",
						Condition:     _gentoo.PkgCondGreaterEqual,
					},
				},
			))
		})

		It("Check dep3", func() {
			Expect(*gr.Dependencies[2]).Should(Equal(
				GentooDependency{
					Use:          "mount",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps: []*GentooDependency{
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:     "fuse",
								Category: "sys-fs",
								Slot:     "0",
							},
						},
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:          "pmount",
								Category:      "sys-apps",
								Condition:     _gentoo.PkgCondEqual,
								Version:       "0.9.99",
								VersionSuffix: "_alpha-r5",
								Slot:          "",
							},
						},
					},
					Dep: nil,
				},
			))
		})

	})

	Context("Parse Dependencies 6", func() {

		rdepend := `
							   	app-crypt/sbsigntools
							   	>=sys-libs/ncurses-5.2-r5:0=
							   	mount? (
							   		sys-fs/fuse
							   		=sys-apps/pmount-0.9.99_alpha-r5:= )
							   `
		gr, err := ParseDependenciesMultiline(rdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(3))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sbsigntools",
						Category: "app-crypt",
						Slot:     "0",
					},
				},
			))
		})
	})
	Context("Parse Dependencies 7", func() {

		rdepend := `
							   	app-crypt/sbsigntools
							   	>=sys-libs/ncurses-5.2-r5:0=
							   	mount? (
							   		sys-fs/fuse
							   		=sys-apps/pmount-0.9.99_alpha-r5:=
							   		ext2? (
							   			sys-fs/genext2fs
							   		)
							   	)
							   `
		gr, err := ParseDependenciesMultiline(rdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(3))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "sbsigntools",
						Category: "app-crypt",
						Slot:     "0",
					},
				},
			))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:          "ncurses",
						Category:      "sys-libs",
						Slot:          "0",
						Version:       "5.2",
						VersionSuffix: "-r5",
						Condition:     _gentoo.PkgCondGreaterEqual,
					},
				},
			))
		})

		It("Check dep3", func() {
			Expect(*gr.Dependencies[2]).Should(Equal(
				GentooDependency{
					Use:          "mount",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps: []*GentooDependency{
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:     "fuse",
								Category: "sys-fs",
								Slot:     "0",
							},
						},
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:          "pmount",
								Category:      "sys-apps",
								Condition:     _gentoo.PkgCondEqual,
								Version:       "0.9.99",
								VersionSuffix: "_alpha-r5",
								Slot:          "",
							},
						},
						&GentooDependency{
							Use:          "ext2",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps: []*GentooDependency{
								&GentooDependency{
									Use:          "",
									UseCondition: _gentoo.PkgCondInvalid,
									SubDeps:      make([]*GentooDependency, 0),
									Dep: &_gentoo.GentooPackage{
										Name:     "genext2fs",
										Category: "sys-fs",
										Slot:     "0",
									},
								},
							},
							Dep: nil,
						},
					},
				},
			))
		})

	})

	Context("Parse Dependencies 8", func() {

		rdepend := `sys-libs/zlib  virtual/libintl`
		gr, err := ParseDependencies(rdepend)

		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(2))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "zlib",
						Category: "sys-libs",
						Slot:     "0",
					},
				},
			))
		})
		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "libintl",
						Category: "virtual",
						Slot:     "0",
					},
				},
			))
		})
	})

	Context("Parse Dependencies 9", func() {

		rdepend := `sys-libs/zlib nls? ( virtual/libintl )`
		gr, err := ParseDependencies(rdepend)

		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(2))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "zlib",
						Category: "sys-libs",
						Slot:     "0",
					},
				},
			))
		})
		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(
				GentooDependency{
					Use:          "nls",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps: []*GentooDependency{
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:     "libintl",
								Category: "virtual",
								Slot:     "0",
							},
						},
					},
				},
			))
		})
	})

	Context("Parse Dependencies 10", func() {

		rdepend := `sys-libs/zlib nls? ( virtual/libintl ) virtual/libiconv >=dev-libs/gmp-4.3.2:0= >=dev-libs/mpfr-2.4.2:0= >=dev-libs/mpc-0.8.1:0= objc-gc? ( >=dev-libs/boehm-gc-7.4.2 ) graphite? ( >=dev-libs/isl-0.14:0= )`
		gr, err := ParseDependencies(rdepend)

		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(8))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "zlib",
						Category: "sys-libs",
						Slot:     "0",
					},
				},
			))
		})
		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(
				GentooDependency{
					Use:          "nls",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps: []*GentooDependency{
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:     "libintl",
								Category: "virtual",
								Slot:     "0",
							},
						},
					},
				},
			))
		})
		It("Check dep3", func() {
			Expect(*gr.Dependencies[2]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:     "libiconv",
						Category: "virtual",
						Slot:     "0",
					},
				},
			))
		})
		It("Check dep4", func() {
			Expect(*gr.Dependencies[3]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:      "gmp",
						Category:  "dev-libs",
						Version:   "4.3.2",
						Condition: _gentoo.PkgCondGreaterEqual,
						Slot:      "0",
					},
				},
			))
		})

		It("Check dep5", func() {
			Expect(*gr.Dependencies[4]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:      "mpfr",
						Category:  "dev-libs",
						Version:   "2.4.2",
						Condition: _gentoo.PkgCondGreaterEqual,
						Slot:      "0",
					},
				},
			))
		})

		It("Check dep6", func() {
			Expect(*gr.Dependencies[5]).Should(Equal(
				GentooDependency{
					Use:          "",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps:      make([]*GentooDependency, 0),
					Dep: &_gentoo.GentooPackage{
						Name:      "mpc",
						Category:  "dev-libs",
						Version:   "0.8.1",
						Condition: _gentoo.PkgCondGreaterEqual,
						Slot:      "0",
					},
				},
			))
		})

		It("Check dep7", func() {
			Expect(*gr.Dependencies[6]).Should(Equal(
				GentooDependency{
					Use:          "objc-gc",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps: []*GentooDependency{
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:      "boehm-gc",
								Category:  "dev-libs",
								Version:   "7.4.2",
								Condition: _gentoo.PkgCondGreaterEqual,
								Slot:      "0",
							},
						},
					},
				},
			))
		})

		It("Check dep8", func() {
			Expect(*gr.Dependencies[7]).Should(Equal(
				GentooDependency{
					Use:          "graphite",
					UseCondition: _gentoo.PkgCondInvalid,
					SubDeps: []*GentooDependency{
						&GentooDependency{
							Use:          "",
							UseCondition: _gentoo.PkgCondInvalid,
							SubDeps:      make([]*GentooDependency, 0),
							Dep: &_gentoo.GentooPackage{
								Name:      "isl",
								Category:  "dev-libs",
								Version:   "0.14",
								Condition: _gentoo.PkgCondGreaterEqual,
								Slot:      "0",
							},
						},
					},
				},
			))
		})

	})

	Context("Parse Dependencies 11", func() {
		bdepend := `minizip? ( || ( >=sys-devel/automake-1.16.1:1.16 >=sys-devel/automake-1.15.1:1.15 ) >=sys-devel/autoconf-2.69 >=sys-devel/libtool-2.4 ) >=app-portage/elt-patches-20170815`

		gr, err := ParseDependencies(bdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(2))
		})

		d1, e1 := NewGentooDependency(">=sys-devel/autoconf-2.69", "")
		d2, e2 := NewGentooDependency(">=sys-devel/libtool-2.4", "")
		d3, e3 := NewGentooDependency(">=app-portage/elt-patches-20170815", "")
		d5, e4 := NewGentooDependency(">=sys-devel/automake-1.16.1:1.16", "")
		d6, e5 := NewGentooDependency(">=sys-devel/automake-1.15.1:1.15", "")
		It("Check error d1", func() {
			Expect(e1).Should(BeNil())
			Expect(e2).Should(BeNil())
			Expect(e3).Should(BeNil())
			Expect(e4).Should(BeNil())
			Expect(e5).Should(BeNil())
		})
		aggr1, err_aggr := NewGentooDependencyWithSubdeps("", "", []*GentooDependency{d5, d6})
		It("Check error dor", func() {
			Expect(err_aggr).Should(BeNil())
		})
		dor, err_dor := NewGentooDependencyWithSubdeps("", "", []*GentooDependency{aggr1})
		It("Check error dor", func() {
			Expect(err_dor).Should(BeNil())
		})
		dor.DepInOr = true
		minizip, err := NewGentooDependencyWithSubdeps("", "minizip",
			[]*GentooDependency{dor, d1, d2})
		It("Check error minizip", func() {
			Expect(err).Should(BeNil())
		})
		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(*minizip))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(*d3))
		})

	})
	Context("Parse Dependencies 12", func() {

		rdepend := `aqua? ( x11-libs/gtk+:2[aqua=,abi_x86_32(-)?,abi_x86_64(-)?,abi_x86_x32(-)?,abi_mips_n32(-)?,abi_mips_n64(-)?,abi_mips_o32(-)?,abi_s390_32(-)?,abi_s390_64(-)?] virtual/jpeg:0=[abi_x86_32(-)?,abi_x86_64(-)?,abi_x86_x32(-)?,abi_mips_n32(-)?,abi_mips_n64(-)?,abi_mips_o32(-)?,abi_s390_32(-)?,abi_s390_64(-)?] tiff? ( media-libs/tiff:0[abi_x86_32(-)?,abi_x86_64(-)?,abi_x86_x32(-)?,abi_mips_n32(-)?,abi_mips_n64(-)?,abi_mips_o32(-)?,abi_s390_32(-)?,abi_s390_64(-)?] ) )`
		gr, err := ParseDependencies(rdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(1))
		})

		gtk, e1 := NewGentooDependency("x11-libs/gtk+:2", "")
		vjpeg, e1 := NewGentooDependency("virtual/jpeg:0=", "")
		tiff, e2 := NewGentooDependency("media-libs/tiff:0", "")

		tiffUse, e3 := NewGentooDependencyWithSubdeps("", "tiff",
			[]*GentooDependency{tiff})

		aquaUse, e4 := NewGentooDependencyWithSubdeps("", "aqua",
			[]*GentooDependency{gtk, vjpeg, tiffUse})

		It("Check error", func() {
			Expect(e1).Should(BeNil())
			Expect(e2).Should(BeNil())
			Expect(e3).Should(BeNil())
			Expect(e4).Should(BeNil())
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(*aquaUse))
		})
	})

	Context("Parse Dependencies 13", func() {

		bdepend := `virtual/pkgconfig test? ( || ( dev-lang/python:3.8 dev-lang/python:3.7 dev-lang/python:3.6 ) || ( ( dev-lang/python:3.8 dev-python/pytest[python_targets_python3_8(-),python_single_target_python3_8(+)] ) ( dev-lang/python:3.7 dev-python/pytest[python_targets_python3_7(-),python_single_target_python3_7(+)] ) ( dev-lang/python:3.6 dev-python/pytest[python_targets_python3_6(-),python_single_target_python3_6(+)] ) ) ) >=dev-util/meson-0.54.0 >=dev-util/ninja-1.8.2 virtual/pkgconfig`

		gr, err := ParseDependencies(bdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(5))
		})

		pkgConfig, e1 := NewGentooDependency("virtual/pkgconfig", "")

		py38, e2 := NewGentooDependency("dev-lang/python:3.8", "")
		py37, e3 := NewGentooDependency("dev-lang/python:3.7", "")
		py36, e4 := NewGentooDependency("dev-lang/python:3.6", "")

		pyContainer, _ := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{py38, py37, py36})

		orPy, e5 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{pyContainer})
		orPy.DepInOr = true

		pytest, e6 := NewGentooDependency("dev-python/pytest", "")

		orPytestpy38, e7 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{py38, pytest})

		orPytestpy37, e8 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{py37, pytest})

		orPytestpy36, e9 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{py36, pytest})

		orPyTests, e10 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{orPytestpy38, orPytestpy37, orPytestpy36})
		orPyTests.DepInOr = true

		testUse, e10 := NewGentooDependencyWithSubdeps("", "test",
			[]*GentooDependency{orPy, orPyTests})

		meson, e11 := NewGentooDependency(">=dev-util/meson-0.54.0", "")
		ninja, e12 := NewGentooDependency(">=dev-util/ninja-1.8.2", "")

		fmt.Println(gr)

		It("Check error", func() {
			Expect(e1).Should(BeNil())
			Expect(e2).Should(BeNil())
			Expect(e3).Should(BeNil())
			Expect(e4).Should(BeNil())
			Expect(e5).Should(BeNil())
			Expect(e6).Should(BeNil())
			Expect(e7).Should(BeNil())
			Expect(e8).Should(BeNil())
			Expect(e9).Should(BeNil())
			Expect(e10).Should(BeNil())
			Expect(e11).Should(BeNil())
			Expect(e12).Should(BeNil())
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(*pkgConfig))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(*testUse))
		})

		It("Check dep3", func() {
			Expect(*gr.Dependencies[2]).Should(Equal(*meson))
		})

		It("Check dep4", func() {
			Expect(*gr.Dependencies[3]).Should(Equal(*ninja))
		})

		It("Check dep5", func() {
			Expect(*gr.Dependencies[4]).Should(Equal(*pkgConfig))
		})

	})
	Context("Parse Dependencies 14", func() {

		bdepend := `gnutls? ( net-libs/gnutls[tools] ) !gnutls? ( !libressl? ( dev-libs/openssl:0= ) libressl? ( dev-libs/libressl:0= ) ) >=net-libs/courier-authlib-0.66.4 >=net-libs/courier-unicode-2 >=net-mail/mailbase-0.00-r8 net-dns/libidn:= berkdb? ( sys-libs/db:= ) fam? ( virtual/fam ) gdbm? ( >=sys-libs/gdbm-1.8.0 ) selinux? ( sec-policy/selinux-courier ) !mail-mta/courier !net-mail/bincimap !net-mail/cyrus-imapd !net-mail/uw-imap`

		gr, err := ParseDependencies(bdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(14))
		})

		gnutls, e1 := NewGentooDependency("net-libs/gnutls", "")
		gnutlsUse, e2 := NewGentooDependencyWithSubdeps("", "gnutls",
			[]*GentooDependency{gnutls})

		openssl, e3 := NewGentooDependency("dev-libs/openssl:0=", "")
		libresslUseNot, e4 := NewGentooDependencyWithSubdeps("", "libressl",
			[]*GentooDependency{openssl})
		libresslUseNot.UseCondition = _gentoo.PkgCondNot

		libressl, e5 := NewGentooDependency("dev-libs/libressl:0=", "")
		libresslUse, e6 := NewGentooDependencyWithSubdeps("", "libressl",
			[]*GentooDependency{libressl})

		gnutlsNot, e7 := NewGentooDependencyWithSubdeps("", "!gnutls",
			[]*GentooDependency{libresslUseNot, libresslUse})

		courierAuthlib, e8 := NewGentooDependency(">=net-libs/courier-authlib-0.66.4", "")
		courierUnicode, e9 := NewGentooDependency(">=net-libs/courier-unicode-2", "")
		mailbase, e10 := NewGentooDependency(">=net-mail/mailbase-0.00-r8", "")
		libidn, e11 := NewGentooDependency("net-dns/libidn:=", "")

		db, e12 := NewGentooDependency("sys-libs/db:=", "")
		berkdbUse, e13 := NewGentooDependencyWithSubdeps("", "berkdb",
			[]*GentooDependency{db})

		fam, e14 := NewGentooDependency("virtual/fam", "")
		famUse, e15 := NewGentooDependencyWithSubdeps("", "fam",
			[]*GentooDependency{fam})

		gdbm, e16 := NewGentooDependency(">=sys-libs/gdbm-1.8.0", "")
		gdbmUse, e17 := NewGentooDependencyWithSubdeps("", "gdbm",
			[]*GentooDependency{gdbm})

		selinuxCourier, e18 := NewGentooDependency("sec-policy/selinux-courier", "")
		selinuxUse, e19 := NewGentooDependencyWithSubdeps("", "selinux",
			[]*GentooDependency{selinuxCourier})

		courier, e20 := NewGentooDependency("!mail-mta/courier", "")
		bincimap, e21 := NewGentooDependency("!net-mail/bincimap", "")
		imapd, e22 := NewGentooDependency("!net-mail/cyrus-imapd", "")
		uwimap, e23 := NewGentooDependency("!net-mail/uw-imap", "")

		It("Check error", func() {
			Expect(e1).Should(BeNil())
			Expect(e2).Should(BeNil())
			Expect(e3).Should(BeNil())
			Expect(e4).Should(BeNil())
			Expect(e5).Should(BeNil())
			Expect(e6).Should(BeNil())
			Expect(e7).Should(BeNil())
			Expect(e8).Should(BeNil())
			Expect(e9).Should(BeNil())
			Expect(e10).Should(BeNil())
			Expect(e11).Should(BeNil())
			Expect(e12).Should(BeNil())
			Expect(e13).Should(BeNil())
			Expect(e14).Should(BeNil())
			Expect(e15).Should(BeNil())
			Expect(e16).Should(BeNil())
			Expect(e17).Should(BeNil())
			Expect(e18).Should(BeNil())
			Expect(e19).Should(BeNil())
			Expect(e20).Should(BeNil())
			Expect(e21).Should(BeNil())
			Expect(e22).Should(BeNil())
			Expect(e23).Should(BeNil())
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(*gnutlsUse))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(*gnutlsNot))
		})

		It("Check dep3", func() {
			Expect(*gr.Dependencies[2]).Should(Equal(*courierAuthlib))
		})

		It("Check dep4", func() {
			Expect(*gr.Dependencies[3]).Should(Equal(*courierUnicode))
		})

		It("Check dep5", func() {
			Expect(*gr.Dependencies[4]).Should(Equal(*mailbase))
		})

		It("Check dep6", func() {
			Expect(*gr.Dependencies[5]).Should(Equal(*libidn))
		})

		It("Check dep7", func() {
			Expect(*gr.Dependencies[6]).Should(Equal(*berkdbUse))
		})

		It("Check dep8", func() {
			Expect(*gr.Dependencies[7]).Should(Equal(*famUse))
		})

		It("Check dep9", func() {
			Expect(*gr.Dependencies[8]).Should(Equal(*gdbmUse))
		})

		It("Check dep10", func() {
			Expect(*gr.Dependencies[9]).Should(Equal(*selinuxUse))
		})

		It("Check dep11", func() {
			Expect(*gr.Dependencies[10]).Should(Equal(*courier))
		})

		It("Check dep12", func() {
			Expect(*gr.Dependencies[11]).Should(Equal(*bincimap))
		})

		It("Check dep13", func() {
			Expect(*gr.Dependencies[12]).Should(Equal(*imapd))
		})

		It("Check dep14", func() {
			Expect(*gr.Dependencies[13]).Should(Equal(*uwimap))
		})
		fmt.Println(gr)
	})

	Context("Parse Dependencies 15", func() {

		bdepend := `|| ( dev-lang/python:3.9 dev-lang/python:3.8 dev-lang/python:3.7 dev-lang/python:3.6 ) opencl? ( >=sys-devel/gcc-4.6 ) sys-devel/bison sys-devel/flex sys-devel/gettext virtual/pkgconfig || ( ( dev-lang/python:3.9 >=dev-python/mako-0.8.0[python_targets_python3_9(-),python_single_target_python3_9(+)] ) ( dev-lang/python:3.8 >=dev-python/mako-0.8.0[python_targets_python3_8(-),python_single_target_python3_8(+)] ) ( dev-lang/python:3.7 >=dev-python/mako-0.8.0[python_targets_python3_7(-),python_single_target_python3_7(+)] ) ( dev-lang/python:3.6 >=dev-python/mako-0.8.0[python_targets_python3_6(-),python_single_target_python3_6(+)] ) ) >=dev-util/meson-0.54.0 >=dev-util/ninja-1.8.2`

		gr, err := ParseDependencies(bdepend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		py39, e1 := NewGentooDependency("dev-lang/python:3.9", "")
		py38, e2 := NewGentooDependency("dev-lang/python:3.8", "")
		py37, e3 := NewGentooDependency("dev-lang/python:3.7", "")
		py36, e4 := NewGentooDependency("dev-lang/python:3.6", "")

		orPy, e5 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{py39, py38, py37, py36})
		orPy.DepInOr = true

		gcc, e6 := NewGentooDependency(">=sys-devel/gcc-4.6", "")
		openclUse, e7 := NewGentooDependencyWithSubdeps("", "opencl",
			[]*GentooDependency{gcc})

		bison, e7 := NewGentooDependency("sys-devel/bison", "")
		flex, e8 := NewGentooDependency("sys-devel/flex", "")
		getText, e9 := NewGentooDependency("sys-devel/gettext", "")
		pkgConfig, e10 := NewGentooDependency("virtual/pkgconfig", "")

		mako, e11 := NewGentooDependency(">=dev-python/mako-0.8.0", "")

		orMakopy39, e12 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{py39, mako})

		orMakopy38, e13 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{py38, mako})

		orMakopy37, e14 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{py37, mako})

		orMakopy36, e15 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{py36, mako})

		orMako, e16 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{orMakopy39, orMakopy38, orMakopy37, orMakopy36})
		orMako.DepInOr = true

		meson, e17 := NewGentooDependency(">=dev-util/meson-0.54.0", "")
		ninja, e18 := NewGentooDependency(">=dev-util/ninja-1.8.2", "")

		It("Check error", func() {
			Expect(e1).Should(BeNil())
			Expect(e2).Should(BeNil())
			Expect(e3).Should(BeNil())
			Expect(e4).Should(BeNil())
			Expect(e5).Should(BeNil())
			Expect(e6).Should(BeNil())
			Expect(e7).Should(BeNil())
			Expect(e8).Should(BeNil())
			Expect(e9).Should(BeNil())
			Expect(e10).Should(BeNil())
			Expect(e11).Should(BeNil())
			Expect(e12).Should(BeNil())
			Expect(e13).Should(BeNil())
			Expect(e14).Should(BeNil())
			Expect(e15).Should(BeNil())
			Expect(e16).Should(BeNil())
			Expect(e17).Should(BeNil())
			Expect(e18).Should(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(9))
		})

		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(*orPy))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(*openclUse))
		})

		It("Check dep3", func() {
			Expect(*gr.Dependencies[2]).Should(Equal(*bison))
		})

		It("Check dep4", func() {
			Expect(*gr.Dependencies[3]).Should(Equal(*flex))
		})

		It("Check dep5", func() {
			Expect(*gr.Dependencies[4]).Should(Equal(*getText))
		})

		It("Check dep6", func() {
			Expect(*gr.Dependencies[5]).Should(Equal(*pkgConfig))
		})

		It("Check dep7", func() {
			Expect(*gr.Dependencies[6]).Should(Equal(*orMako))
		})

		It("Check dep8", func() {
			Expect(*gr.Dependencies[7]).Should(Equal(*meson))
		})

		It("Check dep9", func() {
			Expect(*gr.Dependencies[8]).Should(Equal(*ninja))
		})
		fmt.Println(gr)
	})

	Context("Parse Dependencies 16", func() {

		depend := `acct-group/sshd acct-user/sshd !static? ( audit? ( sys-process/audit ) ldns? ( net-libs/ldns !bindist? ( net-libs/ldns[ecdsa,ssl(+)] ) bindist? ( net-libs/ldns[-ecdsa,ssl(+)]
		       ) ) libedit? ( dev-libs/libedit:= ) sctp? ( net-misc/lksctp-tools ) selinux? ( >=sys-libs/libselinux-1.28 ) ssl? ( !libressl? ( || ( ( >=dev-libs/openssl-1.0.1:0[bindist=] <dev-libs/openssl-1.1.0:0[bindist=] ) >=dev-libs/openssl-1.1.0g:0[bindist=] ) dev-libs/openssl:0= ) libressl? ( dev-libs/libressl:0= ) ) virtual/libcrypt:= >=sys-libs/zlib-1.2.3:= ) pam? ( sys-libs/pam ) kerberos? ( virtual/krb5 ) virtual/os-headers kernel_linux? ( >=sys-kernel/linux-headers-5.1 ) static? ( audit? ( sys-process/audit[static-libs(+)] ) ldns? ( net-libs/ldns[static-libs(+)] !bindist? ( net-libs/ldns[ecdsa,ssl(+)] ) bindist? ( net-libs/ldns[-ecdsa,ssl(+)] ) ) libedit? ( dev-libs/libedit:=[static-libs(+)] ) sctp? ( net-misc/lksctp-tools[static-libs(+)] ) selinux? ( >=sys-libs/libselinux-1.28[static-libs(+)] ) ssl? ( !libressl? ( || ( ( >=dev-libs/openssl-1.0.1:0[bindist=] <dev-libs/openssl-1.1.0:0[bindist=] ) >=dev-libs/openssl-1.1.0g:0[bindist=] ) dev-libs/openssl:0=[static-libs(+)] ) libressl? ( dev-libs/libressl:0=[static-libs(+)] ) ) virtual/libcrypt:=[static-libs(+)] >=sys-libs/zlib-1.2.3:=[static-libs(+)] )`

		gr, err := ParseDependencies(depend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(8))
		})

		groupSshd, e1 := NewGentooDependency("acct-group/sshd", "")
		userSshd, e2 := NewGentooDependency("acct-user/sshd", "")

		audit, e3 := NewGentooDependency("sys-process/audit", "")
		auditUse, e3 := NewGentooDependencyWithSubdeps("", "audit",
			[]*GentooDependency{audit})

		ldns, e4 := NewGentooDependency("net-libs/ldns", "")
		bindist, e5 := NewGentooDependencyWithSubdeps("", "bindist",
			[]*GentooDependency{ldns})
		bindistNot, e6 := NewGentooDependencyWithSubdeps("", "bindist",
			[]*GentooDependency{ldns})
		bindistNot.UseCondition = _gentoo.PkgCondNot

		ldnsUse, e7 := NewGentooDependencyWithSubdeps("", "ldns",
			[]*GentooDependency{ldns, bindistNot, bindist})

		libedit, e8 := NewGentooDependency("dev-libs/libedit:=", "")
		libeditUse, e9 := NewGentooDependencyWithSubdeps("", "libedit",
			[]*GentooDependency{libedit})

		sctp, e10 := NewGentooDependency("net-misc/lksctp-tools", "")
		sctpUse, e11 := NewGentooDependencyWithSubdeps("", "sctp", []*GentooDependency{sctp})

		selinux, e12 := NewGentooDependency(">=sys-libs/libselinux-1.28", "")
		selinuxUse, e13 := NewGentooDependencyWithSubdeps("", "selinux", []*GentooDependency{selinux})

		openssl2, e14 := NewGentooDependency(">=dev-libs/openssl-1.0.1:0", "")
		openssl3, e15 := NewGentooDependency("<dev-libs/openssl-1.1.0:0", "")
		openssl1, e16 := NewGentooDependencyWithSubdeps("", "", []*GentooDependency{openssl2, openssl3})
		opensslg, e17 := NewGentooDependency(">=dev-libs/openssl-1.1.0g:0", "")
		openssl0, e18 := NewGentooDependency("dev-libs/openssl:0=", "")

		opensslOr, e19 := NewGentooDependencyWithSubdeps("", "",
			[]*GentooDependency{openssl1, opensslg})
		opensslOr.DepInOr = true

		libresslNot, e20 := NewGentooDependencyWithSubdeps("", "libressl",
			[]*GentooDependency{opensslOr, openssl0})
		libresslNot.UseCondition = _gentoo.PkgCondNot

		libressl, e21 := NewGentooDependency("dev-libs/libressl:0=", "")
		libresslUse, e22 := NewGentooDependencyWithSubdeps("", "libressl",
			[]*GentooDependency{libressl})

		ssl, e23 := NewGentooDependencyWithSubdeps("", "ssl",
			[]*GentooDependency{libresslNot, libresslUse})

		libgcrypt, e24 := NewGentooDependency("virtual/libcrypt:=", "")
		zlib, e25 := NewGentooDependency(">=sys-libs/zlib-1.2.3:=", "")
		static, e26 := NewGentooDependencyWithSubdeps("", "static",
			[]*GentooDependency{
				auditUse,
				ldnsUse,
				libeditUse,
				sctpUse,
				selinuxUse,
				ssl,
				libgcrypt,
				zlib,
			})
		staticNot, e27 := NewGentooDependencyWithSubdeps("", "static",
			[]*GentooDependency{
				auditUse,
				ldnsUse,
				libeditUse,
				sctpUse,
				selinuxUse,
				ssl,
				libgcrypt,
				zlib,
			})
		staticNot.UseCondition = _gentoo.PkgCondNot
		pam, e28 := NewGentooDependency("sys-libs/pam", "")
		pamUse, e29 := NewGentooDependencyWithSubdeps("", "pam",
			[]*GentooDependency{pam})

		kerberos, e30 := NewGentooDependency("virtual/krb5", "")
		kerberosUse, e31 := NewGentooDependencyWithSubdeps("", "kerberos",
			[]*GentooDependency{kerberos})

		osHeader, e32 := NewGentooDependency("virtual/os-headers", "")

		kernelLinux, e33 := NewGentooDependency(">=sys-kernel/linux-headers-5.1", "")
		kernelLinuxUse, e34 := NewGentooDependencyWithSubdeps("", "kernel_linux",
			[]*GentooDependency{kernelLinux})

		It("Check error", func() {
			Expect(e1).Should(BeNil())
			Expect(e2).Should(BeNil())
			Expect(e3).Should(BeNil())
			Expect(e4).Should(BeNil())
			Expect(e5).Should(BeNil())
			Expect(e6).Should(BeNil())
			Expect(e7).Should(BeNil())
			Expect(e8).Should(BeNil())
			Expect(e9).Should(BeNil())
			Expect(e10).Should(BeNil())
			Expect(e11).Should(BeNil())
			Expect(e12).Should(BeNil())
			Expect(e13).Should(BeNil())
			Expect(e14).Should(BeNil())
			Expect(e15).Should(BeNil())
			Expect(e16).Should(BeNil())
			Expect(e17).Should(BeNil())
			Expect(e18).Should(BeNil())
			Expect(e19).Should(BeNil())
			Expect(e20).Should(BeNil())
			Expect(e21).Should(BeNil())
			Expect(e22).Should(BeNil())
			Expect(e23).Should(BeNil())
			Expect(e24).Should(BeNil())
			Expect(e25).Should(BeNil())
			Expect(e26).Should(BeNil())
			Expect(e27).Should(BeNil())
			Expect(e28).Should(BeNil())
			Expect(e29).Should(BeNil())
			Expect(e30).Should(BeNil())
			Expect(e31).Should(BeNil())
			Expect(e32).Should(BeNil())
			Expect(e33).Should(BeNil())
			Expect(e34).Should(BeNil())
		})
		It("Check dep1", func() {
			Expect(*gr.Dependencies[0]).Should(Equal(*groupSshd))
		})

		It("Check dep2", func() {
			Expect(*gr.Dependencies[1]).Should(Equal(*userSshd))
		})

		It("Check dep3", func() {
			//			Expect(*gr.Dependencies[2].SubDeps[2]).Should(Equal(*staticNot.SubDeps[2]))
			Expect(*gr.Dependencies[2]).Should(Equal(*staticNot))
		})

		It("Check dep4", func() {
			Expect(*gr.Dependencies[3]).Should(Equal(*pamUse))
		})

		It("Check dep5", func() {
			Expect(*gr.Dependencies[4]).Should(Equal(*kerberosUse))
		})

		It("Check dep6", func() {
			Expect(*gr.Dependencies[5]).Should(Equal(*osHeader))
		})

		It("Check dep7", func() {
			Expect(*gr.Dependencies[6]).Should(Equal(*kernelLinuxUse))
		})

		It("Check dep8", func() {
			Expect(*gr.Dependencies[7]).Should(Equal(*static))
		})

	})

	Context("Parse Dependencies 17", func() {

		depend := `|| ( fam? ( sys-fs/fam ) fam2? ( sys-fs/fam ) virtual/libc )`

		gr, err := ParseDependencies(depend)
		It("Check error", func() {
			Expect(err).Should(BeNil())
		})
		It("Check gr", func() {
			Expect(gr).ShouldNot(BeNil())
		})

		It("Check deps #", func() {
			Expect(len(gr.Dependencies)).Should(Equal(1))
		})

	})
})
