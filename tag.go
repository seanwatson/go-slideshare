package slideshare

type Tag struct {
	Count uint32 `xml:"count,attr"`
	Name  string `xml:"tag"`
}

type Tags struct {
    Values []Tag `xml:"tag"`
}
