package os

import "strings"

type OSKind int

const (
	None OSKind = iota
	CentOS
)

func (o OSKind) GenerateOSInstance() OS {
	switch o {
	case CentOS:
		return &centOs{}
	}

	return nil
}

func (o OSKind) toString() string {
	switch o {
	case CentOS:
		return "centos"
	}
	return ""
}

func ValueOf(osType string) OSKind {
	switch osType {
	case CentOS.toString():
		return CentOS
	}
	return None
}

func AvailableOS() string {
	return strings.Join([]string{CentOS.toString()}, ",")
}
