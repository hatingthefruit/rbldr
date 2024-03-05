package util

import (
	"io"
	"text/template"
)

func FillTemplate(resume Resume, resumeTemplate string, outputWriter io.Writer) error {
	t := template.Must(template.New("resumeTemplate").Parse(resumeTemplate))

	return t.Execute(outputWriter, resume)
}
