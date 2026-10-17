/*
	Copyright © 2021-2023 Macaroni OS Linux
	See AUTHORS and LICENSE for the license details and contributors.
*/

package converter

import (
	"fmt"
	"path/filepath"

	"github.com/macaroni-os/anise/anise-build/pkg/v2/solver"
	anise_pkg "github.com/macaroni-os/anise/pkg/package"
	"github.com/macaroni-os/anise/pkg/v2/compiler/types/artifact"
	anise_tree "github.com/macaroni-os/anise/pkg/v2/tree"
)

func (pc *PortageConverter) FindRuntimePackages(pTarget *anise_pkg.DefaultPackage) (anise_pkg.DefaultPackages, error) {
	ans := []*anise_pkg.DefaultPackage{}

	runtimeOpts := solver.NewRuntimeSolverOpts()
	s := solver.NewRuntimeSolver(pc.Config, runtimeOpts)

	// Setup the forest guard for the solver
	s.SetForestGuard(pc.ForestGuard)

	treesIdx, err := pc.ForestGuard.SearchPackage(pTarget)
	if err != nil {
		return nil, err
	}

	if len(treesIdx) == 0 {
		treesIdx, err = pc.ForestGuard.SearchProvides(pTarget)
		if err != nil {
			return nil, err
		}
	}

	if len(treesIdx) == 0 {
		return nil, fmt.Errorf("no packages found for package %s", pTarget.HumanReadableString())
	}

	for _, tree := range treesIdx {
		// Load all availables packages

		versions, _ := tree.GetPackageVersions(pTarget.PackageName())

		for _, version := range versions {

			// Using render engine to read build.yaml
			defFile := filepath.Join(tree.TreePath, tree.BaseDir, version.Path)
			runtimePackage, err := anise_tree.ReadDefinitionFile(defFile)
			if err != nil {
				return nil, err
			}

			ans = append(ans, runtimePackage)

		}

	}

	return anise_pkg.DefaultPackages(ans), nil
}

func (pc *PortageConverter) FindBuildtimePackages(pTarget *anise_pkg.DefaultPackage) (anise_pkg.DefaultPackages, error) {
	ans := []*anise_pkg.DefaultPackage{}

	buildOpts := solver.NewBuildSolverOpts()
	s := solver.NewBuildSolver(pc.Config, buildOpts)

	// Setup the forest guard for the solver
	s.SetForestGuard(pc.ForestGuard)
	s.SetRenderEngine(pc.RenderEngine)

	treesIdx, err := pc.ForestGuard.SearchPackage(pTarget)
	if err != nil {
		return nil, err
	}

	if len(treesIdx) == 0 {
		treesIdx, err = pc.ForestGuard.SearchProvides(pTarget)
		if err != nil {
			return nil, err
		}
	}

	if len(treesIdx) == 0 {
		return nil, fmt.Errorf("no packages found for package %s", pTarget.HumanReadableString())
	}

	for _, tree := range treesIdx {
		// Load all availables packages

		versions, _ := tree.GetPackageVersions(pTarget.PackageName())

		for _, version := range versions {

			cs, pkgPath, err := s.LoadCompilationSpec(tree, version, pTarget)
			if err != nil {
				return nil, fmt.Errorf("error on load compilation specs for %s-%s: %s",
					pTarget.PackageName(), version.Version, err.Error())
			}

			art := artifact.NewPackageArtifact(pkgPath)
			art.CompileSpec = cs.ToSpec()

			ans = append(ans, art.GetPackage())

		}

	}

	return anise_pkg.DefaultPackages(ans), nil
}

func (pc *PortageConverter) FindBuildtimePackage(pTarget *anise_pkg.DefaultPackage) (*anise_pkg.DefaultPackage, error) {
	var ans *anise_pkg.DefaultPackage = nil

	buildOpts := solver.NewBuildSolverOpts()
	s := solver.NewBuildSolver(pc.Config, buildOpts)

	// Setup the forest guard for the solver
	s.SetForestGuard(pc.ForestGuard)
	s.SetRenderEngine(pc.RenderEngine)

	treesIdx, err := pc.ForestGuard.SearchPackage(pTarget)
	if err != nil {
		return nil, err
	}

	if len(treesIdx) == 0 {
		treesIdx, err = pc.ForestGuard.SearchProvides(pTarget)
		if err != nil {
			return nil, err
		}
	}

	if len(treesIdx) == 0 {
		return nil, fmt.Errorf("no packages found for package %s", pTarget.HumanReadableString())
	}

	for _, tree := range treesIdx {
		// Load all availables packages

		versions, _ := tree.GetPackageVersions(pTarget.PackageName())

		for _, version := range versions {

			cs, pkgPath, err := s.LoadCompilationSpec(tree, version, pTarget)
			if err != nil {
				return nil, fmt.Errorf("error on load compilation specs for %s-%s: %s",
					pTarget.PackageName(), version.Version, err.Error())
			}

			art := artifact.NewPackageArtifact(pkgPath)
			art.CompileSpec = cs.ToSpec()

			ans = art.GetPackage()

			break
		}

		if ans != nil {
			break
		}
	}

	return ans, nil
}

func (pc *PortageConverter) FindRuntimePackage(pTarget *anise_pkg.DefaultPackage) (*anise_pkg.DefaultPackage, error) {
	var ans *anise_pkg.DefaultPackage = nil

	runtimeOpts := solver.NewRuntimeSolverOpts()
	s := solver.NewRuntimeSolver(pc.Config, runtimeOpts)

	// Setup the forest guard for the solver
	s.SetForestGuard(pc.ForestGuard)

	treesIdx, err := pc.ForestGuard.SearchPackage(pTarget)
	if err != nil {
		return nil, err
	}

	if len(treesIdx) == 0 {
		treesIdx, err = pc.ForestGuard.SearchProvides(pTarget)
		if err != nil {
			return nil, err
		}
	}

	if len(treesIdx) == 0 {
		return nil, fmt.Errorf("no packages found for package %s", pTarget.HumanReadableString())
	}

	for _, tree := range treesIdx {
		// Load all availables packages

		versions, _ := tree.GetPackageVersions(pTarget.PackageName())

		for _, version := range versions {

			// Using render engine to read build.yaml
			defFile := filepath.Join(tree.TreePath, tree.BaseDir, version.Path)
			runtimePackage, err := anise_tree.ReadDefinitionFile(defFile)
			if err != nil {
				return nil, err
			}

			ans = runtimePackage
			break

		}

		if ans != nil {
			break
		}

	}

	return ans, nil
}
