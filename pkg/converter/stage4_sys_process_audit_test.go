/*
Copyright © 2021-2023 Funtoo Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package converter_test

import (
	"fmt"

	anise_pkg "github.com/macaroni-os/anise/pkg/package"

	. "github.com/macaroni-os/anise-portage-converter/pkg/converter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Converter", func() {
	Context("Stage4 sys-process/audit", func() {

		bzip2 := NewPackage("bzip2", "app-arch", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("build-seed", "layers", ">=0", nil),
			},
		)

		layer_gentoo_portage := NewPackage("gentoo-portage", "layer", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("gentoo-stage3", "layer", ">=0", nil),
			},
		)

		lang_swig := NewPackage("swig", "dev-lang", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("libpcre", "dev-libs-3", ">=0", nil),
			},
		)

		virtual_pkgconfig := NewPackage("pkgconfig", "virtual", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("build-seed", "layers", ">=0", nil),
			},
		)

		autoconf := NewPackage("autoconf", "sys-devel-2.69", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("build-seed", "layers", ">=0", nil),
			},
		)

		linux_headers := NewPackage("linux-headers", "sys-kernel", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("build-seed", "layers", ">=0", nil),
			},
		)

		python_exec := NewPackage("python-exec", "dev-lang-2", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("build-seed", "layers", ">=0", nil),
			},
		)

		mos_overlay_x := NewPackage("mocaccino-overlay-x", "development", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("gentoo-portage", "layer", ">=0", nil),
			},
		)

		build_seed := NewPackage("build-seed", "layers", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("mocaccino-overlay-x", "development", ">=0", nil),
			},
		)

		zlib := NewPackage("zlib", "sys-libs", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("build-seed", "layers", ">=0", nil),
			},
		)

		libpcre := NewPackage("libpcre", "dev-libs-3", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("bzip2", "app-arch", ">=0", nil),
				NewPackage("readline", "sys-libs", ">=0", nil),
				NewPackage("zlib", "sys-libs", ">=0", nil),
			},
		)

		automake := NewPackage("automake", "sys-devel-1.16", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("build-seed", "layers", ">=0", nil),
			},
		)

		libtool := NewPackage("libtool", "sys-devel-2", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("build-seed", "layers", ">=0", nil),
			},
		)

		audit := NewPackage("audit", "sys-process", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("libcap-ng", "sys-libs", ">=0", nil),
				NewPackage("pkgconfig", "virtual", ">=0", nil),
			},
		)

		readline := NewPackage("readline", "sys-libs", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("ncurses", "sys-libs", ">=0", nil),
				NewPackage("pkgconfig", "virtual", ">=0", nil),
			},
		)

		libcap_ng := NewPackage("libcap-ng", "sys-libs", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("python-exec", "dev-lang-2", ">=0", nil),
				NewPackage("swig", "dev-lang", ">=0", nil),
				NewPackage("autoconf", "sys-devel-2.69", ">=0", nil),
				NewPackage("automake", "sys-devel-1.16", ">=0", nil),
				NewPackage("libtool", "sys-devel-2", ">=0", nil),
				NewPackage("linux-headers", "sys-kernel", ">=0", nil),
			},
		)

		gentoo_stage3 := NewPackage("gentoo-stage3", "layer", "1.0",
			[]*anise_pkg.DefaultPackage{},
		)

		ncurses := NewPackage("ncurses", "sys-libs", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("build-seed", "layers", ">=0", nil),
			},
		)

		It("Resolve", func() {
			levels := NewStage4LevelsWithSize(18)
			errs := make([]error, 0)

			errs = append(errs, levels.AddDependency(bzip2, nil, 0))
			errs = append(errs, levels.AddDependency(layer_gentoo_portage, nil, 0))
			errs = append(errs, levels.AddDependency(lang_swig, nil, 0))
			errs = append(errs, levels.AddDependency(virtual_pkgconfig, nil, 0))
			errs = append(errs, levels.AddDependency(autoconf, nil, 0))
			errs = append(errs, levels.AddDependency(linux_headers, nil, 0))
			errs = append(errs, levels.AddDependency(python_exec, nil, 0))
			errs = append(errs, levels.AddDependency(mos_overlay_x, nil, 0))
			errs = append(errs, levels.AddDependency(build_seed, nil, 0))
			errs = append(errs, levels.AddDependency(zlib, nil, 0))
			errs = append(errs, levels.AddDependency(libpcre, nil, 0))
			errs = append(errs, levels.AddDependency(automake, nil, 0))
			errs = append(errs, levels.AddDependency(libtool, nil, 0))
			errs = append(errs, levels.AddDependency(audit, nil, 0))
			errs = append(errs, levels.AddDependency(readline, nil, 0))
			errs = append(errs, levels.AddDependency(libcap_ng, nil, 0))
			errs = append(errs, levels.AddDependency(gentoo_stage3, nil, 0))
			errs = append(errs, levels.AddDependency(ncurses, nil, 0))

			emptyStack := []string{}
			var err error

			_, err = levels.AddDependencyRecursive(build_seed, bzip2, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(gentoo_stage3, layer_gentoo_portage, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(libpcre, lang_swig, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(build_seed, virtual_pkgconfig, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(build_seed, autoconf, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(build_seed, linux_headers, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(build_seed, python_exec, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(layer_gentoo_portage, mos_overlay_x, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(mos_overlay_x, build_seed, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(build_seed, zlib, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(bzip2, libpcre, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(readline, libpcre, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(zlib, libpcre, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(build_seed, automake, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(build_seed, libtool, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(libcap_ng, audit, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(virtual_pkgconfig, audit, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(ncurses, readline, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(virtual_pkgconfig, readline, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(python_exec, libcap_ng, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(lang_swig, libcap_ng, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(autoconf, libcap_ng, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(automake, libcap_ng, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(libtool, libcap_ng, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(linux_headers, libcap_ng, emptyStack, 1)
			errs = append(errs, err)
			_, err = levels.AddDependencyRecursive(build_seed, ncurses, emptyStack, 1)
			errs = append(errs, err)
			for i, _ := range errs {
				if errs[i] != nil {
					Expect(errs[i]).Should(BeNil())
				}
			}

			// Check Deps
			Expect(levels.Levels[9].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					gentoo_stage3,
				},
			))

			Expect(levels.Levels[8].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					gentoo_stage3,
					layer_gentoo_portage,
				},
			))

			Expect(levels.Levels[7].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					gentoo_stage3,
					layer_gentoo_portage,
					mos_overlay_x,
				},
			))

			Expect(levels.Levels[6].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					gentoo_stage3,
					layer_gentoo_portage,
					mos_overlay_x,
					build_seed,
				},
			))

			Expect(levels.Levels[5].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
					mos_overlay_x,
					gentoo_stage3,
					build_seed,
					ncurses,
					virtual_pkgconfig,
				},
			))

			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					gentoo_stage3,
					mos_overlay_x,
					build_seed,
					layer_gentoo_portage,
					bzip2,
					readline,
					zlib,
					ncurses,
					virtual_pkgconfig,
				},
			))

			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
					build_seed,
					ncurses,
					virtual_pkgconfig,
					gentoo_stage3,
					mos_overlay_x,
					libpcre,
					bzip2,
					readline,
					zlib,
				},
			))

			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
					bzip2,
					readline,
					zlib,
					gentoo_stage3,
					layer_gentoo_portage,
					build_seed,
					ncurses,
					virtual_pkgconfig,
					python_exec,
					lang_swig,
					autoconf,
					automake,
					libtool,
					linux_headers,
					libpcre,
				},
			))

			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
					gentoo_stage3,
					libpcre,
					layer_gentoo_portage,
					mos_overlay_x,
					bzip2,
					readline,
					zlib,
					libcap_ng,
					virtual_pkgconfig,
					ncurses,
					python_exec,
					lang_swig,
					autoconf,
					automake,
					libtool,
					linux_headers,
				},
			))

			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					bzip2,
					layer_gentoo_portage,
					lang_swig,
					virtual_pkgconfig,
					autoconf,
					linux_headers,
					python_exec,
					mos_overlay_x,
					build_seed,
					zlib,
					libpcre,
					automake,
					libtool,
					audit,
					readline,
					libcap_ng,
					gentoo_stage3,
					ncurses,
				},
			))

			Expect(levels.Levels[9].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[8].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 0,
						Counter:  1,
					},
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 1,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[7].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 0,
						Counter:  1,
					},
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 1,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 2,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[6].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 0,
						Counter:  1,
					},
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 1,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 2,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{ncurses, virtual_pkgconfig},
						Position: 3,
						Counter:  2,
					},
				},
			))

			Expect(levels.Levels[5].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 0,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 1,
						Counter:  1,
					},
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 2,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{bzip2, zlib, ncurses, virtual_pkgconfig},
						Position: 3,
						Counter:  4,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 4,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 5,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 0,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 1,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{ncurses, virtual_pkgconfig, bzip2, zlib},
						Position: 2,
						Counter:  4,
					},
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 3,
						Counter:  1,
					},
					"app-arch/bzip2": &Stage4Leaf{
						Package:  bzip2,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 4,
						Counter:  1,
					},
					"sys-libs/readline": &Stage4Leaf{
						Package:  readline,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 5,
						Counter:  1,
					},
					"sys-libs/zlib": &Stage4Leaf{
						Package:  zlib,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 6,
						Counter:  1,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 7,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 8,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 0,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package: build_seed,
						Father: []*anise_pkg.DefaultPackage{
							bzip2, zlib, ncurses, virtual_pkgconfig, python_exec,
							autoconf, automake, libtool, linux_headers,
						},
						Position: 1,
						Counter:  9,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 2,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 3,
						Counter:  1,
					},
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 4,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 5,
						Counter:  1,
					},
					"dev-libs-3/libpcre": &Stage4Leaf{
						Package:  libpcre,
						Father:   []*anise_pkg.DefaultPackage{lang_swig},
						Position: 6,
						Counter:  1,
					},
					"app-arch/bzip2": &Stage4Leaf{
						Package:  bzip2,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 7,
						Counter:  1,
					},
					"sys-libs/readline": &Stage4Leaf{
						Package:  readline,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 8,
						Counter:  1,
					},
					"sys-libs/zlib": &Stage4Leaf{
						Package:  zlib,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 9,
						Counter:  1,
					},
				},
			))

			v, _ := levels.Levels[2].Map["development/mocaccino-overlay-x"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  mos_overlay_x,
				Father:   []*anise_pkg.DefaultPackage{build_seed},
				Position: 0,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["layer/gentoo-stage3"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  gentoo_stage3,
				Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["layer/gentoo-portage"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  layer_gentoo_portage,
				Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package: build_seed,
				Father: []*anise_pkg.DefaultPackage{
					bzip2, zlib, virtual_pkgconfig, ncurses, python_exec,
					autoconf, automake, libtool, linux_headers,
				},
				Position: 6,
				Counter:  9,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 13,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 14,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{lang_swig},
				Position: 15,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[2].Map)).Should(Equal(16))

			v, _ = levels.Levels[1].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package: build_seed,
				Father: []*anise_pkg.DefaultPackage{
					bzip2, virtual_pkgconfig, autoconf,
					linux_headers, python_exec, zlib,
					automake, libtool, ncurses,
				},
				Position: 0,
				Counter:  9,
			},
			))

			v, _ = levels.Levels[1].Map["layer/gentoo-stage3"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  gentoo_stage3,
				Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{lang_swig},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["layer/gentoo-portage"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  layer_gentoo_portage,
				Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["development/mocaccino-overlay-x"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  mos_overlay_x,
				Father:   []*anise_pkg.DefaultPackage{build_seed},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/libcap-ng"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libcap_ng,
				Father:   []*anise_pkg.DefaultPackage{audit},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{audit, readline},
				Position: 9,
				Counter:  2,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 13,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 14,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 15,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 16,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[1].Map)).Should(Equal(17))

			v, _ = levels.Levels[0].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 0,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["layer/gentoo-portage"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  layer_gentoo_portage,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["development/mocaccino-overlay-x"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  mos_overlay_x,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  build_seed,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-process/audit"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  audit,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 13,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 14,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/libcap-ng"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libcap_ng,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 15,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["layer/gentoo-stage3"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  gentoo_stage3,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 16,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 17,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[0].Map)).Should(Equal(18))

			fmt.Println("LEVELS\n", levels.Dump())

			key := "layer/gentoo-stage3"
			key_level := 10
			fmt.Println("====================================")
			fmt.Println(fmt.Sprintf("ANALYZE %s level %d", key, key_level))
			fmt.Println("====================================")
			rescan, err := levels.AnalyzeLeaf(key_level-1, levels.Levels[key_level-1],
				levels.Levels[key_level-1].Map[key],
			)
			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(false))
			fmt.Println("RESCAN ", rescan)
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[9].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					gentoo_stage3,
				},
			))

			Expect(levels.Levels[8].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
				},
			))

			Expect(levels.Levels[7].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
					mos_overlay_x,
				},
			))

			Expect(levels.Levels[6].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
					mos_overlay_x,
					build_seed,
				},
			))

			Expect(levels.Levels[5].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
					mos_overlay_x,
					build_seed,
					ncurses,
					virtual_pkgconfig,
				},
			))

			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
					build_seed,
					layer_gentoo_portage,
					bzip2,
					readline,
					zlib,
					ncurses,
					virtual_pkgconfig,
				},
			))

			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
					build_seed,
					ncurses,
					virtual_pkgconfig,
					mos_overlay_x,
					libpcre,
					bzip2,
					readline,
					zlib,
				},
			))

			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
					bzip2,
					readline,
					zlib,
					layer_gentoo_portage,
					build_seed,
					ncurses,
					virtual_pkgconfig,
					python_exec,
					lang_swig,
					autoconf,
					automake,
					libtool,
					linux_headers,
					libpcre,
				},
			))

			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
					libpcre,
					layer_gentoo_portage,
					mos_overlay_x,
					bzip2,
					readline,
					zlib,
					libcap_ng,
					virtual_pkgconfig,
					ncurses,
					python_exec,
					lang_swig,
					autoconf,
					automake,
					libtool,
					linux_headers,
				},
			))

			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					bzip2,
					layer_gentoo_portage,
					lang_swig,
					virtual_pkgconfig,
					autoconf,
					linux_headers,
					python_exec,
					mos_overlay_x,
					build_seed,
					zlib,
					libpcre,
					automake,
					libtool,
					audit,
					readline,
					libcap_ng,
					ncurses,
				},
			))

			Expect(levels.Levels[9].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[8].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[7].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 0,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 1,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[6].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 0,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 1,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{ncurses, virtual_pkgconfig},
						Position: 2,
						Counter:  2,
					},
				},
			))

			Expect(levels.Levels[5].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 0,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 1,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{bzip2, zlib, ncurses, virtual_pkgconfig},
						Position: 2,
						Counter:  4,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 3,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 4,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 0,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{ncurses, virtual_pkgconfig, bzip2, zlib},
						Position: 1,
						Counter:  4,
					},
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 2,
						Counter:  1,
					},
					"app-arch/bzip2": &Stage4Leaf{
						Package:  bzip2,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 3,
						Counter:  1,
					},
					"sys-libs/readline": &Stage4Leaf{
						Package:  readline,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 4,
						Counter:  1,
					},
					"sys-libs/zlib": &Stage4Leaf{
						Package:  zlib,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 5,
						Counter:  1,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 6,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 7,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 0,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package: build_seed,
						Father: []*anise_pkg.DefaultPackage{
							bzip2, zlib, ncurses, virtual_pkgconfig, python_exec,
							autoconf, automake, libtool, linux_headers,
						},
						Position: 1,
						Counter:  9,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 2,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 3,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 4,
						Counter:  1,
					},
					"dev-libs-3/libpcre": &Stage4Leaf{
						Package:  libpcre,
						Father:   []*anise_pkg.DefaultPackage{lang_swig},
						Position: 5,
						Counter:  1,
					},
					"app-arch/bzip2": &Stage4Leaf{
						Package:  bzip2,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 6,
						Counter:  1,
					},
					"sys-libs/readline": &Stage4Leaf{
						Package:  readline,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 7,
						Counter:  1,
					},
					"sys-libs/zlib": &Stage4Leaf{
						Package:  zlib,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 8,
						Counter:  1,
					},
				},
			))

			v, _ = levels.Levels[2].Map["development/mocaccino-overlay-x"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  mos_overlay_x,
				Father:   []*anise_pkg.DefaultPackage{build_seed},
				Position: 0,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["layer/gentoo-portage"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  layer_gentoo_portage,
				Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package: build_seed,
				Father: []*anise_pkg.DefaultPackage{
					bzip2, zlib, virtual_pkgconfig, ncurses, python_exec,
					autoconf, automake, libtool, linux_headers,
				},
				Position: 5,
				Counter:  9,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 13,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{lang_swig},
				Position: 14,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[2].Map)).Should(Equal(15))

			v, _ = levels.Levels[1].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package: build_seed,
				Father: []*anise_pkg.DefaultPackage{
					bzip2, virtual_pkgconfig, autoconf,
					linux_headers, python_exec, zlib,
					automake, libtool, ncurses,
				},
				Position: 0,
				Counter:  9,
			},
			))

			v, _ = levels.Levels[1].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{lang_swig},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["layer/gentoo-portage"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  layer_gentoo_portage,
				Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["development/mocaccino-overlay-x"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  mos_overlay_x,
				Father:   []*anise_pkg.DefaultPackage{build_seed},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/libcap-ng"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libcap_ng,
				Father:   []*anise_pkg.DefaultPackage{audit},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{audit, readline},
				Position: 8,
				Counter:  2,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 13,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 14,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 15,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[1].Map)).Should(Equal(16))

			v, _ = levels.Levels[0].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 0,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["layer/gentoo-portage"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  layer_gentoo_portage,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["development/mocaccino-overlay-x"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  mos_overlay_x,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  build_seed,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-process/audit"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  audit,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 13,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 14,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/libcap-ng"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libcap_ng,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 15,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 16,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[0].Map)).Should(Equal(17))

			fmt.Println("LEVELS\n", levels.Dump())

			key = "layer/gentoo-portage"
			key_level = 9
			fmt.Println("====================================")
			fmt.Println(fmt.Sprintf("ANALYZE %s level %d", key, key_level))
			fmt.Println("====================================")
			rescan, err = levels.AnalyzeLeaf(key_level-1, levels.Levels[key_level-1],
				levels.Levels[key_level-1].Map[key],
			)
			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(false))
			fmt.Println("RESCAN ", rescan)
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[9].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					gentoo_stage3,
				},
			))

			Expect(levels.Levels[8].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
				},
			))

			Expect(levels.Levels[7].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
				},
			))

			Expect(levels.Levels[6].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
					build_seed,
				},
			))

			Expect(levels.Levels[5].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
					build_seed,
					ncurses,
					virtual_pkgconfig,
				},
			))

			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
					build_seed,
					bzip2,
					readline,
					zlib,
					ncurses,
					virtual_pkgconfig,
				},
			))

			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
					ncurses,
					virtual_pkgconfig,
					mos_overlay_x,
					libpcre,
					bzip2,
					readline,
					zlib,
				},
			))

			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
					bzip2,
					readline,
					zlib,
					build_seed,
					ncurses,
					virtual_pkgconfig,
					python_exec,
					lang_swig,
					autoconf,
					automake,
					libtool,
					linux_headers,
					libpcre,
				},
			))

			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
					libpcre,
					mos_overlay_x,
					bzip2,
					readline,
					zlib,
					libcap_ng,
					virtual_pkgconfig,
					ncurses,
					python_exec,
					lang_swig,
					autoconf,
					automake,
					libtool,
					linux_headers,
				},
			))

			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					bzip2,
					lang_swig,
					virtual_pkgconfig,
					autoconf,
					linux_headers,
					python_exec,
					mos_overlay_x,
					build_seed,
					zlib,
					libpcre,
					automake,
					libtool,
					audit,
					readline,
					libcap_ng,
					ncurses,
				},
			))

			Expect(levels.Levels[9].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[8].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[7].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[6].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 0,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{ncurses, virtual_pkgconfig},
						Position: 1,
						Counter:  2,
					},
				},
			))

			Expect(levels.Levels[5].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 0,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{bzip2, zlib, ncurses, virtual_pkgconfig},
						Position: 1,
						Counter:  4,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 2,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 3,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 0,
						Counter:  1,
					},
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{ncurses, virtual_pkgconfig, bzip2, zlib},
						Position: 1,
						Counter:  4,
					},
					"app-arch/bzip2": &Stage4Leaf{
						Package:  bzip2,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 2,
						Counter:  1,
					},
					"sys-libs/readline": &Stage4Leaf{
						Package:  readline,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 3,
						Counter:  1,
					},
					"sys-libs/zlib": &Stage4Leaf{
						Package:  zlib,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 4,
						Counter:  1,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 5,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 6,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layers/build-seed": &Stage4Leaf{
						Package: build_seed,
						Father: []*anise_pkg.DefaultPackage{
							bzip2, zlib, ncurses, virtual_pkgconfig, python_exec,
							autoconf, automake, libtool, linux_headers,
						},
						Position: 0,
						Counter:  9,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 1,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 2,
						Counter:  1,
					},
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 3,
						Counter:  1,
					},
					"dev-libs-3/libpcre": &Stage4Leaf{
						Package:  libpcre,
						Father:   []*anise_pkg.DefaultPackage{lang_swig},
						Position: 4,
						Counter:  1,
					},
					"app-arch/bzip2": &Stage4Leaf{
						Package:  bzip2,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 5,
						Counter:  1,
					},
					"sys-libs/readline": &Stage4Leaf{
						Package:  readline,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 6,
						Counter:  1,
					},
					"sys-libs/zlib": &Stage4Leaf{
						Package:  zlib,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 7,
						Counter:  1,
					},
				},
			))

			v, _ = levels.Levels[2].Map["development/mocaccino-overlay-x"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  mos_overlay_x,
				Father:   []*anise_pkg.DefaultPackage{build_seed},
				Position: 0,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package: build_seed,
				Father: []*anise_pkg.DefaultPackage{
					bzip2, zlib, virtual_pkgconfig, ncurses, python_exec,
					autoconf, automake, libtool, linux_headers,
				},
				Position: 4,
				Counter:  9,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{lang_swig},
				Position: 13,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[2].Map)).Should(Equal(14))

			v, _ = levels.Levels[1].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package: build_seed,
				Father: []*anise_pkg.DefaultPackage{
					bzip2, virtual_pkgconfig, autoconf,
					linux_headers, python_exec, zlib,
					automake, libtool, ncurses,
				},
				Position: 0,
				Counter:  9,
			},
			))

			v, _ = levels.Levels[1].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{lang_swig},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["development/mocaccino-overlay-x"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  mos_overlay_x,
				Father:   []*anise_pkg.DefaultPackage{build_seed},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/libcap-ng"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libcap_ng,
				Father:   []*anise_pkg.DefaultPackage{audit},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{audit, readline},
				Position: 7,
				Counter:  2,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 13,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 14,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[1].Map)).Should(Equal(15))

			v, _ = levels.Levels[0].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 0,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["development/mocaccino-overlay-x"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  mos_overlay_x,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  build_seed,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-process/audit"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  audit,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 13,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/libcap-ng"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libcap_ng,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 14,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 15,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[0].Map)).Should(Equal(16))

			fmt.Println("LEVELS\n", levels.Dump())

			key = "development/mocaccino-overlay-x"
			key_level = 8
			fmt.Println("====================================")
			fmt.Println(fmt.Sprintf("ANALYZE %s level %d", key, key_level))
			fmt.Println("====================================")
			rescan, err = levels.AnalyzeLeaf(key_level-1, levels.Levels[key_level-1],
				levels.Levels[key_level-1].Map[key],
			)
			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(false))
			fmt.Println("RESCAN ", rescan)
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[9].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					gentoo_stage3,
				},
			))

			Expect(levels.Levels[8].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
				},
			))

			Expect(levels.Levels[7].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
				},
			))

			Expect(levels.Levels[6].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
				},
			))

			Expect(levels.Levels[5].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
					ncurses,
					virtual_pkgconfig,
				},
			))

			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
					bzip2,
					readline,
					zlib,
					ncurses,
					virtual_pkgconfig,
				},
			))

			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
					ncurses,
					virtual_pkgconfig,
					libpcre,
					bzip2,
					readline,
					zlib,
				},
			))

			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					bzip2,
					readline,
					zlib,
					build_seed,
					ncurses,
					virtual_pkgconfig,
					python_exec,
					lang_swig,
					autoconf,
					automake,
					libtool,
					linux_headers,
					libpcre,
				},
			))

			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
					libpcre,
					bzip2,
					readline,
					zlib,
					libcap_ng,
					virtual_pkgconfig,
					ncurses,
					python_exec,
					lang_swig,
					autoconf,
					automake,
					libtool,
					linux_headers,
				},
			))

			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					bzip2,
					lang_swig,
					virtual_pkgconfig,
					autoconf,
					linux_headers,
					python_exec,
					build_seed,
					zlib,
					libpcre,
					automake,
					libtool,
					audit,
					readline,
					libcap_ng,
					ncurses,
				},
			))

			Expect(levels.Levels[9].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-stage3": &Stage4Leaf{
						Package:  gentoo_stage3,
						Father:   []*anise_pkg.DefaultPackage{layer_gentoo_portage},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[8].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layer/gentoo-portage": &Stage4Leaf{
						Package:  layer_gentoo_portage,
						Father:   []*anise_pkg.DefaultPackage{mos_overlay_x},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[7].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"development/mocaccino-overlay-x": &Stage4Leaf{
						Package:  mos_overlay_x,
						Father:   []*anise_pkg.DefaultPackage{build_seed},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[6].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{ncurses, virtual_pkgconfig},
						Position: 0,
						Counter:  2,
					},
				},
			))

			Expect(levels.Levels[5].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{bzip2, zlib, ncurses, virtual_pkgconfig},
						Position: 0,
						Counter:  4,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 1,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 2,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layers/build-seed": &Stage4Leaf{
						Package:  build_seed,
						Father:   []*anise_pkg.DefaultPackage{ncurses, virtual_pkgconfig, bzip2, zlib},
						Position: 0,
						Counter:  4,
					},
					"app-arch/bzip2": &Stage4Leaf{
						Package:  bzip2,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 1,
						Counter:  1,
					},
					"sys-libs/readline": &Stage4Leaf{
						Package:  readline,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 2,
						Counter:  1,
					},
					"sys-libs/zlib": &Stage4Leaf{
						Package:  zlib,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 3,
						Counter:  1,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 4,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 5,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"layers/build-seed": &Stage4Leaf{
						Package: build_seed,
						Father: []*anise_pkg.DefaultPackage{
							bzip2, zlib, ncurses, virtual_pkgconfig, python_exec,
							autoconf, automake, libtool, linux_headers,
						},
						Position: 0,
						Counter:  9,
					},
					"sys-libs/ncurses": &Stage4Leaf{
						Package:  ncurses,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 1,
						Counter:  1,
					},
					"virtual/pkgconfig": &Stage4Leaf{
						Package:  virtual_pkgconfig,
						Father:   []*anise_pkg.DefaultPackage{readline},
						Position: 2,
						Counter:  1,
					},
					"dev-libs-3/libpcre": &Stage4Leaf{
						Package:  libpcre,
						Father:   []*anise_pkg.DefaultPackage{lang_swig},
						Position: 3,
						Counter:  1,
					},
					"app-arch/bzip2": &Stage4Leaf{
						Package:  bzip2,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 4,
						Counter:  1,
					},
					"sys-libs/readline": &Stage4Leaf{
						Package:  readline,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 5,
						Counter:  1,
					},
					"sys-libs/zlib": &Stage4Leaf{
						Package:  zlib,
						Father:   []*anise_pkg.DefaultPackage{libpcre},
						Position: 6,
						Counter:  1,
					},
				},
			))

			v, _ = levels.Levels[2].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 0,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package: build_seed,
				Father: []*anise_pkg.DefaultPackage{
					bzip2, zlib, virtual_pkgconfig, ncurses, python_exec,
					autoconf, automake, libtool, linux_headers,
				},
				Position: 3,
				Counter:  9,
			},
			))

			v, _ = levels.Levels[2].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[2].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{lang_swig},
				Position: 12,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[2].Map)).Should(Equal(13))

			v, _ = levels.Levels[1].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package: build_seed,
				Father: []*anise_pkg.DefaultPackage{
					bzip2, virtual_pkgconfig, autoconf,
					linux_headers, python_exec, zlib,
					automake, libtool, ncurses,
				},
				Position: 0,
				Counter:  9,
			},
			))

			v, _ = levels.Levels[1].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{lang_swig},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{libpcre},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/libcap-ng"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libcap_ng,
				Father:   []*anise_pkg.DefaultPackage{audit},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{audit, readline},
				Position: 6,
				Counter:  2,
			},
			))

			v, _ = levels.Levels[1].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{readline},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[1].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{libcap_ng},
				Position: 13,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[1].Map)).Should(Equal(14))

			v, _ = levels.Levels[0].Map["app-arch/bzip2"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  bzip2,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 0,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-lang/swig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  lang_swig,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 1,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["virtual/pkgconfig"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  virtual_pkgconfig,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 2,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-2.69/autoconf"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  autoconf,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 3,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-kernel/linux-headers"]
			Expect(v).Should(Equal(&Stage4Leaf{ ///
				Package:  linux_headers,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 4,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-lang-2/python-exec"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  python_exec,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 5,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["layers/build-seed"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  build_seed,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 6,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/zlib"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  zlib,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 7,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["dev-libs-3/libpcre"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libpcre,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 8,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-1.16/automake"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  automake,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 9,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-devel-2/libtool"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libtool,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 10,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-process/audit"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  audit,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 11,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/readline"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  readline,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 12,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/libcap-ng"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  libcap_ng,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 13,
				Counter:  1,
			},
			))

			v, _ = levels.Levels[0].Map["sys-libs/ncurses"]
			Expect(v).Should(Equal(&Stage4Leaf{
				Package:  ncurses,
				Father:   []*anise_pkg.DefaultPackage{},
				Position: 14,
				Counter:  1,
			},
			))

			Expect(len(levels.Levels[0].Map)).Should(Equal(15))

			err = levels.Resolve()
			Expect(err).Should(BeNil())

			// Check Deps
			Expect(levels.Levels[17].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					gentoo_stage3,
				},
			))

			Expect(levels.Levels[16].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					layer_gentoo_portage,
				},
			))

			Expect(levels.Levels[15].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					mos_overlay_x,
				},
			))

			Expect(levels.Levels[14].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					build_seed,
				},
			))

			Expect(levels.Levels[13].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					ncurses,
				},
			))

			Expect(levels.Levels[12].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					virtual_pkgconfig,
				},
			))

			Expect(levels.Levels[11].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					bzip2,
				},
			))

			Expect(levels.Levels[10].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					zlib,
				},
			))

			Expect(levels.Levels[9].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					python_exec,
				},
			))

			Expect(levels.Levels[8].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					autoconf,
				},
			))

			Expect(levels.Levels[7].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					automake,
				},
			))

			Expect(levels.Levels[6].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					readline,
				},
			))

			Expect(levels.Levels[5].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					libpcre,
				},
			))

			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					libtool,
				},
			))

			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					linux_headers,
				},
			))

			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					lang_swig,
				},
			))

			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					libcap_ng,
				},
			))

			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{
					audit,
				},
			))

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())
		})

	})

})
