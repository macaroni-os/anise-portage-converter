/*
	Copyright © 2021 Funtoo Macaroni OS Linux
	See AUTHORS and LICENSE for the license details and contributors.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Luet-lab/luet-portage-converter/pkg/converter"

	. "github.com/mudler/luet/pkg/config"
	. "github.com/mudler/luet/pkg/logger"

	"github.com/spf13/cobra"
)

func newGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate luet specs.",
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
			to, _ := cmd.Flags().GetString("to")
			rulesFile, _ := cmd.Flags().GetString("rules")
			override, _ := cmd.Flags().GetBool("override")
			stage2, _ := cmd.Flags().GetBool("disable-stage2")
			stage3, _ := cmd.Flags().GetBool("disable-stage3")
			stage4, _ := cmd.Flags().GetBool("enable-stage4")
			debug, _ := cmd.Flags().GetBool("debug")
			backend, _ := cmd.Flags().GetString("backend")
			ignoreMissingDeps, _ := cmd.Flags().GetBool("ignore-missing-deps")
			continueWithError, _ := cmd.Flags().GetBool("continue-with-error")
			pkgs, _ := cmd.Flags().GetStringArray("pkg")
			withPortagePkgs, _ := cmd.Flags().GetBool("with-portage-pkg")
			disableConflicts, _ := cmd.Flags().GetBool("disable-conflicts")
			layer4Rdepends, _ := cmd.Flags().GetBool("layer4rdepends")

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

			if debug {
				LuetCfg.GetGeneral().Debug = debug
			}

			if len(pkgs) > 0 {
				converter.SetFilteredPackages(pkgs)
			}

			if len(treePath) == 0 {
				DebugC(GetAurora().Bold("ATTENTION! No trees defined."))
			}

			err := converter.LoadRules(rulesFile)
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

	return cmd
}
