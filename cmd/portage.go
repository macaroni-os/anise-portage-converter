/*
Copyright © 2021-2026 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package cmd

import (
	cmd_portage "github.com/macaroni-os/anise-portage-converter/cmd/portage"

	"github.com/spf13/cobra"
)

func newPortageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "portage",
		Aliases: []string{"p"},
		Short:   "Execute operations related to Portage.",
		Long: `
Execute operations to purge edb cache (portage-cache) or to update
Portage world file.
`,
	}

	cmd.AddCommand(
		cmd_portage.NewPortageCacheCommand(),
		cmd_portage.NewPortageShowCommand(),
		cmd_portage.NewPortageAddWorldCommand(),
	)

	return cmd
}
