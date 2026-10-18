/*
	Copyright © 2021-2026 Macaroni OS Linux
	See AUTHORS and LICENSE for the license details and contributors.
*/

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/macaroni-os/anise-portage-converter/pkg/converter"

	. "github.com/macaroni-os/anise/pkg/config"
	. "github.com/macaroni-os/anise/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	cliName = `Copyright © 2021-2026 Macaroni OS Linux

Portage/Overlay bridge for Macaroni OS Sambuca stack.`

	version = "0.17.0"
)

func initConfig() error {
	AniseCfg.Viper.SetEnvPrefix("ANISE")
	AniseCfg.Viper.AutomaticEnv() // read in environment variables that match

	// Create EnvKey Replacer for handle complex structure
	replacer := strings.NewReplacer(".", "__")
	AniseCfg.Viper.SetEnvKeyReplacer(replacer)
	AniseCfg.Viper.SetTypeByDefaultValue(true)

	err := AniseCfg.Viper.Unmarshal(&AniseCfg)
	if err != nil {
		return err
	}

	InitAurora()
	NewSpinner()

	return nil
}

func Execute() {

	var rootCmd = &cobra.Command{
		Use:     "anise-portage-converter --",
		Short:   cliName,
		Version: fmt.Sprintf("%s-g%s %s", version, converter.BuildCommit, converter.BuildTime),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := initConfig()
			if err != nil {
				fmt.Println("Error on setup config/logger: " + err.Error())
				os.Exit(1)
			}
		},
	}

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug verbosity.")

	rootCmd.AddCommand(
		newGenerateCommand(),
		newReposcanResolveCommand(),
		newReposcanMetadataCommand(),
		newReposcanGenerateCommand(),
		newSyncCommand(),
		newPortageCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
