package common

import "strings"

func RemoveAllLineSpace(s string) string{
	return RemoveAllLine(RemoveAllSpace(s))
}
func RemoveAllLine(s string) string {
	return strings.Replace(s, "\n", "", -1)
}
func RemoveAllSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}
