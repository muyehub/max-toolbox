package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
)

// getOSVersion 获取系统版本
func getOSVersion() string {
	var version string

	switch runtime.GOOS {
	case "windows":
		// 使用 "ver" 命令获取 Windows 版本
		cmd := exec.Command("cmd", "/C", "ver")
		output, err := cmd.Output()
		if err != nil {
			return "Unable to get version: " + err.Error()
		}
		version = string(output)

	case "linux":
		// 使用 "lsb_release -a" 命令获取 Linux 版本
		cmd := exec.Command("lsb_release", "-a")
		output, err := cmd.Output()
		if err != nil {
			return "Unable to get version: " + err.Error()
		}
		version = string(output)

	case "darwin": // macOS
		// 使用 "sw_vers" 命令获取 macOS 版本
		cmd := exec.Command("sw_vers")
		output, err := cmd.Output()
		if err != nil {
			return "Unable to get version: " + err.Error()
		}
		version = string(output)

	default:
		version = "Unsupported OS"
	}

	return version
}

// getInternalIP 获取内网地址
func getInternalIP() ([]string, error) {
	var ipList []string

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// 遍历网络接口
	for _, iface := range interfaces {
		// 忽略关闭的接口和回环接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口的地址
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		// 遍历接口的地址
		for _, addr := range addrs {
			// 检查地址类型是否为 IP 地址
			if ipnet, ok := addr.(*net.IPNet); ok {
				ipList = append(ipList, ipnet.IP.String())
			}
		}
	}

	return ipList, nil
}

// getExternalIP 获取外网地址
func getExternalIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get external IP, status: %s", resp.Status)
	}

	// 读取响应体
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

// filterIPv4 过滤出 IPv4 地址
func filterIPv4(ips []string) []string {
	var ipv4List []string
	ipv4Regex := regexp.MustCompile(`^\d{1,3}(\.\d{1,3}){3}$`) // 匹配 IPv4 格式的正则表达式

	for _, ip := range ips {
		if ipv4Regex.MatchString(ip) {
			ipv4List = append(ipv4List, ip)
		}
	}

	return ipv4List
}

// GeoIPResponse 用于解析 IP 地址返回的地理位置信息
type GeoIPResponse struct {
	IP           string `json:"ip"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	City         string `json:"city"`
	Location     string `json:"loc"`
	Organization string `json:"org"`
}

func getGeoLocation() (*GeoIPResponse, error) {
	// 请求 ipinfo.io 的 API
	resp, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get geolocation: %s", resp.Status)
	}

	var geoInfo GeoIPResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoInfo); err != nil {
		return nil, err
	}

	return &geoInfo, nil
}
