/*
Copyright © 2021-2026 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/macaroni-os/anise-portage-converter/pkg/converter"
	"github.com/macaroni-os/anise-portage-converter/pkg/specs"

	. "github.com/macaroni-os/anise/pkg/config"
	. "github.com/macaroni-os/anise/pkg/logger"

	"github.com/spf13/cobra"
)

func newReposcanResolveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reposcan-resolve",
		Aliases: []string{"rr"},
		Short:   "Resolve a package from reposcan tree.",
		PreRun: func(cmd *cobra.Command, args []string) {
			pkg, _ := cmd.Flags().GetString("pkg")
			if pkg == "" {
				fmt.Println("Missing --pkg argument")
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
			jsonOutput, _ := cmd.Flags().GetBool("json")
			backend, _ := cmd.Flags().GetString("backend")
			ignoreMissingDeps, _ := cmd.Flags().GetBool("ignore-missing-deps")
			continueWithError, _ := cmd.Flags().GetBool("continue-with-error")
			pkg, _ := cmd.Flags().GetString("pkg")
			allowEmptyKeywords, _ := cmd.Flags().GetBool("allow-empty-keywords")

			converter := converter.NewPortageConverter(to, backend)
			converter.Override = override
			converter.IgnoreMissingDeps = ignoreMissingDeps
			converter.TreePaths = treePath
			converter.DisableStage2 = stage2
			converter.DisableStage3 = stage3
			converter.DisableStage4 = !stage4
			converter.ContinueWithError = continueWithError

			if debug {
				AniseCfg.GetGeneral().Debug = debug
			}

			err := converter.LoadRules(rulesFile)
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

			err = converter.InitConverter(false)
			if err != nil {
				Error(fmt.Sprintf("Error on init converter: %s", err.Error()))
				os.Exit(1)
			}

			artefact := *specs.NewPortageConverterArtefact(pkg)

			// Check if it's present artefact from map
			art, err := converter.Specs.GetArtefactByPackage(pkg)
			if err == nil {
				DebugC(fmt.Sprintf("[%s] Using artefact from map. Uses disabled: %s, enabled: %s",
					pkg, art.Uses.Disabled, art.Uses.Enabled))
				// POST: use artefact from map.
				artefact = *art
			}

			opts := specs.PortageResolverOpts{
				EnableUseFlags:   artefact.Uses.Enabled,
				DisabledUseFlags: artefact.Uses.Disabled,
				Conditions:       artefact.Conditions,
			}

			solution, err := converter.Resolver.Resolve(pkg, &opts)
			if err != nil {
				Error(fmt.Sprintf("Error on resolve %s: %s", pkg, err.Error()))
				os.Exit(1)
			}

			if jsonOutput {
				fmt.Println(solution)
			} else {

				fmt.Println(fmt.Sprintf("Package: %s-%s",
					solution.Package.GetPackageNameWithSlot(),
					solution.Package.GetPVR()))

				fmt.Println(fmt.Sprintf("Description: %s",
					solution.Description))

				fmt.Println("Use flags:")
				for _, use := range solution.Package.UseFlags {
					fmt.Println("- " + use)
				}

				fmt.Println("Build deps:")
				for _, p := range solution.BuildDeps {
					fmt.Println("-" + p.GetPackageNameWithSlot())
				}

				fmt.Println("Runtime deps:")
				for _, p := range solution.RuntimeDeps {
					fmt.Println("-" + p.GetPackageNameWithSlot())
				}

				fmt.Println("Labels:")
				for k, v := range solution.Labels {
					fmt.Println(k + ": " + v)
				}

			}

		},
	}

	cmd.Flags().String("to", "", "Targer tree where bump new specs.")
	cmd.Flags().Bool("override", false, "Override existing specs if already present.")
	cmd.Flags().Bool("json", false, "JSON output.")
	cmd.Flags().Bool("continue-with-error", false, "Continue processing with errors (for example: no KEYWORDS defined).")
	cmd.Flags().StringP("pkg", "p", "", "Define the package to analyze.")
	cmd.PersistentFlags().Bool("with-portage-pkg", false, "Generate portage packages for every required package.")
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
	return cmd
}
