package components

// Stat represents a single impact statistic.
type Stat struct {
	Value string
	Label string
}

// DefaultStats returns placeholder impact statistics for the homepage.
func DefaultStats() []Stat {
	return []Stat{
		{Value: "5,000+", Label: "Veterans Served"},
		{Value: "50", Label: "States Reached"},
		{Value: "500+", Label: "Hunts Completed"},
		{Value: "100%", Label: "Free to Veterans"},
	}
}
