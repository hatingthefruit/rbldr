package util

import (
	"fmt"
	"strings"
)

func (edu Education) FormattedGPA() string {
	if edu.Gpa > 4 || edu.Gpa < 0 {
		return ""
	} else {
		return fmt.Sprintf("GPA: %.2f", edu.Gpa)
	}
}

func (edu Education) GradString() string {
	var gradString string
	if edu.Graduated {
		gradString = "Graduated: "
	} else {
		gradString = "Expected Graduation: "
	}
	gradString += edu.GradDate.Month + " " + edu.GradDate.Year
	return gradString
}

func (exp Experience) MonthRange() string {
	startDate := exp.StartDate.Month + " " + exp.StartDate.Year
	endDate := exp.EndDate.Month + " " + exp.EndDate.Year
	return startDate + "--" + endDate
}

func EscapeCharacters(input string) string {
	invalidChars := []string{"$", "&", "%"}
	for _, char := range invalidChars {
		input = strings.ReplaceAll(input, char, "\\"+char)
	}
	return input
}
