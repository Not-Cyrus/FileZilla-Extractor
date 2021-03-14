package core

import (
	"encoding/xml"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/Not-Cyrus/FileZilla-Exporter/core/structs"
	"github.com/Not-Cyrus/FileZilla-Exporter/utils"
)

func GetRecentServers() (servers structs.RecentServerList, err error) {
	var xmlData []byte

	xmlData, err = utils.OpenFile(fileZillaPath + recentServerFile)
	if err != nil {
		return
	}

	err = xml.Unmarshal(xmlData, &servers)

	return
}

func GetSiteManagers() (servers structs.SiteManagerList, err error) {
	var xmlData []byte

	xmlData, err = utils.OpenFile(fileZillaPath + siteManagerFile)
	if err != nil {
		return
	}

	err = xml.Unmarshal(xmlData, &servers)

	return
}

func InitOS() {
	// we have this if statement for when I (if I add) macOS support. (Properly)
	switch runtime.GOOS {
	case "windows":
		fileZillaPath = os.Getenv("APPDATA") + fileZillaPath
	default:
		fmt.Print("Enter the FileZilla path (Do not add the \"FileZilla part\"): ")
		fmt.Scan(&fileZillaPath)

		fileZillaPath = strings.TrimSuffix(fileZillaPath, "/FileZilla") // I know there will be that one person who does it anyways
	}
}

const (
	siteManagerFile  = "sitemanager.xml"
	recentServerFile = "recentservers.xml"
)

type (
	TotalServers struct {
		Recent   structs.RecentServerList
		Managers structs.SiteManagerList
	}
)

var (
	fileZillaPath = "/FileZilla/"
)
