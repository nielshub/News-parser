package model

import (
	"encoding/xml"
	"time"
)

type News struct {
	ID          string    `json:"id"`
	TeamID      string    `json:"teamId"`
	OptaMatchID string    `json:"optaMatchId"`
	Title       string    `json:"title"`
	Type        []string  `json:"type"`
	Teaser      string    `json:"teaser"`
	Content     string    `json:"content"`
	URL         string    `json:"url"`
	ImageURL    string    `json:"imageUrl"`
	GalleryUrls string    `json:"galleryUrls"`
	VideoURL    string    `json:"videoUrl"`
	Published   time.Time `json:"published"`
}

type NewListInformation struct {
	XMLName             xml.Name `xml:"NewListInformation"`
	ClubName            string   `xml:"ClubName"`
	ClubWebsiteURL      string   `xml:"ClubWebsiteURL"`
	NewsletterNewsItems struct {
		NewsletterNewsItem []struct {
			ArticleURL        string `xml:"ArticleURL"`
			NewsArticleID     string `xml:"NewsArticleID"`
			PublishDate       string `xml:"PublishDate"`
			Taxonomies        string `xml:"Taxonomies"`
			TeaserText        string `xml:"TeaserText"`
			ThumbnailImageURL string `xml:"ThumbnailImageURL"`
			Title             string `xml:"Title"`
			OptaMatchId       string `xml:"OptaMatchId"`
			LastUpdateDate    string `xml:"LastUpdateDate"`
			IsPublished       string `xml:"IsPublished"`
		} `xml:"NewsletterNewsItem"`
	} `xml:"NewsletterNewsItems"`
}

type NewsArticleInformation struct {
	XMLName        xml.Name `xml:"NewsArticleInformation"`
	ClubName       string   `xml:"ClubName"`
	ClubWebsiteURL string   `xml:"ClubWebsiteURL"`
	NewsArticle    struct {
		ArticleURL        string `xml:"ArticleURL"`
		NewsArticleID     string `xml:"NewsArticleID"`
		PublishDate       string `xml:"PublishDate"`
		Taxonomies        string `xml:"Taxonomies"`
		TeaserText        string `xml:"TeaserText"`
		Subtitle          string `xml:"Subtitle"`
		ThumbnailImageURL string `xml:"ThumbnailImageURL"`
		Title             string `xml:"Title"`
		BodyText          string `xml:"BodyText"`
		GalleryImageURLs  string `xml:"GalleryImageURLs"`
		VideoURL          string `xml:"VideoURL"`
		OptaMatchId       string `xml:"OptaMatchId"`
		LastUpdateDate    string `xml:"LastUpdateDate"`
		IsPublished       string `xml:"IsPublished"`
	} `xml:"NewsArticle"`
}

type NewsResponse struct {
	Status   string               `json:"status"`
	Data     []News               `json:"data"`
	Metadata NewsResponseMetadata `json:"metadata"`
}

type NewsResponseMetadata struct {
	CreatedAt  time.Time `json:"createdAt"`
	TotalItems int       `json:"totalItems"`
	Sort       string    `json:"sort"`
}

type NewsByIDResponse struct {
	Status   string               `json:"status"`
	Data     News                 `json:"data"`
	Metadata NewsResponseMetadata `json:"metadata"`
}

type NewsResponseByIDMetadata struct {
	CreatedAt time.Time `json:"createdAt"`
}
