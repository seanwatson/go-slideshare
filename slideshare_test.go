package slideshare

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	apiKey       = "abc123"
	sharedSecret = "123abc"
)

func getTestServer(response string) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, response)
	}
	testServer := httptest.NewServer(http.HandlerFunc(handler))
	baseUrl = testServer.URL
	return testServer
}

func TestGetUrl(t *testing.T) {
}

func TestGetSlideshow(t *testing.T) {
}

func TestGetSlideshowsByTag(t *testing.T) {
	testServer := getTestServer(getSlideshowsByTagResponse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
	slideshows, err := ss.GetSlideshowsByTag("slideshare", 10, 0)
	if err != nil {
		t.Fatalf("TestGetSlideshowsByTag: %v", err)
	}
	if len(slideshows) != len(getSlideshowsByTagExpectedResult) {
		t.Fatalf("TestGetSlideshowsByTag:\n\tGot     : %v\n\tExpected: %v",
			slideshows, getSlideshowsByTagExpectedResult)
	}
	for i, v := range getSlideshowsByTagExpectedResult {
		if slideshows[i] != v {
			t.Fatalf("TestGetSlideshowsByTag:\n\tGot     : %v\n\tExpected: %v",
				slideshows, getSlideshowsByTagExpectedResult)
		}
	}
}

func TestGetSlideshowsByGroup(t *testing.T) {
}

func TestGetSlideshowsByUser(t *testing.T) {
}

func TestSearchSlideshows(t *testing.T) {
}

func TestGetUserGroups(t *testing.T) {
}

func TestGetUserFavorites(t *testing.T) {
}

func TestGetUserContacts(t *testing.T) {
}

func TestGetUserTags(t *testing.T) {
}

func TestAddFavorite(t *testing.T) {
}

func TestCheckFavorite(t *testing.T) {
}
