package components

// NavItem represents a navigation menu item.
type NavItem struct {
	Label string
	Href  string
}

// DefaultNavItems returns the standard navigation items for the site.
func DefaultNavItems() []NavItem {
	return []NavItem{
		{Label: "Home", Href: "/"},
		{Label: "About", Href: "/about"},
		{Label: "Gallery", Href: "/gallery"},
		{Label: "Contact", Href: "/contact"},
	}
}
