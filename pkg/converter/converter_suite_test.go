/*
Copyright © 2021-2023 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package converter_test

import (
	"strings"
	"testing"

	. "github.com/macaroni-os/anise/pkg/config"
	. "github.com/macaroni-os/anise/pkg/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func initConfig() error {
	AniseCfg.Viper.SetEnvPrefix("LUET")
	AniseCfg.Viper.AutomaticEnv() // read in environment variables that match

	// Create EnvKey Replacer for handle complex structure
	replacer := strings.NewReplacer(".", "__")
	AniseCfg.Viper.SetEnvKeyReplacer(replacer)
	AniseCfg.Viper.SetTypeByDefaultValue(true)

	err := AniseCfg.Viper.Unmarshal(&AniseCfg)
	if err != nil {
		return err
	}

	AniseCfg.GetGeneral().Debug = true

	InitAurora()
	NewSpinner()

	return nil
}

func TestSolver(t *testing.T) {
	RegisterFailHandler(Fail)
	initConfig()
	RunSpecs(t, "Converter Suite")
}
