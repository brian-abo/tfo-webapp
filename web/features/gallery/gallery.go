package gallery

// GalleryImage represents an image in the gallery.
type GalleryImage struct {
	Src     string
	Alt     string
	Caption string
}

// DefaultGalleryImages returns placeholder images for the gallery.
// Uses picsum.photos for placeholder images.
func DefaultGalleryImages() []GalleryImage {
	return []GalleryImage{
		{Src: "https://picsum.photos/seed/tfo1/800/600", Alt: "Hunt outing", Caption: "Veterans on a fall hunt in Montana"},
		{Src: "https://picsum.photos/seed/tfo2/800/600", Alt: "Fishing trip", Caption: "Fishing trip at Lake Tahoe"},
		{Src: "https://picsum.photos/seed/tfo3/800/600", Alt: "Group photo", Caption: "TFO team after a successful hunt"},
		{Src: "https://picsum.photos/seed/tfo4/800/600", Alt: "Outdoor adventure", Caption: "Hiking in the Rockies"},
		{Src: "https://picsum.photos/seed/tfo5/800/600", Alt: "Campfire", Caption: "Evening campfire stories"},
		{Src: "https://picsum.photos/seed/tfo6/800/600", Alt: "Wildlife", Caption: "Wildlife spotted on the trail"},
		{Src: "https://picsum.photos/seed/tfo7/800/600", Alt: "Award ceremony", Caption: "Annual volunteer recognition"},
		{Src: "https://picsum.photos/seed/tfo8/800/600", Alt: "Training", Caption: "Safety training session"},
		{Src: "https://picsum.photos/seed/tfo9/800/600", Alt: "Family event", Caption: "Gold Star family day"},
		{Src: "https://picsum.photos/seed/tfo10/800/600", Alt: "Sunrise hunt", Caption: "Early morning in the blind"},
		{Src: "https://picsum.photos/seed/tfo11/800/600", Alt: "Gear prep", Caption: "Getting ready for the day"},
		{Src: "https://picsum.photos/seed/tfo12/800/600", Alt: "Victory", Caption: "A successful day outdoors"},
	}
}
