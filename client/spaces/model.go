package spaces

import (
	"github.com/mittwald/spacectl/client/lowlevel"
	"time"
)

type SpaceName struct {
	DNSName           string `json:"dnsName"`
	HumanReadableName string `json:"humanReadableName"`
}

type SoftwareRef struct {
	ID   string `json:"id"`
	HREF string `json:"href,omitempty"`
}

type VersionRef struct {
	Number string `json:"number"`
}

type TeamRef struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	DNSLabel string `json:"dnsLabel"`
}

type Stage struct {
	Links   lowlevel.LinkList `json:"_links"`
	Actions lowlevel.LinkList `json:"_actions"`

	Name              string      `json:"name"`
	Application       SoftwareRef `json:"application"`
	Cronjobs          []Cronjob   `json:"cronjobs"`
	Version           VersionRef  `json:"version"`
	VersionConstraint string      `json:"versionConstraint"`
	UserData          interface{} `json:"userData"`
	DNSNames          []string    `json:"dnsNames"`
	Running           bool        `json:"running"`
}

type StageDeclaration struct {
	Name              string      `json:"name"`
	Application       SoftwareRef `json:"application"`
	Cronjobs          []Cronjob   `json:"cronjobs"`
	VersionConstraint string      `json:"versionConstraint"`
	UserData          interface{} `json:"userData"`
}

type Space struct {
	ID        string            `json:"id"`
	Links     lowlevel.LinkList `json:"_links"`
	Name      SpaceName         `json:"name"`
	CreatedAt time.Time         `json:"createdAt"`
	Status    string            `json:"status"`
	DNSNames  []string          `json:"dnsNames"`
	Stages    []Stage           `json:"stages"`
	Team      TeamRef           `json:"team"`
	Running   bool              `json:"running"`
}

type SpaceDeclaration struct {
	Name   SpaceName          `json:"name"`
	Stages []StageDeclaration `json:"stages"`
}

func (s Space) StagesCount() int {
	return len(s.Stages)
}

func (s Space) StagesNames() []string {
	names := make([]string, len(s.Stages))
	for i := range s.Stages {
		names[i] = s.Stages[i].Name
	}
	return names
}
