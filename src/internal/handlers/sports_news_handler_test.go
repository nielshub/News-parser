package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"incrowd/src/internal/model"
	"incrowd/src/internal/ports"
	"incrowd/src/internal/services"
	"incrowd/src/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockSportNewsHandler struct {
	router           *gin.RouterGroup
	sportNewsService ports.SportNewsService
}

type mocksSportNewsService struct {
	nonRelationalSportNewsDBRepository *mocks.MockNonRelationalSportNewsDBRepository
}

func TestGetNews(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// · Mocks · //
	news := []model.News{}
	response := model.NewsResponse{
		Status: "success",
		Data:   news,
		Metadata: model.NewsResponseMetadata{
			CreatedAt:  time.Now(),
			TotalItems: len(news),
			Sort:       "",
		},
	}
	jsonResponse, _ := json.Marshal(response)
	// · Tests · //
	type want struct {
		code     int
		response string
		err      error
	}

	tests := []struct {
		name  string
		user  model.News
		url   string
		want  want
		mocks func(mSNS mocksSportNewsService)
	}{
		{
			name: "Should get news succesfully",
			url:  "/v1/teams/t94/news",
			want: want{
				code:     http.StatusOK,
				response: string(jsonResponse),
				err:      nil,
			},
			mocks: func(mSNS mocksSportNewsService) {
				mSNS.nonRelationalSportNewsDBRepository.EXPECT().GetNews(gomock.Any()).Return(news, nil)
			},
		},
		{
			name: "Should return error - Failed to query DB",
			url:  "/v1/teams/t94/news",
			want: want{
				code: http.StatusInternalServerError,
				response: `{
					"message": "Error getting news from DB"
				}`,
				err: errors.New("Error getting news from DB"),
			},
			mocks: func(mSNS mocksSportNewsService) {
				mSNS.nonRelationalSportNewsDBRepository.EXPECT().GetNews(gomock.Any()).Return(news, errors.New("Error getting news from DB"))
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			mSNS := mocksSportNewsService{
				nonRelationalSportNewsDBRepository: mocks.NewMockNonRelationalSportNewsDBRepository(gomock.NewController(t)),
			}
			w := httptest.NewRecorder()
			r := gin.Default()
			app := r.Group("/")

			mSNH := mockSportNewsHandler{
				router:           app,
				sportNewsService: services.NewSportNewsService(mSNS.nonRelationalSportNewsDBRepository),
			}

			tt.mocks(mSNS)
			NewSportNewsHandler(mSNH.router, mSNH.sportNewsService)

			req, err := http.NewRequest("GET", tt.url, bytes.NewBufferString(""))
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			//assert.JSONEq(t, tt.want.response, w.Body.String())
			assert.Equal(t, tt.want.code, w.Code)
		})

	}

}
