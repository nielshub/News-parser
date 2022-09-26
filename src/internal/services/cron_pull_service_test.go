package services

import (
	"encoding/json"
	"errors"
	"incrowd/src/internal/model"
	"incrowd/src/mocks"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mocksCronPullService struct {
	pullNewsURL                        string
	pullArticleURL                     string
	nonRelationalSportNewsDBRepository *mocks.MockNonRelationalSportNewsDBRepository
}

func queryNewsToFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	bodyResp := model.NewListInformation{}
	bytesBodyResp, _ := json.Marshal(bodyResp)
	w.Write(bytesBodyResp)
}

func TestGetNewsFromFeed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// · Mocks · //
	result := model.NewListInformation{}
	queryError := errors.New("error sending req for news feed. Error:")
	// · Tests · //
	type want struct {
		result model.NewListInformation
		err    error
	}

	tests := []struct {
		name              string
		want              want
		newsServerHandler func(w http.ResponseWriter, r *http.Request)
	}{
		{
			name: "Should get news succesfully",
			want: want{
				result: result,
				err:    nil,
			},
			newsServerHandler: func(w http.ResponseWriter, r *http.Request) {
				queryNewsToFeed(w, r)
			},
		},
		{
			name: "Should return error - Failed to query DB",
			want: want{
				result: result,
				err:    queryError,
			},
			newsServerHandler: func(w http.ResponseWriter, r *http.Request) {
				queryNewsToFeed(w, r)
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			feedServer := httptest.NewServer(http.HandlerFunc(tt.newsServerHandler))
			os.Setenv("NEWSURL", feedServer.URL)
			defer feedServer.Close()

			m := mocksCronPullService{
				nonRelationalSportNewsDBRepository: mocks.NewMockNonRelationalSportNewsDBRepository(gomock.NewController(t)),
			}

			cronPullService := NewCronPullService(m.nonRelationalSportNewsDBRepository)
			result, err := cronPullService.GetNewsFromFeed()

			assert.Equal(t, tt.want.result, result)
			assert.Equal(t, tt.want.err, err)
		})

	}
}
