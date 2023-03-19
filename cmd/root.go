package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"glowing-potato/cmd/http"
	"os"
	"strings"
)

type AppInfo struct {
	AppName        string
	AppVersion     string
	AppCommit      string
	BuildGoVersion string
	BuildArch      string
	BuildDate      string
}

var (
	app *AppInfo

	rootCmd = &cobra.Command{
		Use:   "glowing-potato",
		Short: "glowing-potato",
		Long:  "glowing-potato",
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print version info",
		Long:  "Print version information of glowing-potato",
		Run: func(command *cobra.Command, args []string) {
			infoStr := strings.Builder{}
			fmt.Println(infoStr.String())
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(http.ServeHTTP())
}

func Execute(appInfo *AppInfo) {
	app = appInfo
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
