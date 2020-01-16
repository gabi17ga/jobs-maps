package main

import (
	"regexp"
	"strconv"
	"strings"
)

type company struct {
	ID      int64
	Name    string
	Address string
	Jobs    []job `bigquery:"-"`
}

type job struct {
	ID        int64
	CompanyID int64
	Name      string
	Link      string
	City      string
}

// extract extrage inf din pagina
func extract(companyID int64, content string) (*company, error) {
	c := company{
		ID: companyID,
	}

	rule := regexp.MustCompile(`<p\sclass="profile-company-name">\s*(?P<company>.*)\s*</p>`)
	matches := rule.FindAllStringSubmatch(content, -1)

	if len(matches) > 0 {
		c.Name = matches[0][1]

		rule = regexp.MustCompile(`<span class="contact-info\scontact-info--adress">\s*<i(?:.*)</i>\s*(.*)\s*</span>`)
		matches = rule.FindAllStringSubmatch(content, -1)
		if len(matches) > 0 {
			c.Address = matches[0][1]
		}

		rule = regexp.MustCompile(`<h2\sclass="jobitem-title"\sitemprop="title">\s*<a(?:.*)>(.*)</a></h2>`)
		nameMatches := rule.FindAllStringSubmatch(content, -1)

		rule = regexp.MustCompile(`<a\sclass="title\sdataLayerItemLink"\shref="(.*)">(?:.*)</a>\s*</h2>`)
		linkMatches := rule.FindAllStringSubmatch(content, -1)

		rule = regexp.MustCompile(`<a\shref="javascript:;"\sclass="jobitem-info-item\sjobitem-icon\sjobitem-cities-qtip"\stitle="(.*)"\s*itemprop="jobLocation">`)
		cityMatches := rule.FindAllStringSubmatch(content, -1)

		jobs := []job{}

		for index, j := range nameMatches {
			link := linkMatches[index][1]
			list := strings.Split(link, "/")
			jID := list[len(list)-1]

			jobID, err := strconv.ParseInt(jID, 10, 64)
			if err != nil {
				return nil, err
			}

			jb := job{
				ID:        jobID,
				CompanyID: companyID,
				Name:      j[1],
				Link:      link,
				City:      cityMatches[index][1],
			}

			jobs = append(jobs, jb)
		}

		c.Jobs = jobs
	}

	return &c, nil
}
