package util

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func BuildResume(resume Resume, template string) string {
	rootPath := "templates/" + template + "/"
	files, err := ioutil.ReadDir(rootPath)
	CheckErr(err)
	fileMap := make(map[string]string)
	for _, eachFile := range files {

		if !eachFile.IsDir() {
			tempStr, err := ioutil.ReadFile(rootPath + eachFile.Name())
			fileMap[eachFile.Name()] = string(tempStr)
			CheckErr(err)
		}
	}
	outStr := fileMap["base"]

	outStr += fmt.Sprintf(fileMap["contact"], resume.Contact.Name, resume.Contact.Email, resume.Contact.Email, resume.Contact.Phone)

	for _, section := range strings.Split(fileMap["order"], "\n") {
		var itemStr, sectionStr, listStr string
		if section == "Education" {
			for _, edu := range resume.Education {
				itemStr += fmt.Sprintf(fileMap["sectionItem"], edu.Organization, edu.GradString(), edu.CertName, edu.FormattedGPA())
			}
		} else if section == "Experience" {
			for _, exp := range resume.Experience {
				itemStr += fmt.Sprintf(fileMap["sectionItem"], exp.Employer, exp.Location, exp.Role, exp.MonthRange())
				if len(exp.Responsibilities) > 0 {
					listStr = ""
					for _, resp := range exp.Responsibilities {
						listStr += fmt.Sprintf(fileMap["item"], resp)
					}
					itemStr += fmt.Sprintf(fileMap["itemList"], listStr)
				}
			}
		}
		sectionStr += fmt.Sprintf(fileMap["sectionBase"], section, itemStr)
		sectionStr = EscapeCharacters(sectionStr)
		outStr += sectionStr
	}

	for section, value := range resume.Other {
		//var itemStr,
		var listStr, sectionStr string
		for bold, item := range value {
			if len(item) > 0 {
				listStr += fmt.Sprintf(fileMap["skillItem"], bold, item)
			}
		}
		//itemStr = fmt.Sprintf(fileMap["itemList"], listStr)
		sectionStr += fmt.Sprintf(fileMap["sectionBase"], section, listStr)
		sectionStr = EscapeCharacters(sectionStr)
		outStr += sectionStr
	}

	outStr += fileMap["end"]
	return outStr
}
