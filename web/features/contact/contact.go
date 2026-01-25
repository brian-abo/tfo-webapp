package contact

// RegionalLeader represents a regional director/leader.
type RegionalLeader struct {
	Name   string
	Region string
	Email  string
}

// DefaultRegionalLeaders returns placeholder regional leaders.
func DefaultRegionalLeaders() []RegionalLeader {
	return []RegionalLeader{
		{
			Name:   "Tom Anderson",
			Region: "Northeast",
			Email:  "northeast@thefallenoutdoors.org",
		},
		{
			Name:   "Maria Garcia",
			Region: "Southeast",
			Email:  "southeast@thefallenoutdoors.org",
		},
		{
			Name:   "James Wilson",
			Region: "Midwest",
			Email:  "midwest@thefallenoutdoors.org",
		},
		{
			Name:   "Sarah Brown",
			Region: "Southwest",
			Email:  "southwest@thefallenoutdoors.org",
		},
		{
			Name:   "David Lee",
			Region: "West",
			Email:  "west@thefallenoutdoors.org",
		},
		{
			Name:   "Jennifer Martinez",
			Region: "Pacific Northwest",
			Email:  "pnw@thefallenoutdoors.org",
		},
	}
}
