package poundctl

import (
	"encoding/xml"
	"os/exec"
)

type ListenerStatus string

const (
	ListenerStatus_Active   ListenerStatus = "active"
	ListenerStatus_Disabled ListenerStatus = "DISABLED"
)

type ListenerProtocol string

const (
	ListenerProtocol_HTTP  ListenerProtocol = "http"
	ListenerProtocol_HTTPS ListenerProtocol = "HTTPS"
)

type AliveStatus string

const (
	AliveStatus_Alive AliveStatus = "yes"
	AliveStatus_Dead  AliveStatus = "DEAD"
)

type PoundStatus struct {
	XMLName xml.Name   `xml:"pound"`
	Queue   PoundQueue `xml:"queue"`

	Listeners []PoundListener `xml:"listener"`
}

type PoundQueue struct {
	Size int `xml:"size,attr"`
}

type PoundListener struct {
	ID       int              `xml:"index,attr"`
	Protocol ListenerProtocol `xml:"protocol,attr"`
	Address  string           `xml:"address,attr"`
	Status   ListenerStatus   `xml:"status,attr"`

	Services []PoundService `xml:"service"`
}

type PoundService struct {
	ID     int            `xml:"index,attr"`
	Name   string         `xml:"name,attr"`
	Status ListenerStatus `xml:"status,attr"`

	Backends []PoundBackend `xml:"backend"`
}

type PoundBackend struct {
	ID       int            `xml:"index,attr"`
	Address  string         `xml:"address,attr"`
	Avg      float32        `xml:"avg,attr"`
	Priority int            `xml:"priority,attr"`
	Alive    AliveStatus    `xml:"alive,attr"`
	Status   ListenerStatus `xml:"status,attr"`
}

func GetStatus(socketFile string) (*PoundStatus, error) {
	x, err := exec.Command("poundctl", "-c", socketFile, "-X").Output()
	if err != nil {
		return nil, err
	}

	return ParseStatusXml(x)
}

func ParseStatusXml(resultXml []byte) (*PoundStatus, error) {
	v := PoundStatus{}

	err := xml.Unmarshal(resultXml, &v)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
