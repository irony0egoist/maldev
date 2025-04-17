package all

import "github.com/irony0egoist/maldev/src/system"

func GetHome() (string, error) {
	return system.GetHome()
}

func GetGroups() ([]string, error) {
	return system.GetGroups()
}

func Env() []string {
	return system.Env()
}
