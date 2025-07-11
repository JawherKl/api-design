package utils

import (
	"strings"
)

func ExtractValue(xmlStr, tag string) string {
	start := strings.Index(xmlStr, "<"+tag+">")
	end := strings.Index(xmlStr, "</"+tag+">")
	if start == -1 || end == -1 {
		return ""
	}
	return xmlStr[start+len(tag)+2 : end]
}
