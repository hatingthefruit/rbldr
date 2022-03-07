package util

import (
	"fmt"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func BuildResume(resume Resume, template Template) string {
	outStr := template.Base

	outStr += fmt.Sprintf(template.Contact, resume.Contact.Name, resume.Contact.Email, resume.Contact.Email, resume.Contact.Phone)

	for _, section := range strings.Split(template.Order, "\n") {
		var itemStr, sectionStr, listStr string
		if section == "Education" {
			for _, edu := range resume.Education {
				itemStr += fmt.Sprintf(template.SectionItem, edu.Organization, edu.GradString(), edu.CertName, edu.FormattedGPA())
			}
		} else if section == "Experience" {
			for _, exp := range resume.Experience {
				itemStr += fmt.Sprintf(template.SectionItem, exp.Employer, exp.MonthRange(), exp.Role, exp.Location)
				if len(exp.Responsibilities) > 0 {
					listStr = ""
					for _, resp := range exp.Responsibilities {
						listStr += fmt.Sprintf(template.Item, resp)
					}
					itemStr += fmt.Sprintf(template.ItemList, listStr)
				}
			}
		}
		sectionStr += fmt.Sprintf(template.SectionBase, section, itemStr)
		sectionStr = EscapeCharacters(sectionStr)
		outStr += sectionStr
	}

	for section, value := range resume.Other {
		//var itemStr,
		var listStr, sectionStr string
		for bold, item := range value {
			if len(item) > 0 {
				listStr += fmt.Sprintf(template.SkillItem, bold, item)
			}
		}
		//itemStr = fmt.Sprintf(template.ItemList, listStr)
		sectionStr += fmt.Sprintf(template.SectionBase, section, listStr)
		sectionStr = EscapeCharacters(sectionStr)
		outStr += sectionStr
	}

	outStr += template.End
	return outStr
}
