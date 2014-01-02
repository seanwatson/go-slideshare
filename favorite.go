package slideshare

type Favorite struct {
	SlideshowId uint32 `xml:"slideshow_id"`
	TagText     string `xml:"tag_text"`
}

type Favorites struct {
	Values []Favorite `xml:"favorite"`
}

type AddFavoriteResponse struct {
	SlideshowId uint32 `xml:"slideshowid"`
}

type CheckFavoriteResponse struct {
	SlideshowId uint32 `xml:"slideshowid"`
	User        string `xml:"user"`
	Favorited   bool   `xml:"favorited"`
}
