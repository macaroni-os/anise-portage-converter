/*
Copyright © 2021-2024 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/macaroni-os/anise-portage-converter/pkg/reposcan"

	//. "github.com/macaroni-os/anise/pkg/config"
	. "github.com/macaroni-os/anise/pkg/logger"
	"github.com/spf13/cobra"
)

func newReposcanMetadataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reposcan-metadata [ebuild]",
		Aliases: []string{"rm"},
		Short:   "Parse an ebuild and get metadata in reposcan mode.",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("Invalid argument. Only one ebuild admitted.")
				os.Exit(1)
			}

			ebuildFile := args[0]
			if filepath.Ext(ebuildFile) != ".ebuild" {
				fmt.Println("File hasn't .ebuild extension:", filepath.Ext(ebuildFile))
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			portageBinPath, _ := cmd.Flags().GetString("portage-binpath")
			eclassDirs, _ := cmd.Flags().GetStringArray("eclass-dir")
			output, _ := cmd.Flags().GetString("output")
			kit, _ := cmd.Flags().GetString("kit")
			branch, _ := cmd.Flags().GetString("branch")

			generator := reposcan.NewRepoScanGenerator(portageBinPath)

			ebuildFile := args[0]

			for _, ed := range eclassDirs {
				generator.AddEclassesPath(ed)
			}

			// Retrieve category and package name
			basedir := filepath.Dir(ebuildFile)
			fragDir := strings.Split(basedir, "/")
			if len(fragDir) < 3 {
				fmt.Println("The supplied path doesn't respect Portage tree!")
				os.Exit(1)
			}

			ans, err := generator.ParseAtom(ebuildFile,
				fragDir[len(fragDir)-2],
				fragDir[len(fragDir)-1],
				kit,    // kit
				branch, // branch
			)
			if err != nil {
				Fatal(err.Error())
			}

			var out string

			switch output {
			case "yaml":
				out, err = ans.Yaml()
			case "json":
				out, err = ans.Json()
			default:
				out = fmt.Sprintf("%v", ans)
			}

			if err != nil {
				Fatal(err.Error())
			}
			fmt.Println(out)
		},
	}

	cmd.Flags().StringArray("eclass-dir", []string{},
		"Set the eclass directories to use (The path is without eclass/ subdirectory).")
	cmd.Flags().String("portage-binpath", "/usr/lib/portage/python3.9",
		"Override default portage path and python version.")
	cmd.Flags().StringP("output", "o", "json", "Output mode: yaml|json|terminal")
	cmd.Flags().String("branch", "", "Define optionally the branch of the kit used.")
	cmd.Flags().String("kit", "", "Define optionally the name of the kit used.")

	return cmd
}
