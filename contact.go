package slideshare

type Contact struct {
	Username      string `xml:"Username"`
	NumSlideshows uint32 `xml:"NumSlideshows"`
	NumComments   uint32 `xml:"NumComments"`
}

type Contacts struct {
	Values []Contact `xml:"Contact"`
}
