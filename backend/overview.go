package backend

import (
	"fmt"
)

// SystemInfo returns a system info
func (a *App) SystemInfo() string {
	return getOSVersion()
}

func (a *App) IpInfo() string {
	internalIp, err := getInternalIP()
	if err != nil {
		return err.Error()
	}

	// 过滤出 IPv4 地址
	ipv4IPs := filterIPv4(internalIp)

	externalIp, err := getExternalIP()
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("内网 ip：%s 外网 ip: %s", ipv4IPs, externalIp)
}

func (a *App) LocationInfo() string {
	data, err := getGeoLocation()
	if err != nil {
		return err.Error()
	}

	return data.City
}
