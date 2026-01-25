package about

// Leader represents a leadership team member.
type Leader struct {
	Name  string
	Title string
	Bio   string
	Photo string
}

// Document represents a transparency document.
type Document struct {
	Name        string
	Description string
	Href        string
}

// DefaultLeaders returns placeholder leadership team members.
func DefaultLeaders() []Leader {
	return []Leader{
		{
			Name:  "John Smith",
			Title: "Founder & Executive Director",
			Bio:   "Army veteran with 20 years of service. Founded TFO in 2015 after experiencing firsthand the healing power of outdoor adventures with fellow veterans.",
			Photo: "",
		},
		{
			Name:  "Sarah Johnson",
			Title: "Board Chair",
			Bio:   "Gold Star mother and passionate advocate for veteran support services. Brings 15 years of nonprofit leadership experience.",
			Photo: "",
		},
		{
			Name:  "Mike Davis",
			Title: "Operations Director",
			Bio:   "Marine Corps veteran and experienced outdoorsman. Coordinates all hunt logistics and volunteer management.",
			Photo: "",
		},
		{
			Name:  "Lisa Chen",
			Title: "Development Director",
			Bio:   "Experienced fundraiser with a heart for veteran causes. Manages donor relations and grant writing.",
			Photo: "",
		},
	}
}

// DefaultDocuments returns placeholder transparency documents.
func DefaultDocuments() []Document {
	return []Document{
		{
			Name:        "501(c)(3) Determination Letter",
			Description: "IRS determination letter confirming our tax-exempt status.",
			Href:        "#",
		},
		{
			Name:        "Form 990 (2024)",
			Description: "Annual return filed with the IRS for fiscal year 2024.",
			Href:        "#",
		},
		{
			Name:        "Form 990 (2023)",
			Description: "Annual return filed with the IRS for fiscal year 2023.",
			Href:        "#",
		},
		{
			Name:        "Annual Report (2024)",
			Description: "Comprehensive overview of our programs, impact, and financials.",
			Href:        "#",
		},
	}
}
