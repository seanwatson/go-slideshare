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
	slideshows, err := ss.GetSlideshowsByTag("slideshare", 2, 0)
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
	testServer := getTestServer(getSlideshowsByGroupResponse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
}

func TestGetSlideshowsByUser(t *testing.T) {
	testServer := getTestServer(getSlideshowsByUserResponse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
}

func TestSearchSlideshows(t *testing.T) {
	testServer := getTestServer(searchSlideshowsResponse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
}

func TestGetUserGroups(t *testing.T) {
	testServer := getTestServer(getUserGroupsResponse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
	groups, err := ss.GetUserGroups("test")
	if err != nil {
		t.Fatalf("TestGetUserGroups: %v", err)
	}
	if len(groups) != len(getUserGroupsExpectedResult) {
		t.Fatalf("TestGetUserGroups:\n\tGot     : %v\n\tExpected: %v",
			groups, getUserGroupsExpectedResult)
	}
	for i, v := range getUserGroupsExpectedResult {
		if groups[i] != v {
			t.Fatalf("TestGetUserGroups:\n\tGot     : %v\n\tExpected: %v",
				groups, getUserGroupsExpectedResult)
		}
	}
}

func TestGetUserFavorites(t *testing.T) {
	testServer := getTestServer(getUserFavoritesResponse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
	favorites, err := ss.GetUserFavorites("test")
	if err != nil {
		t.Fatalf("TestGetUserFavorites: %v", err)
	}
	if len(favorites) != len(getUserFavoritesExpectedResult) {
		t.Fatalf("TestGetUserFavorites:\n\tGot     : %v\n\tExpected: %v",
			favorites, getUserFavoritesExpectedResult)
	}
	for i, v := range getUserFavoritesExpectedResult {
		if favorites[i] != v {
			t.Fatalf("TestGetUserFavorites:\n\tGot     : %v\n\tExpected: %v",
				favorites, getUserFavoritesExpectedResult)
		}
	}
}

func TestGetUserContacts(t *testing.T) {
	testServer := getTestServer(getUserContactsResponse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
	contacts, err := ss.GetUserContacts("test", 2, 0)
	if err != nil {
		t.Fatalf("TestGetUserContacts: %v", err)
	}
	if len(contacts) != len(getUserContactsExpectedResult) {
		t.Fatalf("TestGetUserContacts:\n\tGot     : %v\n\tExpected: %v",
			contacts, getUserContactsExpectedResult)
	}
	for i, v := range getUserContactsExpectedResult {
		if contacts[i] != v {
			t.Fatalf("TestGetUserContacts:\n\tGot     : %v\n\tExpected: %v",
				contacts, getUserContactsExpectedResult)
		}
	}
}

func TestGetUserTags(t *testing.T) {
	testServer := getTestServer(getUserTagsResponse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
	tags, err := ss.GetUserTags("test", "test")
	if err != nil {
		t.Fatalf("TestGetUserTags: %v", err)
	}
	if len(tags) != len(getUserTagsExpectedResult) {
		t.Fatalf("TestGetUserTags:\n\tGot     : %v\n\tExpected: %v",
			tags, getUserTagsExpectedResult)
	}
	for i, v := range getUserTagsExpectedResult {
		if tags[i] != v {
			t.Fatalf("TestGetUserTags:\n\tGot     : %v\n\tExpected: %v",
				tags, getUserTagsExpectedResult)
		}
	}
}

func TestAddFavorite(t *testing.T) {
	testServer := getTestServer(addFavoriteResponse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
	err := ss.AddFavorite("test", "test", 123)
	if err != nil {
		t.Fatalf("TestAddFavorite: %v", err)
	}
}

func TestCheckFavoriteTrue(t *testing.T) {
	testServer := getTestServer(checkFavoriteResponseTrue)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
	fav, err := ss.CheckFavorite("test", "test", 123)
	if err != nil {
		t.Fatalf("TestCheckFavorite: %v", err)
	}
	if !fav {
		t.Fatalf("TestCheckFavorite: Got: false, Expected: true")
	}
}

func TestCheckFavoriteFalse(t *testing.T) {
	testServer := getTestServer(checkFavoriteResponseFalse)
	defer testServer.Close()

	ss := SlideShare{apiKey, sharedSecret}
	fav, err := ss.CheckFavorite("test", "test", 123)
	if err != nil {
		t.Fatalf("TestCheckFavorite: %v", err)
	}
	if fav {
		t.Fatalf("TestCheckFavorite: Got: true, Expected: false")
	}
}
