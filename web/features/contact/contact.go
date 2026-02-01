package contact

// RegionalLeader represents a regional director/leader.
type RegionalLeader struct {
	ID     string
	Name   string
	Region string
	Email  string
}

// DefaultRegionalLeaders returns placeholder regional leaders.
func DefaultRegionalLeaders() []RegionalLeader {
	return []RegionalLeader{
		{
			ID:     "west-coast",
			Name:   "David Lee",
			Region: "West Coast",
			Email:  "westcoast@thefallenoutdoors.org",
		},
		{
			ID:     "midwest",
			Name:   "James Wilson",
			Region: "Midwest",
			Email:  "midwest@thefallenoutdoors.org",
		},
		{
			ID:     "east-coast",
			Name:   "Tom Anderson",
			Region: "East Coast",
			Email:  "eastcoast@thefallenoutdoors.org",
		},
		{
			ID:     "southern",
			Name:   "Maria Garcia",
			Region: "Southern",
			Email:  "southern@thefallenoutdoors.org",
		},
	}
}

// LeadersByID returns a map of leaders keyed by their region ID.
func LeadersByID() map[string]RegionalLeader {
	leaders := DefaultRegionalLeaders()
	result := make(map[string]RegionalLeader, len(leaders))
	for _, l := range leaders {
		result[l.ID] = l
	}
	return result
}
