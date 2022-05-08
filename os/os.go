package os

import "strings"

type OS int

const (
	None OS = iota
	CentOS
)

func (o OS) GenerateOSInstance() Commader {
	switch o {
	case CentOS:
		return &centOs{}
	}

	return nil
}

func (o OS) toString() string {
	switch o {
	case CentOS:
		return "centos"
	}
	return ""
}

func ValueOf(osType string) OS {
	switch osType {
	case CentOS.toString():
		return CentOS
	}
	return None
}

func AvailableOS() string {
	return strings.Join([]string{CentOS.toString()}, ",")
}
