package components

// MissionProps contains configuration for the mission section.
type MissionProps struct {
	Heading     string
	Description string
}

// DefaultMissionProps returns standard mission content for the homepage.
func DefaultMissionProps() MissionProps {
	return MissionProps{
		Heading:     "Our Mission",
		Description: "The Fallen Outdoors is a 501(c)(3) nonprofit organization dedicated to honoring our fallen heroes by providing free outdoor adventures to veterans, active duty military, and Gold Star families. Through hunting, fishing, and outdoor experiences, we create a supportive community that promotes healing, camaraderie, and connection with nature.",
	}
}
