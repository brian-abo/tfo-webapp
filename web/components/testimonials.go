package components

// Testimonial represents a veteran testimonial.
type Testimonial struct {
	Quote  string
	Name   string
	Detail string
}

// DefaultTestimonials returns placeholder testimonials for the homepage.
func DefaultTestimonials() []Testimonial {
	return []Testimonial{
		{
			Quote:  "The Fallen Outdoors gave me something I didn't know I was missing. Being out in nature with fellow veterans who understand—it's healing in a way nothing else has been.",
			Name:   "John M.",
			Detail: "U.S. Army Veteran",
		},
		{
			Quote:  "After losing my son, I didn't think I'd ever find peace again. This organization brought me into a community that truly cares. I'm forever grateful.",
			Name:   "Sarah T.",
			Detail: "Gold Star Mother",
		},
		{
			Quote:  "These hunts aren't just about the outdoors. They're about brotherhood, healing, and remembering why we served. TFO gets it.",
			Name:   "Mike R.",
			Detail: "U.S. Marine Corps Veteran",
		},
	}
}
