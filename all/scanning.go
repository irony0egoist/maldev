package all

import (
	"github.com/irony0egoist/maldev/src/scanning"
	wp "github.com/likexian/whois-parser"
)

func CheckUrl(url_to_check string, timeout int) bool {
	return scanning.CheckUrl(url_to_check, timeout)
}

func GetHttpFromDomain(domain string, timeout int) (string, error) {
	return scanning.GetHttpFromDomain(domain, timeout)
}

func Hostscan(ip_range string, ping_timeout int) ([]string, error) {
	return scanning.Hostscan(ip_range, ping_timeout)
}

func CheckIfUp(ip_to_check string, timeout int) (bool, error) {
	return scanning.CheckIfUp(ip_to_check, timeout)
}

func PortscanAll(ip string) ([]int, error) {
	return scanning.PortscanAll(ip)
}

func PortscanCommon(ip string) ([]int, error) {
	return scanning.PortscanCommon(ip)
}

func CheckOpenTcpPort(ip string, port int) bool {
	return scanning.CheckOpenTcpPort(ip, port)
}

func CheckOpenUdpPort(ip string, port int) bool {
	return scanning.CheckOpenUdpPort(ip, port)
}

func GetCrt(dom string) ([]string, error) {
	return scanning.GetCrt(dom)
}

func GetHackerTarget(dom string) ([]string, error) {
	return scanning.GetHackerTarget(dom)
}

func GetAlienVault(dom string) ([]string, error) {
	return scanning.GetAlienVault(dom)
}

func GetAllSubdomains(dom string) ([]string, error) {
	return scanning.GetAllSubdomains(dom)
}

func Wappalyzer(url string) ([]string, error) {
	return scanning.Wappalyzer(url)
}

func WhoisDomain(dom string) (wp.WhoisInfo, error) {
	return scanning.WhoisDomain(dom)
}
