package services

import (
	"incrowd/src/mocks"
)

type mocksCronPullService struct {
	pullNewsURL                        string
	pullArticleURL                     string
	nonRelationalSportNewsDBRepository *mocks.MockNonRelationalSportNewsDBRepository
}
