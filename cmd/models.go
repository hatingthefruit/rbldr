package cmd


struct Date {
	day string
	month string
	year string
}


struct Experience {
	role string
	employer string
	location string
	start_date Date
	end_date Date
	responsibilities []string
}

struct Education {
	organization string
	org_short string
	cert_name string
	graduated bool
	grad_date Date
	gpa float
}

struct Contact {
	name string
	phone string
	email string
}