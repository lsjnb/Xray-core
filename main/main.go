package main

import (
	"flag"
	"fmt"
	"os"

	authclient "github.com/Paper-Dragon/py-auth/client/go"
	"github.com/xtls/xray-core/main/commands/base"
	_ "github.com/xtls/xray-core/main/distro/all"
)

func main() {
	os.Args = getArgsV4Compatible()

	if err := checkAuthorization(); err != nil {
		fmt.Fprintf(os.Stderr, "exit status 99\n")
		os.Exit(99)
	}

	base.RootCommand.Long = "Core is a platform for building proxies."
	base.RootCommand.Commands = append(
		[]*base.Command{
			cmdRun,
			cmdVersion,
		},
		base.RootCommand.Commands...,
	)
	base.Execute()
}

func getArgsV4Compatible() []string {
	if len(os.Args) == 1 {
		return []string{os.Args[0], "run"}
	}
	if os.Args[1][0] != '-' {
		return os.Args
	}
	version := false
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.BoolVar(&version, "version", false, "")
	// parse silently, no usage, no error output
	fs.Usage = func() {}
	fs.SetOutput(&null{})
	err := fs.Parse(os.Args[1:])
	if err == flag.ErrHelp {
		// fmt.Println("DEPRECATED: -h, WILL BE REMOVED IN V5.")
		// fmt.Println("PLEASE USE: core help")
		// fmt.Println()
		return []string{os.Args[0], "help"}
	}
	if version {
		// fmt.Println("DEPRECATED: -version, WILL BE REMOVED IN V5.")
		// fmt.Println("PLEASE USE: core version")
		// fmt.Println()
		return []string{os.Args[0], "version"}
	}
	// fmt.Println("COMPATIBLE MODE, DEPRECATED.")
	// fmt.Println("PLEASE USE: core run [arguments] INSTEAD.")
	// fmt.Println()
	return append([]string{os.Args[0], "run"}, os.Args[1:]...)
}

type null struct{}

func (n *null) Write(p []byte) (int, error) {
	return len(p), nil
}

func checkAuthorization() error {
	serverURL := "https://auth.certauth.cn"
	softwareName := "xray-core"
	clientSecret := "aB3cD5eF7gH9iJ1kL3mN5oP7qR9sT1uV3wX5yZ7aB9cD1eF3gH5iJ7kL9mN1oP3qR5sT7uV9wX1yZ3"

	client, err := authclient.NewAuthClient(authclient.AuthClientConfig{
		ServerURL:         serverURL,
		SoftwareName:      softwareName,
		ClientSecret:      clientSecret,
		CacheValidityDays: 7,
		CheckIntervalDays: 2,
		Debug:             false,
	})
	if err != nil {
		return err
	}

	return client.RequireAuthorization(false)
}
