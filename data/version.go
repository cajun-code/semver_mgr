package data

import "fmt"

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Build int `json:"build"`
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Build)
}

func (v *Version) Increment(field string) {
	switch field {
	case "major":
		v.Major++
	case "minor":
		v.Minor++
	case "build":
		v.Build++
	}

}

func NewVersion() *Version {
	return &Version{
		Major: 0,
		Minor: 0,
		Build: 0,
	}
}
