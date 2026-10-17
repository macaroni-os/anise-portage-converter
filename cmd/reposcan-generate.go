/*
Copyright © 2021-2024 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/macaroni-os/anise-portage-converter/pkg/reposcan"

	//. "github.com/macaroni-os/anise/pkg/config"
	. "github.com/macaroni-os/anise/pkg/logger"
	"github.com/spf13/cobra"
)

func newReposcanGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reposcan-generate [kit-path]",
		Aliases: []string{"rg"},
		Short:   "Generate the reposcan JSON file of a specific kit.",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("Invalid argument. Only one kit admitted.")
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			portageBinPath, _ := cmd.Flags().GetString("portage-binpath")
			eclassDirs, _ := cmd.Flags().GetStringArray("eclass-dir")
			output, _ := cmd.Flags().GetString("output")
			kit, _ := cmd.Flags().GetString("kit")
			branch, _ := cmd.Flags().GetString("branch")
			file, _ := cmd.Flags().GetString("file")
			concurrency, _ := cmd.Flags().GetInt("concurrency")

			if concurrency < 1 {
				concurrency = 1
			}
			if file == "" {
				file = fmt.Sprintf("%s-%s", kit, branch)
			}

			generator := reposcan.NewRepoScanGenerator(portageBinPath)
			generator.SetConcurrency(concurrency)

			kitPath := args[0]

			for _, ed := range eclassDirs {
				generator.AddEclassesPath(ed)
			}

			ans, err := generator.ProcessKit(kit, branch, kitPath)
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
				err = ans.WriteJsonFile(file)
			}

			if err != nil {
				Fatal(err.Error())
			}
			if output == "yaml" || output == "json" {
				fmt.Println(out)
			} else {
				fmt.Println(fmt.Sprintf("File %s written with %d packages and %d errors.",
					file, generator.Counter, generator.Errors))
			}
		},
	}

	cmd.Flags().StringArray("eclass-dir", []string{},
		"Set the eclass directories to use (The path is without eclass/ subdirectory).")
	cmd.Flags().String("portage-binpath", "/usr/lib/portage/python3.9",
		"Override default portage path and python version.")
	cmd.Flags().StringP("output", "o", "json", "Output mode: json|yaml|file")
	cmd.Flags().StringP("file", "f", "", "Write kit cache file.")
	cmd.Flags().String("branch", "", "Define optionally the branch of the kit used.")
	cmd.Flags().String("kit", "", "Define optionally the name of the kit used.")
	cmd.Flags().Int("concurrency", runtime.NumCPU(), "Override concurrency elaboration option.")

	return cmd
}
