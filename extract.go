package main

import (
	"regexp"
)

type info struct {
	CompanyName string
	Address     string
	Jobs        []job
}

type job struct {
	Name string
	Link string
	City string
}

func extract(content string) info {
	i := info{}

	rule := regexp.MustCompile(`<p\sclass="profile-company-name">\s*(?P<company>.*)\s*</p>`)
	matches := rule.FindAllStringSubmatch(content, -1)
	i.CompanyName = matches[0][1]

	rule = regexp.MustCompile(`<span class="contact-info\scontact-info--adress">\s*<i(?:.*)</i>\s*(.*)\s*</span>`)
	matches = rule.FindAllStringSubmatch(content, -1)
	i.Address = matches[0][1]

	rule = regexp.MustCompile(`<h2\sclass="jobitem-title"\sitemprop="title">\s*<a(?:.*)>(.*)</a></h2>`)
	nameMatches := rule.FindAllStringSubmatch(content, -1)

	rule = regexp.MustCompile(`<a\sclass="title\sdataLayerItemLink"\shref="(.*)">(?:.*)</a>\s*</h2>`)
	linkMatches := rule.FindAllStringSubmatch(content, -1)

	rule = regexp.MustCompile(`<a\shref="javascript:;"\sclass="jobitem-info-item\sjobitem-icon\sjobitem-cities-qtip"\stitle="(.*)"\s*itemprop="jobLocation">`)
	cityMatches := rule.FindAllStringSubmatch(content, -1)

	jobs := []job{}
	for index, j := range nameMatches {
		jb := job{
			Name: j[1],
			Link: linkMatches[index][1],
			City: cityMatches[index][1],
		}
		jobs = append(jobs, jb)
	}
	i.Jobs = jobs

	return i
}
