package services

import (
	"encoding/xml"
	"errors"
	"incrowd/src/internal/model"
	"incrowd/src/internal/ports"
	"net/http"
	"os"
)

type CronPullService struct {
	pullNewsURL                     string
	pullArticleURL                  string
	relationalSportNewsDBRepository ports.NonRelationalSportNewsDBRepository
}

// 			"id": "created uuid4",
// 			"teamId": "t94", by default
// 			"optaMatchId": OptaMatchId XML ,GENERICS
// 			"title": Title XML, GENERICS
// 			"type": [
// 				"Brentford B Team"
// 			], Taxonomies XML DETAIL
// 			"teaser": "Yehoing joined Brentford recently from his boyho Dnipro.",  TEASER TEXT XML GENERICS
// 			"content": "<p>The 18-year-old arrived in West fter his 2021/22 season in his" BodyText XML DETAIL
// 			"url": "https://www.brentforliuk/", ArticleURL XML GENERICS
// 			"imageUrl": "https://www.bre62-9eea-084803035aa3/Medium/yehor-red.jpg", ThumbnailImageURL XML DETAIL
// 			"galleryUrls": null, GalleryImageURLs XML DETAIL
// 			"videoUrl": null, VideoURL XML DETAIL
// 			"published": "2022-07-21T15:48:00.000Z" PublishDate XML GENERICS

func NewCronPullService(relationalSportNewsDBRepository ports.NonRelationalSportNewsDBRepository) *CronPullService {
	return &CronPullService{
		pullNewsURL:                     os.Getenv("NEWSURL"),
		pullArticleURL:                  os.Getenv("ARTICLEURL"),
		relationalSportNewsDBRepository: relationalSportNewsDBRepository,
	}
}

func (cps *CronPullService) CronPullNewsRoutine() {
	//s := gocron.NewScheduler(time.UTC)

}

func (cps *CronPullService) GetNewsFromFeed() (model.NewListInformation, error) {
	client := &http.Client{}
	newsListXML := model.NewListInformation{}
	req, err := http.NewRequest("GET", cps.pullNewsURL, nil)
	if err != nil {
		return newsListXML, errors.New("error creating req for news feed. URL: " + cps.pullNewsURL + " .Error: " + err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return newsListXML, errors.New("error sending req for news feed. Error: " + err.Error())
	}

	err = xml.NewDecoder(resp.Body).Decode(&newsListXML)
	if err != nil {
		return newsListXML, errors.New("error decoding news feed response. Error: " + err.Error())
	}

	return newsListXML, nil
}

func (cps *CronPullService) CreateNewsArrayFromXMLList(newsListInXML model.NewListInformation) []model.News {
	var newsArray []model.News
	newsArrayInXML := newsListInXML.NewsletterNewsItems.NewsletterNewsItem

	for _, newsInXML := range newsArrayInXML {
		news := model.News{}
		news.CreateNewsStructFromGenericXMLNewsList(newsInXML)
		newsArray = append(newsArray, news)
	}

	return newsArray
}

func (cps *CronPullService) GetDetailInformationForEachNews(news []model.News) error {
	for i := range news {
		client := &http.Client{}
		newsDetailXML := model.NewsArticleInformation{}
		articleURL := cps.pullArticleURL + news[i].ArticleID
		req, err := http.NewRequest("GET", articleURL, nil)
		if err != nil {
			return errors.New("error creating req for news feed. URL: " + articleURL + " .Error: " + err.Error())
		}
		resp, err := client.Do(req)
		if err != nil {
			return errors.New("error sending req for news feed. Error: " + err.Error())
		}

		err = xml.NewDecoder(resp.Body).Decode(&newsDetailXML)
		if err != nil {
			return errors.New("error decoding news feed response. Error: " + err.Error())
		}

		news[i].CreateNewsStructFromDetailXMLNews(newsDetailXML.NewsArticle)
	}

	return nil
}
