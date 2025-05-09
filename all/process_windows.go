package all

import "github.com/irony0egoist/maldev/src/process"

func GetProcesses() ([]process.WindowsProcess, error) {
	return process.GetProcesses()
}

func FindPidByName(name string) ([]int, error) {
	return process.FindPidByName(name)
}

func FindNameByPid(pid int) (string, error) {
	return process.FindNameByPid(pid)
}
