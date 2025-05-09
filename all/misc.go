package all

import "github.com/irony0egoist/maldev/src/misc"

func RandomString(length int) string {
	return misc.RandomString(length)
}

func RandomInt(max int, min int) int {
	return misc.RandomInt(max, min)
}

func GeneratePattern(length int) string {
	return misc.GeneratePattern(length)
}

func GetPatternOffset(pattern_str string) int {
	return misc.GetPatternOffset(pattern_str)
}

func BytesToGb(num int) int {
	return misc.BytesToGb(num)
}

func GbToBytes(num int) int {
	return misc.GbToBytes(num)
}

func DateToEpoch(year, month, day, hour, minute, second int) int {
	return misc.DateToEpoch(year, month, day, hour, minute, second)
}

func EpochToDate(epoch int) string {
	return misc.EpochToDate(epoch)
}

func GetRandomAgent() string {
	return misc.GetRandomAgent()
}
