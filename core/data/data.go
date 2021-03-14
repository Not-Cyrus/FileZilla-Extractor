package data

import (
	"fmt"

	"github.com/Not-Cyrus/FileZilla-Exporter/core/decrypt"
	"github.com/Not-Cyrus/FileZilla-Exporter/core/structs"
	"github.com/Not-Cyrus/FileZilla-Exporter/log"
	"github.com/Not-Cyrus/FileZilla-Exporter/utils"
)

func HandleManagedServer(s structs.SiteManagerList) savedServers {
	var (
		err     error
		servers savedServers
	)

	servers.Servers = make(map[int][]structs.Server)

	for index, server := range s.AllServers.Servers {
		server.Pass, err = decrypt.DecodePassword(server.Pass)
		log.HandleError(err)

		if len(server.Keyfile) != 0 {
			keyFile := fmt.Sprintf("/results/%s.pem", server.Host)

			err = utils.CopyFile(server.Keyfile, keyFile)
			if err == nil {
				server.Keyfile = fmt.Sprintf("%s\\%s.pem", utils.CurrentPath, server.Host)
			}

			log.HandleError(err)
		}

		servers.Servers[index] = append(servers.Servers[index], server)
	}

	return servers
}

func HandleRecentServers(s structs.RecentServerList) savedServers {
	var (
		err     error
		servers savedServers
	)

	servers.Servers = make(map[int][]structs.Server)

	for index, server := range s.AllServers.Servers {
		server.Pass, err = decrypt.DecodePassword(server.Pass)
		log.HandleError(err)

		servers.Servers[index] = append(servers.Servers[index], server)
	}

	return servers
}

func SaveBoth(s structs.SiteManagerList, s2 structs.RecentServerList) {
	SaveManagers(s)
	SaveRecent(s2)
}

func SaveManagers(s structs.SiteManagerList) {
	managedServers := HandleManagedServer(s)
	err := utils.WriteJSON("Site Manager Servers.json", managedServers)
	log.HandleError(err)

	fmt.Printf("Exported %d managed servers\n", len(managedServers.Servers))
}

func SaveRecent(s structs.RecentServerList) {
	recentServers := HandleRecentServers(s)
	err := utils.WriteJSON("Recent Servers.json", recentServers)
	log.HandleError(err)

	fmt.Printf("Exported %d recently connected servers\n", len(recentServers.Servers))
}

type (
	savedServers struct {
		Servers map[int][]structs.Server
	}
)
