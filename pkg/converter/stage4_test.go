/*
Copyright © 2021-2026 Macaroni OS Linux
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

/*
func NewPackage(name, cat, version string, deps []*anise_pkg.DefaultPackage) *anise_pkg.DefaultPackage {
	return &anise_pkg.DefaultPackage{
		Name:            name,
		Category:        cat,
		Version:         version,
		PackageRequires: deps,
	}
}
*/

var _ = Describe("Converter", func() {
	Context("Stage4", func() {

		fmt.Println("TEST")
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
		It("Drop dependency1", func() {

			p := NewPackage("C", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("D", "cat1", ">=0", nil),
					NewPackage("A", "cat1", ">=0", nil),
				},
			)

			err := RemoveDependencyFromLuetPackage(p, p1)
			Expect(err).Should(BeNil())
			Expect(len(p.GetRequires())).Should(Equal(1))
			Expect(p.GetRequires()[0].GetName()).Should(Equal(p4.GetName()))
		})
		It("Drop dependency2", func() {

			p := NewPackage("C", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("D", "cat1", ">=0", nil),
					NewPackage("A", "cat1", ">=0", nil),
				},
			)
			err := RemoveDependencyFromLuetPackage(p, p4)
			Expect(err).Should(BeNil())
			Expect(len(p.GetRequires())).Should(Equal(1))
			Expect(p.GetRequires()[0].GetName()).Should(Equal(p1.GetName()))
		})

		It("Drop dependency3", func() {

			p := NewPackage("C", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("D", "cat1", ">=0", nil),
					NewPackage("A", "cat1", ">=0", nil),
				},
			)
			err := RemoveDependencyFromLuetPackage(p, p2)
			Expect(err).ShouldNot(BeNil())
		})

		It("Drop leaf in tail", func() {
			levels := NewStage4Levels()
			tree1 := NewStage4Tree(1)
			levels.AddTree(tree1)

			err1 := tree1.AddDependency(p2, p1)
			err2 := tree1.AddDependency(p3, p1)

			err3 := tree1.DropDependency(p3)

			Expect(err1).Should(BeNil())
			Expect(err2).Should(BeNil())
			Expect(err3).Should(BeNil())

			Expect(len((*tree1.GetDeps()))).Should(Equal(1))
			Expect((*tree1.GetDeps())[0]).Should(Equal(p2))
		})

		It("Check error on Drop dependency", func() {
			levels := NewStage4Levels()
			tree1 := NewStage4Tree(1)
			levels.AddTree(tree1)

			err1 := tree1.AddDependency(p2, p1)
			err2 := tree1.AddDependency(p3, p1)
			err3 := tree1.DropDependency(p4)

			Expect(err1).Should(BeNil())
			Expect(err2).Should(BeNil())
			Expect(err3).ShouldNot(BeNil())
		})

		It("Drop leaf in head", func() {
			levels := NewStage4Levels()
			tree1 := NewStage4Tree(1)
			levels.AddTree(tree1)

			err1 := tree1.AddDependency(p2, p1)
			err2 := tree1.AddDependency(p3, p1)

			err3 := tree1.DropDependency(p2)

			Expect(err1).Should(BeNil())
			Expect(err2).Should(BeNil())
			Expect(err3).Should(BeNil())

			Expect(len((*tree1.GetDeps()))).Should(Equal(1))
			Expect((*tree1.GetDeps())[0]).Should(Equal(p3))
		})

		It("Drop leaf in middle", func() {
			levels := NewStage4Levels()
			tree1 := NewStage4Tree(1)
			levels.AddTree(tree1)

			err1 := tree1.AddDependency(p2, p1)
			err2 := tree1.AddDependency(p3, p1)
			err3 := tree1.AddDependency(p4, p1)

			err4 := tree1.DropDependency(p3)

			Expect(err1).Should(BeNil())
			Expect(err2).Should(BeNil())
			Expect(err3).Should(BeNil())
			Expect(err4).Should(BeNil())

			Expect(len((*tree1.GetDeps()))).Should(Equal(2))
			Expect((*tree1.GetDeps())[0]).Should(Equal(p2))
			Expect((*tree1.GetDeps())[1]).Should(Equal(p4))
		})

		It("Test level.AddDependency1", func() {
			errs := make([]error, 0)
			levels := NewStage4LevelsWithSize(2)

			errs = append(errs, levels.AddDependency(p1, nil, 0))
			errs = append(errs, levels.AddDependency(p2, nil, 0))
			errs = append(errs, levels.AddDependency(p3, nil, 0))
			errs = append(errs, levels.AddDependency(p4, nil, 0))

			// B (father A )
			errs = append(errs, levels.AddDependency(p2, p1, 1))
			// C (father A)
			errs = append(errs, levels.AddDependency(p3, p1, 1))
			// D (father B)
			errs = append(errs, levels.AddDependency(p4, p2, 1))
			// C (father B)
			errs = append(errs, levels.AddDependency(p3, p2, 1))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 1))

			for i, _ := range errs {
				Expect(errs[i]).Should(BeNil())
			}

			Expect(len(levels.Map)).Should(Equal(4))
			Expect(len(levels.Levels)).Should(Equal(2))
			Expect(len(levels.Levels[0].Deps)).Should(Equal(4))
			Expect(len(levels.Levels[1].Deps)).Should(Equal(3))
		})

		It("Test level.AddDependency2", func() {
			errs := make([]error, 0)
			levels := NewStage4LevelsWithSize(4)

			errs = append(errs, levels.AddDependency(p1, nil, 0))
			errs = append(errs, levels.AddDependency(p2, nil, 0))
			errs = append(errs, levels.AddDependency(p3, nil, 0))
			errs = append(errs, levels.AddDependency(p4, nil, 0))

			// B (father A )
			errs = append(errs, levels.AddDependency(p2, p1, 1))
			// C (father A)
			errs = append(errs, levels.AddDependency(p3, p1, 1))
			// D (father B)
			errs = append(errs, levels.AddDependency(p4, p2, 1))
			// C (father B)
			errs = append(errs, levels.AddDependency(p3, p2, 1))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 1))

			// D (father B)
			errs = append(errs, levels.AddDependency(p4, p2, 2))
			// C (father B)
			errs = append(errs, levels.AddDependency(p3, p2, 2))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 2))

			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 3))

			for i, _ := range errs {
				Expect(errs[i]).Should(BeNil())
			}

			Expect(len(levels.Map)).Should(Equal(4))
			Expect(len(levels.Levels)).Should(Equal(4))
			Expect(len(levels.Levels[0].Deps)).Should(Equal(4))
			Expect(len(levels.Levels[1].Deps)).Should(Equal(3))
			Expect(len(levels.Levels[2].Deps)).Should(Equal(2))
			Expect(len(levels.Levels[3].Deps)).Should(Equal(1))

			fmt.Println("LEVELS ", levels.Dump())
		})

		It("Load Tree", func() {
			levels := NewStage4Levels()
			tree1 := NewStage4Tree(1)
			tree2 := NewStage4Tree(2)
			tree3 := NewStage4Tree(3)
			levels.AddTree(tree1)
			levels.AddTree(tree2)
			levels.AddTree(tree3)

			err1 := tree1.AddDependency(p2, p1)
			err2 := tree1.AddDependency(p3, p1)

			err3 := tree2.AddDependency(p4, p2)
			err4 := tree2.AddDependency(p3, p2)

			err5 := tree3.AddDependency(p4, p3)

			Expect(err1).Should(BeNil())
			Expect(err2).Should(BeNil())
			Expect(err3).Should(BeNil())
			Expect(err4).Should(BeNil())
			Expect(err5).Should(BeNil())

			Expect(tree1.GetDependency(0)).Should(Equal(p2))
			Expect(tree1.GetDependency(1)).Should(Equal(p3))
			Expect(tree2.GetDependency(0)).Should(Equal(p4))
			Expect(tree2.GetDependency(1)).Should(Equal(p3))
			Expect(tree3.GetDependency(0)).Should(Equal(p4))
		})

		It("Check clean of the father", func() {

			p1 := NewPackage("A", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("B", "cat1", ">=0", nil),
					NewPackage("C", "cat1", ">=0", nil),
				},
			)

			p2 := NewPackage("B", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("C", "cat1", ">=0", nil),
				},
			)

			p3 := NewPackage("C", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("D", "cat1", ">=0", nil),
				},
			)

			p4 := NewPackage("D", "cat1", "1.0", nil)

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

			for i, _ := range levels.Levels {
				Expect(levels.Levels[i].Id).Should(Equal(i + 1))
			}

			for i, _ := range levels.Levels {
				Expect(len(levels.Levels[i].Map)).Should(Equal(4 - i))
			}

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3, p4},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2, p3, p4},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2, p3, p4},
			))

			// Check Maps
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 1,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 1,
						Counter:  2,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 2,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 3,
						Counter:  1,
					},
				},
			))

			fmt.Println("LEVELS\n", levels.Dump())
			_, err := levels.AnalyzeLeaf(3, levels.Levels[3],
				levels.Levels[3].Map["cat1/D"],
			)
			Expect(err).Should(BeNil())

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2, p3},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2, p3},
			))

			// Check Maps
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 1,
						Counter:  2,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/C level 3")
			fmt.Println("====================================")
			_, err = levels.AnalyzeLeaf(2, levels.Levels[2],
				levels.Levels[2].Map["cat1/C"],
			)
			Expect(err).Should(BeNil())
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2},
			))

			// Check Maps
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/B level 2")
			fmt.Println("====================================")
			_, err = levels.AnalyzeLeaf(1, levels.Levels[1],
				levels.Levels[1].Map["cat1/B"],
			)
			Expect(err).Should(BeNil())
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1},
			))

			// Check Maps
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/A level 1")
			fmt.Println("====================================")
			_, err = levels.AnalyzeLeaf(0, levels.Levels[0],
				levels.Levels[0].Map["cat1/A"],
			)
			Expect(err).Should(BeNil())
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1},
			))

			// Check Maps
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
				},
			))

		})

		It("Resolve Levels test1", func() {

			errs := make([]error, 0)

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

			levels := NewStage4LevelsWithSize(4)

			// Level1
			// A (deps B, C)
			errs = append(errs, levels.AddDependency(p1, nil, 0))
			// B (deps D, C)
			errs = append(errs, levels.AddDependency(p2, nil, 0))
			// C (dep D)
			errs = append(errs, levels.AddDependency(p3, nil, 0))
			// D
			errs = append(errs, levels.AddDependency(p4, nil, 0))

			// Level2
			// B (father A)
			errs = append(errs, levels.AddDependency(p2, p1, 1))
			// C (father A)
			errs = append(errs, levels.AddDependency(p3, p1, 1))
			// D (father B)
			errs = append(errs, levels.AddDependency(p4, p2, 1))
			// C (father B)
			errs = append(errs, levels.AddDependency(p3, p2, 1))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 1))

			// Level3
			// D (father B)
			errs = append(errs, levels.AddDependency(p4, p2, 2))
			// C (father B)
			errs = append(errs, levels.AddDependency(p3, p2, 2))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 2))

			// Level4
			errs = append(errs, levels.AddDependency(p4, p3, 3))

			for i, _ := range errs {
				Expect(errs[i]).Should(BeNil())
			}

			for i, _ := range levels.Levels {
				Expect(levels.Levels[i].Id).Should(Equal(i + 1))
			}

			for i, _ := range levels.Levels {
				Expect(len(levels.Levels[i].Map)).Should(Equal(4 - i))
			}

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4, p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2, p3, p4},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2, p3, p4},
			))

			// Check Maps
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 1,
						Counter:  1,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p2, p3},
						Position: 0,
						Counter:  2,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 1,
						Counter:  2,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p2, p3},
						Position: 2,
						Counter:  2,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 3,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/D level 4")
			fmt.Println("====================================")

			fmt.Println("LEVELS\n", levels.Dump())

			_, err := levels.AnalyzeLeaf(3, levels.Levels[3],
				levels.Levels[3].Map["cat1/D"],
			)
			Expect(err).Should(BeNil())

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2, p3},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2, p3},
			))

			// Check Maps
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 1,
						Counter:  2,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/C level 3")
			fmt.Println("====================================")

			fmt.Println("LEVELS\n", levels.Dump())

			_, err = levels.AnalyzeLeaf(2, levels.Levels[2],
				levels.Levels[2].Map["cat1/C"],
			)
			Expect(err).Should(BeNil())

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2},
			))

			// Check Maps
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/B level 2")
			fmt.Println("====================================")

			fmt.Println("LEVELS\n", levels.Dump())

			_, err = levels.AnalyzeLeaf(1, levels.Levels[1],
				levels.Levels[1].Map["cat1/B"],
			)
			Expect(err).Should(BeNil())

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1},
			))

			// Check Maps
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
				},
			))
		})

		It("Resolve Levels test2", func() {

			errs := make([]error, 0)

			p1 := NewPackage("A", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("C", "cat1", ">=0", nil),
				},
			)

			p2 := NewPackage("B", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("C", "cat1", ">=0", nil),
				},
			)

			p3 := NewPackage("C", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("D", "cat1", ">=0", nil),
				},
			)

			p4 := NewPackage("D", "cat1", "1.0", nil)

			levels := NewStage4LevelsWithSize(4)

			// Level1
			// A (deps C)
			errs = append(errs, levels.AddDependency(p1, nil, 0))
			// B (deps C)
			errs = append(errs, levels.AddDependency(p2, nil, 0))
			// C (deps D)
			errs = append(errs, levels.AddDependency(p3, nil, 0))
			// D
			errs = append(errs, levels.AddDependency(p4, nil, 0))

			// Level2
			// C (father A)
			errs = append(errs, levels.AddDependency(p3, p1, 1))
			// C (father B)
			errs = append(errs, levels.AddDependency(p3, p2, 1))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 1))

			// Leve3
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 2))

			for i, _ := range errs {
				Expect(errs[i]).Should(BeNil())
			}

			fmt.Println("LEVELS\n", levels.Dump())

			Expect(len(levels.Levels[0].Map)).Should(Equal(4))
			Expect(len(levels.Levels[1].Map)).Should(Equal(2))
			Expect(len(levels.Levels[2].Map)).Should(Equal(1))

			// Check Deps
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3, p4},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2, p3, p4},
			))

			// Check Maps
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 0,
						Counter:  2,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 1,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 3,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/D level 3")
			fmt.Println("====================================")
			rescan, err := levels.AnalyzeLeaf(2, levels.Levels[2],
				levels.Levels[2].Map["cat1/D"],
			)

			Expect(err).Should(BeNil())

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			Expect(rescan).Should(Equal(false))
			// Check Deps
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2, p3},
			))

			// Check Maps
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 0,
						Counter:  2,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/C level 2")
			fmt.Println("====================================")
			rescan, err = levels.AnalyzeLeaf(1, levels.Levels[1],
				levels.Levels[1].Map["cat1/C"],
			)
			Expect(err).Should(BeNil())

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			Expect(rescan).Should(Equal(true))
			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4, p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3, p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2},
			))

			// Check Maps

			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 1,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 1,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/D level 4")
			fmt.Println("====================================")
			rescan, err = levels.AnalyzeLeaf(3, levels.Levels[3],
				levels.Levels[3].Map["cat1/D"],
			)

			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(false))
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3, p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2},
			))

			// Check Maps

			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 1,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/C level 3")
			fmt.Println("====================================")
			rescan, err = levels.AnalyzeLeaf(2, levels.Levels[2],
				levels.Levels[2].Map["cat1/C"],
			)
			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(false))

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2},
			))

			// Check Maps

			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/A level 2")
			fmt.Println("====================================")
			rescan, err = levels.AnalyzeLeaf(1, levels.Levels[1],
				levels.Levels[1].Map["cat1/A"],
			)
			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(false))

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2},
			))

			// Check Maps

			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
				},
			))

			fmt.Println("====================================")
			fmt.Println("ANALYZE cat1/B level 1")
			fmt.Println("====================================")
			rescan, err = levels.AnalyzeLeaf(0, levels.Levels[0],
				levels.Levels[0].Map["cat1/B"],
			)
			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(false))

			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2},
			))

			// Check Maps

			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
				},
			))

		})

		It("Resolve Levels test3", func() {

			levels := NewStage4LevelsWithSize(5)

			p1 := NewPackage("A", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("E", "cat1", ">=0", nil),
					NewPackage("C", "cat1", ">=0", nil),
				},
			)

			p2 := NewPackage("B", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("C", "cat1", ">=0", nil),
				},
			)

			p3 := NewPackage("C", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("D", "cat1", ">=0", nil),
				},
			)

			p4 := NewPackage("D", "cat1", "1.0", nil)

			p5 := NewPackage("E", "cat1", "1.0",
				[]*anise_pkg.DefaultPackage{
					NewPackage("D", "cat1", ">=0", nil),
				},
			)

			errs := make([]error, 0)

			// Level 1
			// A (deps E, C)
			errs = append(errs, levels.AddDependency(p1, nil, 0))
			// B (deps C)
			errs = append(errs, levels.AddDependency(p2, nil, 0))
			// C (deps D)
			errs = append(errs, levels.AddDependency(p3, nil, 0))
			// D
			errs = append(errs, levels.AddDependency(p4, nil, 0))
			// E (deps D)
			errs = append(errs, levels.AddDependency(p5, nil, 0))

			// Level2
			// E (father A)
			errs = append(errs, levels.AddDependency(p5, p1, 1))
			// C (father A)
			errs = append(errs, levels.AddDependency(p3, p1, 1))
			// C (father B)
			errs = append(errs, levels.AddDependency(p3, p2, 1))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 1))
			// D (father E)
			errs = append(errs, levels.AddDependency(p4, p5, 1))

			// Level2
			// D (father E)
			errs = append(errs, levels.AddDependency(p4, p5, 2))
			// D (father C)
			errs = append(errs, levels.AddDependency(p4, p3, 2))

			for i, _ := range errs {
				Expect(errs[i]).Should(BeNil())
			}

			fmt.Println("LEVELS\n", levels.Dump())
			// Check Deps
			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{},
			))
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5, p3, p4},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2, p3, p4, p5},
			))

			// Check Maps
			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5, p3},
						Position: 0,
						Counter:  2,
					}},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 1,
						Counter:  2,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p3, p5},
						Position: 2,
						Counter:  2,
					},
				},
			))
			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 3,
						Counter:  1,
					},
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 4,
						Counter:  1,
					},
				},
			))

			key := "cat1/D"
			key_level := 3
			fmt.Println("====================================")
			fmt.Println(fmt.Sprintf("ANALYZE %s level %d", key, key_level))
			fmt.Println("====================================")
			rescan, err := levels.AnalyzeLeaf(key_level-1, levels.Levels[key_level-1],
				levels.Levels[key_level-1].Map[key],
			)
			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(true))
			fmt.Println("RESCAN ", rescan)
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					}},
			))
			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					},

					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 1,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 1,
						Counter:  2,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 3,
						Counter:  1,
					},
				},
			))

			key = "cat1/D"
			key_level = 4
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
			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{},
			))
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5, p3},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2, p3, p5},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 1,
						Counter:  2,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 3,
						Counter:  1,
					},
				},
			))

			key = "cat1/E"
			key_level = 3
			fmt.Println("====================================")
			fmt.Println(fmt.Sprintf("ANALYZE %s level %d", key, key_level))
			fmt.Println("====================================")
			rescan, err = levels.AnalyzeLeaf(key_level-1, levels.Levels[key_level-1],
				levels.Levels[key_level-1].Map[key],
			)
			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(true))
			fmt.Println("RESCAN ", rescan)
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{},
			))
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2, p3},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					}},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1, p2},
						Position: 0,
						Counter:  2,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 2,
						Counter:  1,
					},
				},
			))

			key = "cat1/C"
			key_level = 2
			fmt.Println("====================================")
			fmt.Println(fmt.Sprintf("ANALYZE %s level %d", key, key_level))
			fmt.Println("====================================")
			rescan, err = levels.AnalyzeLeaf(key_level-1, levels.Levels[key_level-1],
				levels.Levels[key_level-1].Map[key],
			)
			Expect(err).Should(BeNil())
			Expect(rescan).Should(Equal(true))
			fmt.Println("RESCAN ", rescan)
			fmt.Println("LEVELS RESOLVED\n", levels.Dump())

			// Check Deps
			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4, p5},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5, p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3, p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					},
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 1,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 1,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 1,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
				},
			))

			key = "cat1/D"
			key_level = 5
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
			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5, p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3, p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					},
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 1,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 1,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
				},
			))

			key = "cat1/E"
			key_level = 4
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
			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3, p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 1,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
				},
			))

			key = "cat1/C"
			key_level = 3
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
			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1, p2},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 1,
						Counter:  1,
					},
				},
			))

			key = "cat1/A"
			key_level = 2
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
			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
				},
			))

			key = "cat1/B"
			key_level = 1
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
			Expect(levels.Levels[4].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p4},
			))
			Expect(levels.Levels[3].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p5},
			))
			Expect(levels.Levels[2].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p3},
			))
			Expect(levels.Levels[1].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p1},
			))
			Expect(levels.Levels[0].Deps).Should(Equal(
				[]*anise_pkg.DefaultPackage{p2},
			))

			Expect(levels.Levels[4].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/D": &Stage4Leaf{
						Package:  p4,
						Father:   []*anise_pkg.DefaultPackage{p5},
						Position: 0,
						Counter:  1,
					},
				},
			))
			Expect(levels.Levels[3].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/E": &Stage4Leaf{
						Package:  p5,
						Father:   []*anise_pkg.DefaultPackage{p3},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[2].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/C": &Stage4Leaf{
						Package:  p3,
						Father:   []*anise_pkg.DefaultPackage{p1},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[1].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/A": &Stage4Leaf{
						Package:  p1,
						Father:   []*anise_pkg.DefaultPackage{p2},
						Position: 0,
						Counter:  1,
					},
				},
			))

			Expect(levels.Levels[0].Map).Should(Equal(
				map[string]*Stage4Leaf{
					"cat1/B": &Stage4Leaf{
						Package:  p2,
						Father:   []*anise_pkg.DefaultPackage{},
						Position: 0,
						Counter:  1,
					},
				},
			))

		})
	})

})
