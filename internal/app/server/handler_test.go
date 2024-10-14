package server

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createHash(url string) string {
	hasher := md5.New()
	return hex.EncodeToString(hasher.Sum([]byte(url))[:12])
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
				response:    "http://localhost:8080/" + createHash("https://practicum.yandex.ru/"),
			},
			method: http.MethodPost,
			target: "/",
			body:   []byte("https://practicum.yandex.ru/"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.target, bytes.NewBuffer(tt.body))
			response := httptest.NewRecorder()

			MainPagePostHandler(response, request)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.target, &errorReader{})
			response := httptest.NewRecorder()

			MainPagePostHandler(response, request)

			result := response.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

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
		}, {
			name: "Test Bad Request",
			want: want{
				statusCode: http.StatusBadRequest,
			},
			method: http.MethodGet,
			target: "http://localhost:8080/" + createHash("unknown URL"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			router := chi.NewRouter()
			router.Route("/", func(router chi.Router) {
				router.Get("/{id}", MainPageGetHandler)
				router.Post("/", MainPagePostHandler)
			})

			postRequest := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("https://practicum.yandex.ru/")))
			postResponse := httptest.NewRecorder()
			router.ServeHTTP(postResponse, postRequest)

			//fmt.Println("\"" + tt.target + "\"")
			getRequest := httptest.NewRequest(tt.method, tt.target, nil)
			getResonse := httptest.NewRecorder()
			router.ServeHTTP(getResonse, getRequest)

			getResult := getResonse.Result()

			assert.Equal(t, tt.want.statusCode, getResult.StatusCode)
			assert.Equal(t, tt.want.Location, getResult.Header.Get("Location"))
		})
	}
}
