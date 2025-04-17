package all

import "github.com/irony0egoist/maldev/src/redteam"

func GetShellcodeFromUrl(sc_url string) ([]byte, error) {
	return redteam.GetShellcodeFromUrl(sc_url)
}

func GetShellcodeFromFile(file string) ([]byte, error) {
	return redteam.GetShellcodeFromFile(file)
}

func WriteShellcodeToFile(filename string, shellcode []byte) error {
	return redteam.WriteShellcodeToFile(filename, shellcode)
}
