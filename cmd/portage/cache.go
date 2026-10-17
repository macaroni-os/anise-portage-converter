/*
Copyright © 2021-2024 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package cmd_portage

import (
	"github.com/macaroni-os/anise-portage-converter/pkg/portage"

	. "github.com/macaroni-os/anise/pkg/config"
	. "github.com/macaroni-os/anise/pkg/logger"

	"github.com/spf13/cobra"
)

func NewPortageCacheCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cache",
		Aliases: []string{"pc", "c"},
		Short:   "Purge Portage cache and optional packages from world file.",
		Long: `This command permits to cleanup Portage edb cache and remove packages from world file.

Only cleanup edb cache:
$> anise-portage-converter portage cache --purge

Cleanup edb cache and remove a package from world file:
$> anise-portage-converter portage c --purge --pkg cat/foo

Remove a package from the world file:
$> anise-portage-converter portage c --pkg cat/foo

Cleanup and rebuild edb cache:
$> anise-portage-converter portage c --rebuild


NOTE: This command needs root permissions.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			purge, _ := cmd.Flags().GetBool("purge")
			rebuild, _ := cmd.Flags().GetBool("rebuild")
			pkgs2remove, _ := cmd.Flags().GetStringArray("pkg")

			jobDone := false

			if debug {
				AniseCfg.GetGeneral().Debug = debug
			}

			if rebuild && !purge {
				purge = true
			}

			if purge {
				err := portage.CleanEdbCache("")
				if err != nil {
					Fatal(err)
				}
				jobDone = true
			}

			if len(pkgs2remove) > 0 {
				// Read the existing world file.
				packages, err := portage.GetWorldPackages("")
				if err != nil {
					Fatal(err)
				}

				packages, err = portage.GetWorldPackagesCleaned(
					packages, pkgs2remove,
				)
				if err != nil {
					Fatal(err)
				}

				err = portage.WriteWorldPackages("", packages)
				if err != nil {
					Fatal(err)
				}

				jobDone = true
			}

			if rebuild {
				err := portage.RebuildEdbCache()
				if err != nil {
					Fatal(err)
				}
				jobDone = true
			}

			if jobDone {
				Info(":panda: All done!")
			} else {
				Warning(":parachute: Nothing done.")
			}
		},
	}

	cmd.Flags().Bool("purge", false, "Purge Portage edb cache.")
	cmd.Flags().Bool("rebuild", false, "Rebuild Portage edb cache.")
	cmd.Flags().StringArrayP("pkg", "p", []string{},
		"Define the package to analyze.")

	return cmd
}
