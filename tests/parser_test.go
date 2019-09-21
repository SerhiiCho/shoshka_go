package tests

import (
	"testing"

	"github.com/SerhiiCho/shoshka_go/utils"
)

func TestGetLinksFromHTML_returns_links_from_html_string(t *testing.T) {
	fileText := utils.FileGetContents("files/example")
	result := utils.GetLinksFromHTML(fileText)

	if len(result) != 8 {
		t.Error("GetLinksFromHTML must return slice with 8 items")
	}
}

func TestGetAllInformation_returns_slice_with_PhotoReport_models(t *testing.T) {
	links := []string{
		utils.FileGetContents("files/link"),
	}

	title := "Title here"
	img := "https://image.jpg"
	url := "https://test.com/hello/"

	result := utils.GetAllInformation(links)

	if result[0].Title != title {
		t.Error("Returned Title from GetAllInformation func must be `" + title + "` but `" + result[0].Title + "` returned")
	}

	if result[0].Image != img {
		t.Error("Returned Image from GetAllInformation func must be `" + img + "` but `" + result[0].Image + "` returned")
	}

	if result[0].URL != url {
		t.Error("Returned URL from GetAllInformation func must be `" + url + "` but `" + result[0].URL + "` returned")
	}
}