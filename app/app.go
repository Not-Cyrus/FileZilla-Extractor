package app

import (
	"os"

	"github.com/Not-Cyrus/FileZilla-Exporter/core"
	"github.com/Not-Cyrus/FileZilla-Exporter/core/data"
	"github.com/Not-Cyrus/FileZilla-Exporter/log"
	"github.com/Not-Cyrus/FileZilla-Exporter/utils"
	"github.com/urfave/cli/v2"
)

func Run() {
	clientApp := &cli.App{
		Name:    "FileZilla Exporter",
		Usage:   "Export server logins/key files",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "savetype", Aliases: []string{"st"}, Destination: &saveType, Value: "all", Usage: "switch the save type between all/managed/recent"},
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"vv"}, Destination: &verbose, Value: false, Usage: "verbose"},
		},
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {
			switch verbose {
			case true:
				log.InitLog("debug")
			default:
				log.InitLog("error")
			}

			log.HandleError(utils.MakeDir("results"))
			core.InitOS()

			managedServers, err := core.GetSiteManagers()
			log.HandleError(err)

			recentServers, err := core.GetRecentServers()
			log.HandleError(err)

			switch saveType {

			case "all":
				data.SaveBoth(managedServers, recentServers)

			case "managed":
				data.SaveManagers(managedServers)

			case "recent":
				data.SaveRecent(recentServers)

			}

			data.SaveManagers(managedServers)

			fmt.Printf("Exported %d managed servers and %d recently connected servers to the results folder.\n", len(managedServers.AllServers.Servers), len(recentServers.AllServers.Servers))
			return nil
		},
	}
	err := clientApp.Run(os.Args)
	if err != nil {
		log.Error(err)
	}
}

var (
	saveType  string
	verbose   bool
)
