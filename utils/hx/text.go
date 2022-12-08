package hx

import "strings"

func FilterUnPrintable(s string) string {
	r := ""
	for _, c := range s {
		if c >= 32 && c <= 126 {
			r += string(c)
		}
	}
	return r
}

func HandleTerminalEscape(s string) string {
	s = strings.ReplaceAll(s, " ", "&nbsp;")
	s = strings.ReplaceAll(s, "\n", "<br/>\n")
	s = strings.ReplaceAll(s, "\x1b[30m", "<i style=\"color:black;\">")
	s = strings.ReplaceAll(s, "\x1b[31m", "<i style=\"color:#e57373;\">")
	s = strings.ReplaceAll(s, "\x1b[32m", "<i style=\"color:#a5d6a7;\">")
	s = strings.ReplaceAll(s, "\x1b[33m", "<i style=\"color:#fdd835;\">")
	s = strings.ReplaceAll(s, "\x1b[34m", "<i style=\"color:#0277bd;\">")
	s = strings.ReplaceAll(s, "\x1b[35m", "<i style=\"color:magenta;\">")
	s = strings.ReplaceAll(s, "\x1b[36m", "<i style=\"color:cyan;\">")
	s = strings.ReplaceAll(s, "\x1b[37m", "<i style=\"color:white;\">")
	s = strings.ReplaceAll(s, "\x1b[0m", "</i>")
	s = strings.ReplaceAll(s, "├", "&boxvr;")
	s = strings.ReplaceAll(s, "└", "&boxur;")
	s = strings.ReplaceAll(s, "─", "&boxh;")
	s = strings.ReplaceAll(s, "│", "&boxv;")
	s = strings.ReplaceAll(s, "┌", "&boxdl;")
	s = strings.ReplaceAll(s, "←", "&larr;")
	return "<p>" + s + "</p>"
}
