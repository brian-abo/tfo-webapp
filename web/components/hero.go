package components

// HeroProps contains configuration for the hero section.
type HeroProps struct {
	Headline    string
	Subheadline string
	CTAText     string
	CTAHref     string
}

// DefaultHeroProps returns standard hero content for the homepage.
func DefaultHeroProps() HeroProps {
	return HeroProps{
		Headline:    "Adventure Awaits",
		Subheadline: "Connecting veterans and Gold Star families with the healing power of the outdoors.",
		CTAText:     "Join a Hunt",
		CTAHref:     "/hunts",
	}
}
