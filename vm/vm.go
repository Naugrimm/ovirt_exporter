package vm

type VMs struct {
	VMs []VM `xml:"vm"`
}

type VM struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name"`
	Host struct {
		ID string `xml:"id,attr"`
	} `xml:"host,omitempty"`
	Cluster struct {
		ID string `xml:"id,attr"`
	} `xml:"cluster,omitempty"`
	Status string `xml:"status>state"`
	CPU    struct {
		Topology struct {
			Cores   int `xml:"cores,attr"`
			Sockets int `xml:"sockets,attr"`
			Threads int `xml:"threads,attr"`
		} `xml:"topology"`
	} `xml:"cpu"`
}
