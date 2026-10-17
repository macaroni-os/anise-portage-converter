/*
Copyright © 2021-2023 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package converter_test

import (
	//"fmt"
	anise_pkg "github.com/macaroni-os/anise/pkg/package"

	. "github.com/macaroni-os/anise-portage-converter/pkg/converter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func NewPackage(name, cat, version string, deps []*anise_pkg.DefaultPackage) *anise_pkg.DefaultPackage {
	return &anise_pkg.DefaultPackage{
		Name:            name,
		Category:        cat,
		Version:         version,
		PackageRequires: deps,
	}
}

var _ = Describe("Converter", func() {
	Context("Stage4 - PackageHasAncient", func() {

		p1 := NewPackage("A", "cat1", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("B", "cat1", ">=0", nil),
				NewPackage("C", "cat1", ">=0", nil),
			},
		)
		p2 := NewPackage("B", "cat1", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("D", "cat1", ">=0", nil),
				NewPackage("C", "cat1", ">=0", nil),
			},
		)

		p3 := NewPackage("C", "cat1", "1.0",
			[]*anise_pkg.DefaultPackage{
				NewPackage("D", "cat1", ">=0", nil),
			},
		)

		p4 := NewPackage("D", "cat1", "1.0", nil)

		It("HasAncient1", func() {
			levels := NewStage4LevelsWithSize(4)

			errs := make([]error, 0)

			// A (deps B, C)
			errs = append(errs, levels.AddDependency(p1, nil, 0))
			// B (deps C)
			errs = append(errs, levels.AddDependency(p2, nil, 0))
			// C (deps D)
			errs = append(errs, levels.AddDependency(p3, nil, 0))
			// D
			errs = append(errs, levels.AddDependency(p4, nil, 0))

			// Tree2
			// B (father A)
			errs = append(errs, levels.AddDependency(p2, p1, 1))
			// C (father A)
			errs = append(errs, levels.AddDependency(p3, p1, 1))
			// C (father B)
			errs = append(errs, levels.AddDependency(p3, p2, 1))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 1))

			// Tree3
			// C (father B)
			errs = append(errs, levels.AddDependency(p3, p2, 2))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 2))

			// Tree4
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 3))

			for i, _ := range errs {
				Expect(errs[i]).Should(BeNil())
			}

			hasAncient, err := levels.PackageHasAncient(p1, p4, 0)
			Expect(err).Should(BeNil())
			Expect(hasAncient).Should(Equal(true))

			hasAncient, err = levels.PackageHasAncient(p3, p1, 0)
			Expect(err).Should(BeNil())
			Expect(hasAncient).Should(Equal(false))
		})
	})

})
