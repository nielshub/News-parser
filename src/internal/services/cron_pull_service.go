package services

import "incrowd/src/internal/ports"

type CronPullService struct {
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
