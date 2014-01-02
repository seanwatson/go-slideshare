//TODO: comments/docs
//TODO: tests
//TODO: check api errors, make real errors
//TODO: add good params

package slideshare

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseUrl = "https://www.slideshare.net/api/2/"

func parseFromHttpResponse(resp *http.Response, container interface{}) error {
	//TODO: check the HEAD
	responseBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
	    return err
	}
	return xml.Unmarshal([]byte(responseBody), container)
}

type SlideShare struct {
	ApiKey       string // Developer key provided by SlideShare
	SharedSecret string // Developer secret provided by SlideShare
}

func (s *SlideShare) getUrl(method string, args map[string]string) string {
	values := url.Values{}
	for k, v := range args {
	    values.Set(k, v)
	}
	values.Set("api_key", s.ApiKey)
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	values.Set("ts", timestamp)
	hash := sha1.New()
	io.WriteString(hash, s.SharedSecret+timestamp)
	values.Set("hash", fmt.Sprintf("%x", hash.Sum(nil)))
	return baseUrl + method + "?" + values.Encode()
}

func (s *SlideShare) GetSlideshow(id int) (Slideshow, error) {
	args := map[string]string{
	    "slideshow_id": strconv.ItoA(id),
	}
	url := s.getUrl("get_slideshow", args)
	resp, err := http.Get(url)
	if err != nil {
	    return Slideshow{}, err
	}
	slideshow := Slideshow{}
	err = parseFromHttpResponse(resp, &slideshow)
	return slideshow, err
}

func (s *SlideShare) GetSlideshowsByTag(tag string) ([]Slideshow, error) {
	args := map[string]string{
	    "tag": tag,
	}
	url := s.getUrl("get_slideshows_by_tag", args)
	resp, err := http.Get(url)
	if err != nil {
	    return []Slideshow{}, err
	}
	slideshows := Slideshows{}
	err = parseFromHttpResponse(resp, &slideshows)
	return slideshows.Values, err
}

func (s *SlideShare) GetSlideshowsByGroup(groupName string) ([]Slideshow, error) {
	args := map[string]string{
	    "group_name": groupName,
	}
	url := s.getUrl("get_slideshows_by_group", args)
	resp, err := http.Get(url)
	if err != nil {
	    return []Slideshow{}, err
	}
	slideshows := Slideshows{}
	err = parseFromHttpResponse(resp, &slideshows)
	return slideshows.Values, err
}

func (s *SlideShare) GetSlideshowsByUser(usernameFor string) ([]Slideshow, error) {
	args := map[string]string{
	    "username_for": usernameFor,
	}
	url := s.getUrl("get_slideshows_by_tag", args)
	resp, err := http.Get(url)
	if err != nil {
	    return []Slideshow{}, err
	}
	slideshows := Slideshows{}
	err = parseFromHttpResponse(resp, &slideshows)
	return slideshows.Values, err
}

func (s *SlideShare) SearchSlideshows(q string) ([]Slideshow, error) {
	args := map[string]string{
	    "q": q,
	}
	url := s.getUrl("search_slideshows", args)
	resp, err := http.Get(url)
	if err != nil {
	    return []Slideshow{}, err
	}
	slideshows := Slideshows{}
	err = parseFromHttpResponse(resp, &slideshows)
	return slideshows.Values, err
}

/*
func (s *SlideShare) EditSlideshow(username, password, slideshowId string) error {
	args := map[string]string{
		"username":     username,
		"password":     password,
		"slideshow_id": slideshowId,
	}
	url := s.getUrl("edit_slideshow", args)
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("OMFG ERROR")
	}
	return nil //checkEditResponse(resp)
}

func (s *SlideShare) DeleteSlideshow(username, password, slideshowId string) error {
	args := map[string]string{
		"username":     username,
		"password":     password,
		"slideshow_id": slideshowId,
	}
	url := s.getUrl("delete_slideshow", args)
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("OMFG ERROR")
	}
	return nil //checkDeleteResponse(resp)
}

func (s *SlideShare) UploadSlideshow(username, password, slideshowTitle, uploadUrl string) (string, error) {
	args := map[string]string{
		"username":        username,
		"password":        password,
		"slideshow_title": slideshowTitle,
		"upload_url":      uploadUrl,
	}
	url := s.getUrl("upload_slideshow", args)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("OMFG ERROR")
	}
	return checkUploadResponse(resp)
}*/

func (s *SlideShare) GetUserGroups(usernameFor string) ([]Group, error) {
	args := map[string]string{
	    "username_for": usernameFor,
	}
	url := s.getUrl("get_user_groups", args)
	resp, err := http.Get(url)
	if err != nil {
	    return []Group{}, err
	}
	groups := Groups{}
	err = parseFromHttpResponse(resp, &groups)
	return groups.Values, err
}

func (s *SlideShare) GetUserFavorites(usernameFor string) ([]Favorite, error) {
	args := map[string]string{
	    "username_for": usernameFor,
	}
	url := s.getUrl("get_user_favorites", args)
	resp, err := http.Get(url)
	if err != nil {
	    return []Favorite{}, err
	}
	favorites := Favorites{}
	err = parseFromHttpResponse(resp, &favorites)
	return favorites.Values, err
}

func (s *SlideShare) GetUserContacts(usernameFor string, limit int, offset int) ([]Contact, error) {
	args := map[string]string{
	    "username_for": usernameFor,
	    "limit":        strconv.Itoa(limit),
	    "offset":       strconv.Itoa(offset),
	}
	url := s.getUrl("get_user_contacts", args)
	resp, err := http.Get(url)
	if err != nil {
	    return []Contact{}, err
	}
	contacts := Contacts{}
	err = parseFromHttpResponse(resp, &contacts)
	return contacts.Values, err
}

func (s *SlideShare) GetUserTags(username, password string) ([]Tag, error) {
	args := map[string]string{
	    "username": username,
	    "password": password,
	}
	url := s.getUrl("get_user_tags", args)
	resp, err := http.Get(url)
	if err != nil {
	    return []Tag{}, err
	}
	tags := Tags{}
	err = parseFromHttpResponse(resp, &tags)
	return tags.Values, err
}

func (s *SlideShare) AddFavorite(username string, password string, slideshowId int) error {
	args := map[string]string{
	    "username":     username,
	    "password":     password,
	    "slideshow_id": strconv.Itoa(slideshowId),
	}
	url := s.getUrl("add_favorite", args)
	resp, err := http.Get(url)
	if err != nil {
	    return err
	}
	addFavoriteResponse := AddFavoriteResponse{}
	err = parseFromHttpResponse(resp, &addFavoriteResponse)
	if err != nil {
	    return err
	}
	if addFavoriteResponse.SlideshowId == 0 {
	    return errors.New("Could not add favorite")
	}
	return nil
}

func (s *SlideShare) CheckFavorite(username string, password string, slideshowId int) (bool, error) {
	args := map[string]string{
	    "username":     username,
	    "password":     password,
	    "slideshow_id": strconv.Itoa(slideshowId),
	}
	url := s.getUrl("check_favorite", args)
	resp, err := http.Get(url)
	if err != nil {
	    return false, err
	}
	checkFavoriteResponse := CheckFavoriteResponse{}
	err = parseFromHttpResponse(resp, &checkFavoriteResponse)
	return checkFavoriteResponse.Favorited, err
}
