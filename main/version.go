package main

import (
	"fmt"

	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/main/commands/base"
)

var cmdVersion = &base.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "Show current version of Core",
	Long: `Version prints the build information for Core executables.
	`,
	Run: executeVersion,
}

func executeVersion(cmd *base.Command, args []string) {
	printVersion()
}

func printVersion() {
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}
