/*
Copyright © 2021-2024 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package cmd_portage

import (
	"fmt"
	"os"

	"github.com/macaroni-os/anise-portage-converter/pkg/portage"

	. "github.com/macaroni-os/anise/pkg/config"
	. "github.com/macaroni-os/anise/pkg/logger"

	"github.com/spf13/cobra"
)

func NewPortageAddWorldCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-world [package1] ... [packageN]",
		Aliases: []string{"aw"},
		Short:   "Add one or more packages to world.",
		Long: `Add one or more packages to Portage world file.

Add a package with category and name:
$> anise-portage-converter portage add-world cat/foo

Add a package with slot (when != 0):
$> anise-portage-converter portage aw cat/foo:2

Add a package with repository:
$> anise-portage-converter portage aw cat/foo::core-kit

Add a package with slot and repository:
$> anise-portage-converter portage aw cat/foo:2::core-kit

NOTE: The command automatically avoid to add duplicate.`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("At least one package is needed.")
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")

			if debug {
				AniseCfg.GetGeneral().Debug = debug
			}

			// Read the existing world file.
			packages, err := portage.GetWorldPackages("")
			if err != nil {
				Fatal(err)
			}

			packages, err = portage.CleanupWorldPackages(
				append(packages, args...),
			)
			if err != nil {
				Fatal(err)
			}

			err = portage.WriteWorldPackages("", packages)
			if err != nil {
				Fatal(err)
			}

			Info(":panda: All done!")
		},
	}

	return cmd
}
