package model

import (
	"encoding/xml"
	"time"

	"github.com/twinj/uuid"
)

type News struct {
	ID          string   `json:"id"`
	TeamID      string   `json:"teamId"`
	OptaMatchID string   `json:"optaMatchId"`
	Title       string   `json:"title"`
	Type        []string `json:"type"`
	Teaser      string   `json:"teaser"`
	Content     string   `json:"content"`
	URL         string   `json:"url"`
	ImageURL    string   `json:"imageUrl"`
	GalleryUrls string   `json:"galleryUrls"`
	VideoURL    string   `json:"videoUrl"`
	Published   string   `json:"published"`
	ArticleID   string   `json:"articleID"`
}

func (n *News) CreateNewsStructFromGenericXMLNewsList(newsXMLGenerics NewsletterNewsItem) {
	n.ID = uuid.NewV4().String()
	n.TeamID = "t94"
	n.OptaMatchID = newsXMLGenerics.OptaMatchId
	n.Title = newsXMLGenerics.Title
	n.Teaser = newsXMLGenerics.TeaserText
	n.URL = newsXMLGenerics.ArticleURL
	n.Published = newsXMLGenerics.PublishDate
	n.ArticleID = newsXMLGenerics.NewsArticleID
}

func (n *News) CreateNewsStructFromDetailXMLNews(newsXMLDetail NewsArticle) {
	n.Type = append(n.Type, newsXMLDetail.Taxonomies)
	n.Content = newsXMLDetail.BodyText
	n.ImageURL = newsXMLDetail.ThumbnailImageURL
	n.GalleryUrls = newsXMLDetail.GalleryImageURLs
	n.VideoURL = newsXMLDetail.VideoURL
}

type NewListInformation struct {
	XMLName             xml.Name `xml:"NewListInformation"`
	ClubName            string   `xml:"ClubName"`
	ClubWebsiteURL      string   `xml:"ClubWebsiteURL"`
	NewsletterNewsItems struct {
		NewsletterNewsItem []NewsletterNewsItem `xml:"NewsletterNewsItem"`
	} `xml:"NewsletterNewsItems"`
}

type NewsletterNewsItem struct {
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
}

type NewsArticleInformation struct {
	XMLName        xml.Name    `xml:"NewsArticleInformation"`
	ClubName       string      `xml:"ClubName"`
	ClubWebsiteURL string      `xml:"ClubWebsiteURL"`
	NewsArticle    NewsArticle `xml:"NewsArticle"`
}

type NewsArticle struct {
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
	Status   string                   `json:"status"`
	Data     News                     `json:"data"`
	Metadata NewsResponseByIDMetadata `json:"metadata"`
}

type NewsResponseByIDMetadata struct {
	CreatedAt time.Time `json:"createdAt"`
}
