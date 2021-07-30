package models

type Date struct {
	Day   string `json:"day"`
	Month string `json:"month`
	Year  string `json:"year"`
}

type Experience struct {
	Role             string   `json:"role"`
	Employer         string   `json:"employer"`
	Location         string   `json:"location"`
	StartDate        Date     `json:"start_date"`
	EndDate          Date     `json:"end_date"`
	responsibilities []string `json:"responsibilities"`
}

type Education struct {
	Organization string  `json:"organization"`
	OrgShort     string  `json:"org_short"`
	CertName     string  `json:"cert_name"`
	Graduated    bool    `json:"graduated"`
	GradDate     Date    `json:"grad_date"`
	Gpa          float64 `json:"gpa`
}

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Cert struct {
	Company string   `json:"company"`
	Certs   []string `json:"certs"`
}

type ResumeDefinition struct {
	Contact    string   `json:"contact"`
	Profile    string   `json:"profile"`
	Education  []string `json:"education"`
	Experience []string `json:"experience"`
}

type Resume struct {
	Contact    Contact
	Profile    []string
	Education  []Education
	Experience []Experience
}
