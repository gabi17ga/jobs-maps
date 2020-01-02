package main

import (
	"regexp"

	"github.com/davecgh/go-spew/spew"
)

func extract(content string) info {
	rule := regexp.MustCompile(`<p\sclass="profile-company-name">\s*(?P<company>.*)\s*</p>`)
	matches := rule.FindAllStringSubmatch(content, -1)
	spew.Dump(matches)
	i := info{}

	return i
}
