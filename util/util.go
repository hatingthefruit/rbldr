package util

import "github.com/hatingthefruit/rbldr/models"

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func BuildResume(resume models.Resume, template string) string {
	return "Built a resume"
}
