/*
Copyright © 2021-2023 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/macaroni-os/anise-portage-converter/pkg/converter"
	"github.com/macaroni-os/anise/pkg/v2/render"

	. "github.com/macaroni-os/anise/pkg/config"
	. "github.com/macaroni-os/anise/pkg/logger"

	"github.com/spf13/cobra"
)

func newGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate anise specs.",
		PreRun: func(cmd *cobra.Command, args []string) {
			to, _ := cmd.Flags().GetString("to")
			if to == "" {
				fmt.Println("Missing --to argument")
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			treePath, _ := cmd.Flags().GetStringArray("tree")
			reposcanSources, _ := cmd.Flags().GetStringArray("reposcan-files")
			disableUseFlags, _ := cmd.Flags().GetStringArray("disable-use-flag")
			templatesDir, _ := cmd.Flags().GetStringArray("templates-dir")
			values, _ := cmd.Flags().GetStringArray("values")
			to, _ := cmd.Flags().GetString("to")
			rulesFile, _ := cmd.Flags().GetString("rules")
			override, _ := cmd.Flags().GetBool("override")
			stage2, _ := cmd.Flags().GetBool("disable-stage2")
			stage3, _ := cmd.Flags().GetBool("disable-stage3")
			stage4, _ := cmd.Flags().GetBool("enable-stage4")
			debug, _ := cmd.Flags().GetBool("debug")
			backend, _ := cmd.Flags().GetString("backend")
			ignoreMissingDeps, _ := cmd.Flags().GetBool("ignore-missing-deps")
			ignoreWrongPackages, _ := cmd.Flags().GetBool("ignore-wrong-packages")
			continueWithError, _ := cmd.Flags().GetBool("continue-with-error")
			checkUpdate4Deps, _ := cmd.Flags().GetBool("check-update4deps")
			pkgs, _ := cmd.Flags().GetStringArray("pkg")
			withPortagePkgs, _ := cmd.Flags().GetBool("with-portage-pkg")
			disableConflicts, _ := cmd.Flags().GetBool("disable-conflicts")
			layer4Rdepends, _ := cmd.Flags().GetBool("layer4rdepends")
			skipRdepsGen, _ := cmd.Flags().GetBool("skip-rdeps-generation")
			allowEmptyKeywords, _ := cmd.Flags().GetBool("allow-empty-keywords")

			converter := converter.NewPortageConverter(to, backend)
			converter.Override = override
			converter.IgnoreMissingDeps = ignoreMissingDeps
			converter.TreePaths = treePath
			converter.WithPortagePkgs = withPortagePkgs
			converter.DisableStage2 = stage2
			converter.DisableStage3 = stage3
			converter.DisableStage4 = !stage4
			converter.DisableConflicts = disableConflicts
			converter.UsingLayerForRuntime = layer4Rdepends
			converter.ContinueWithError = continueWithError
			converter.IgnoreWrongPackages = ignoreWrongPackages
			converter.CheckUpdate4Deps = checkUpdate4Deps
			converter.SkipRDepsGeneration = skipRdepsGen

			rEngine := render.NewRenderEngine(AniseCfg)
			err := rEngine.LoadTemplates(templatesDir)
			if err != nil {
				Fatal(fmt.Errorf(
					"error on load templates: %s", err.Error()))
			}

			err = rEngine.LoadDefaultValues(values)
			if err != nil {
				Fatal(fmt.Errorf(
					"error on load default values: %s", err.Error()))
			}
			converter.RenderEngine = rEngine

			if debug {
				AniseCfg.GetGeneral().Debug = debug
			}

			if len(pkgs) > 0 {
				converter.SetFilteredPackages(pkgs)
			}

			if len(treePath) == 0 {
				DebugC(GetAurora().Bold("ATTENTION! No trees defined."))
			}

			err = converter.LoadRules(rulesFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = converter.LoadTrees(treePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if len(reposcanSources) > 0 {
				for _, source := range reposcanSources {
					converter.Specs.AddReposcanSource(source)
				}
			}

			if len(disableUseFlags) > 0 {
				converter.Specs.ReposcanDisabledUseFlags = append(converter.Specs.ReposcanDisabledUseFlags, disableUseFlags...)
			}

			if allowEmptyKeywords {
				converter.Specs.ReposcanAllowEmpyKeywords = allowEmptyKeywords
			}

			err = converter.Generate()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		},
	}

	cmd.Flags().String("to", "", "Targer tree where bump new specs.")
	cmd.Flags().Bool("override", false, "Override existing specs if already present.")
	cmd.Flags().StringArrayP("pkg", "p", []string{},
		"Define the list of the packages to generate instead of the full list defined in rules file.")
	cmd.Flags().Bool("with-portage-pkg", false, "Generate portage packages for every required package.")
	cmd.Flags().Bool("disable-conflicts", false, "Disable elaboration of runtime and buildtime conflicts.")
	cmd.Flags().Bool("layer4rdepends", false, "Check layer for runtime deps and skip generation.")
	cmd.Flags().Bool("continue-with-error", false, "Continue processing with errors (for example: no KEYWORDS defined).")
	cmd.Flags().Bool("ignore-wrong-packages", false, "Continue processing when a package is not resolved.")
	cmd.Flags().Bool("check-update4deps", false, "Verify if there are update for package dependencies too.")
	cmd.Flags().Bool("skip-rdeps-generation", false,
		"Skip the generation of the runtime dependencies specs.")
	cmd.Flags().Bool("allow-empty-keywords", false, "Override spec option to allow empty KEYWORDS.")

	cmd.Flags().StringArrayP("tree", "t", []string{}, "Path of the tree to use.")
	cmd.Flags().String("rules", "", "Rules file.")
	cmd.Flags().String("backend", "reposcan", "Select backend resolver: qdepends|reposcan.")
	cmd.Flags().Bool("disable-stage2", false, "Disable stage2 phase.")
	cmd.Flags().Bool("disable-stage3", false, "Disable stage3 phase.")
	cmd.Flags().Bool("enable-stage4", false, "Enable experimental stage4 phase.")
	cmd.Flags().StringArray("reposcan-files", []string{}, "Append additional reposcan files. Only for reposcan backend.")
	cmd.Flags().StringArray("disable-use-flag", []string{}, "Append additional use flags to disable.")
	cmd.Flags().Bool("ignore-missing-deps", false, "Ignore missing deps on resolver.")
	cmd.Flags().StringArray("templates-dir", []string{}, "Set templates directories.")
	cmd.Flags().StringArray("values", []string{}, "Set render values.")
	return cmd
}
