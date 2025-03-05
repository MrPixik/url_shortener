package server

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MrPixik/url_shortener/internal/app/middleware"
	"github.com/MrPixik/url_shortener/internal/app/models"
	"github.com/MrPixik/url_shortener/internal/config"
	"github.com/MrPixik/url_shortener/internal/db/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	config.InitConfig()
	middleware.InitLogger()
}

func createHash(url string) string {
	hasher := md5.New()
	return hex.EncodeToString(hasher.Sum([]byte(url))[0:12])
}

func TestMainPagePostHandler(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		response    string
	}
	tests := []struct {
		name   string
		want   want
		method string
		target string
		body   []byte
	}{
		{
			name: "Test OK",
			want: want{
				statusCode:  http.StatusCreated,
				contentType: "text/plain",
				response:    "http://localhost:8080/" + createHash("ok"),
			},
			method: http.MethodPost,
			target: "/",
			body:   []byte("ok"),
		},
		{
			name: "Test Empty Request",
			want: want{
				statusCode:  http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
				response:    "Empty originalURL\n",
			},
			method: http.MethodPost,
			target: "/",
			body:   []byte(""),
		},
	}

	//Mocks initialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDatabaseService(ctrl)

	m.EXPECT().
		CreateUrl(gomock.Any(), createHash("ok"), "ok").
		Return(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			request := httptest.NewRequest(tt.method, tt.target, bytes.NewBuffer(tt.body))
			response := httptest.NewRecorder()

			mainPagePostHandler(response, request, config.Cfg, m)
			result := response.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

			defer result.Body.Close()
			body, err := io.ReadAll(result.Body)
			require.NoError(t, err)

			assert.Equal(t, tt.want.response, string(body))

			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
		})
	}
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("EOF")
}

func TestMainPagePostBadRequestHandler(t *testing.T) {

	type want struct {
		statusCode int
	}
	tests := []struct {
		name   string
		want   want
		method string
		body   []byte
		target string
	}{
		{
			name: "Test Bad Request",
			want: want{
				statusCode: http.StatusBadRequest,
			},
			method: http.MethodPost,
			body:   nil,
			target: "/",
		},
	}

	//Mocks initialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDatabaseService(ctrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.target, &errorReader{})
			response := httptest.NewRecorder()

			mainPagePostHandler(response, request, config.Cfg, m)

			result := response.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

		})
	}
}

func TestShortenPostHandler(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        []byte
	}
	tests := []struct {
		name   string
		want   want
		method string
		target string
		body   []byte
	}{
		{
			name: "Test ok",
			want: want{
				statusCode:  http.StatusCreated,
				contentType: "application/json",
				body:        []byte("{\"short-url\":\"http://localhost:8080/" + createHash("ok") + "\"}"),
			},
			method: http.MethodPost,
			target: "/api/shorten",
			body:   []byte("{\"url\":\"ok\"}"),
		},
	}

	//Mocks initialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDatabaseService(ctrl)

	m.EXPECT().
		CreateUrl(gomock.Any(), createHash("ok"), "ok").
		Return(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.target, bytes.NewBuffer(tt.body))
			recorder := httptest.NewRecorder()

			shortenURLPostHandler(recorder, request, config.Cfg, m)

			response := recorder.Result()

			assert.Equal(t, tt.want.statusCode, response.StatusCode)

			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			require.NoError(t, err)

			fmt.Println(string(body))

			assert.Equal(t, tt.want.body, body)

			assert.Equal(t, tt.want.contentType, response.Header.Get("Content-Type"))
		})
	}
}

func TestMainPageGetHandler(t *testing.T) {
	type want struct {
		statusCode int
		Location   string
	}
	tests := []struct {
		name   string
		want   want
		method string
		target string
	}{
		{
			name: "Test OK",
			want: want{
				statusCode: http.StatusTemporaryRedirect,
				Location:   "https://practicum.yandex.ru/",
			},
			method: http.MethodGet,
			target: "http://localhost:8080/" + createHash("https://practicum.yandex.ru/"),
		},
		{
			name: "Test Bad Request",
			want: want{
				statusCode: http.StatusBadRequest,
			},
			method: http.MethodGet,
			target: "http://localhost:8080/" + "unknown_URL",
		},
		{
			name: "Test DB Error",
			want: want{
				statusCode: http.StatusInternalServerError,
			},
			method: http.MethodGet,
			target: "http://localhost:8080/" + "drop_database_url",
		},
	}

	//Mocks initialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockDatabaseService(ctrl)

	m.EXPECT().
		GetUrlByShortName(gomock.Any(), createHash("https://practicum.yandex.ru/")).
		Return(models.URLDB{Original: "https://practicum.yandex.ru/"}, nil)
	m.EXPECT().
		GetUrlByShortName(gomock.Any(), "unknown_URL").
		Return(models.URLDB{}, nil)
	m.EXPECT().
		GetUrlByShortName(gomock.Any(), "drop_database_url").
		Return(models.URLDB{}, errors.New("crash"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			router := InitHandlers(config.Cfg, middleware.Logger, m)

			getRequest := httptest.NewRequest(tt.method, tt.target, nil)
			getResonse := httptest.NewRecorder()
			router.ServeHTTP(getResonse, getRequest)

			getResult := getResonse.Result()

			assert.Equal(t, tt.want.statusCode, getResult.StatusCode)
			assert.Equal(t, tt.want.Location, getResult.Header.Get("Location"))
		})
	}
}
