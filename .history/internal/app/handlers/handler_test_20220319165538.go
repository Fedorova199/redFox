package handlers_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)
 md := handlers.NewModels()
func TestHandlerPost(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	// создаём массив тестов: имя и желаемый результат
	tests := []struct {
		name string
		want want
	}{
		// определяем все тесты
		{
			name: "positive test #1",
			want: want{
				code:        201,
				response:    "http://localhost:8080/1",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name: "positive test #2",
			want: want{
				code:        201,
				response:    "http://localhost:8080/2",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name: "positive test #3",
			want: want{
				code:        201,
				response:    "http://localhost:8080/3",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name: "positive test #3",
			want: want{
				code:        201,
				response:    "http://localhost:8080/4",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name: "positive test #3",
			want: want{
				code:        201,
				response:    "http://localhost:8080/5",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			var body = []byte("https://practicum.yandex.ru/")
			request := httptest.NewRequest(http.MethodPost, tt.want.response, bytes.NewBuffer(body))

			// создаём новый Recorder
			w := httptest.NewRecorder()
			// определяем хендлер
			h := http.HandlerFunc(md.HandlerPost)
			// запускаем сервер
			h.ServeHTTP(w, request)
			res := w.Result()

			// проверяем код ответа
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(resBody) != tt.want.response {
				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
			}

			// заголовок ответа
			if res.Header.Get("Content-Type") != tt.want.contentType {
				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
			}
		})
	}
}

func TestHandlerGet(t *testing.T) {

	type want struct {
		code        int
		response    string
		contentType string
	}
	// создаём массив тестов: имя и желаемый результат
	tests := []struct {
		name    string
		request string
		want    want
	}{
		// определяем все тесты
		{
			name:    "positive test #1",
			request: "/1",
			want: want{
				code:        307,
				response:    "",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "positive test #2",
			request: "/2",
			want: want{
				code:        307,
				response:    "",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "Invalidtest #1",
			request: "/88",
			want: want{
				code:        400,
				response:    "invalid key",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodGet, tt.request, nil)

			// создаём новый Recorder
			w := httptest.NewRecorder()
			// определяем хендлер
			h := http.HandlerFunc(md.HandlerGet)
			// запускаем сервер
			h.ServeHTTP(w, request)
			res := w.Result()

			// проверяем код ответа
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(resBody) != tt.want.response {
				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
			}

			// заголовок ответа
			if res.Header.Get("Content-Type") != tt.want.contentType {
				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
			}
		})
	}
}
