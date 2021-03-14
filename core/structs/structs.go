package structs

import "encoding/xml"

type (
	Servers struct {
		Servers []Server `xml:"Server"`
	}

	Server struct {
		Host    string `xml:"Host"`
		Port    string `xml:"Port"`
		User    string `xml:"User"`
		Pass    string `xml:"Pass,omitempty"`
		Keyfile string `xml:"Keyfile,omitempty"`
	}

	SiteManagerList struct {
		XMLName    xml.Name `xml:"FileZilla3"`
		AllServers Servers  `xml:"Servers"`
	}

	RecentServerList struct {
		XMLName    xml.Name `xml:"FileZilla3"`
		AllServers Servers  `xml:"RecentServers"`
	}
)
