package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/learies/go-shortener/internal/service"
)

// mockService — простая реализация интерфейса ShortenerService для тестов.
type mockService struct {
	createResult string
	getResult    string
	getErr       error
}

func (m *mockService) Create(originalURL string) string {
	return m.createResult
}

func (m *mockService) Get(shortID string) (string, error) {
	return m.getResult, m.getErr
}

func TestCreateShortURL(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		mockShortID  string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid URL",
			body:         "https://practicum.yandex.ru/",
			mockShortID:  "QrPnX5IU",
			expectedCode: http.StatusCreated,
			expectedBody: "http://localhost:8080/QrPnX5IU",
		},
		{
			name:         "Empty Body",
			body:         "",
			mockShortID:  "",
			expectedCode: http.StatusBadRequest,
			expectedBody: "empty URL\n", // http.Error добавляет перенос строки
		},
	}

	for _, tt := range tests {
		// t.Run создает подтест. Если он упадет, вы увидите имя кейса в выводе
		t.Run(tt.name, func(t *testing.T) {
			// 1. Создаем мок сервиса с нужными для теста данными
			svc := &mockService{createResult: tt.mockShortID}
			h := New(svc)

			// 2. Формируем поддельный HTTP запрос
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))

			// 3. Создаем Recorder для перехвата ответа хендлера
			rr := httptest.NewRecorder()

			// 4. Вызываем наш хендлер, передавая ему фейковый запрос и рекордер
			h.CreateShortURL(rr, req)

			// 5. Проверяем результаты (ассерты)
			if rr.Code != tt.expectedCode {
				t.Errorf("ожидался статус %d, получен %d", tt.expectedCode, rr.Code)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("ожидалось тело %q, получено %q", tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func TestRedirect(t *testing.T) {
	tests := []struct {
		name           string
		shortID        string
		mockOriginal   string
		mockErr        error
		expectedCode   int
		expectedHeader string
	}{
		{
			name:           "Valid Redirect",
			shortID:        "QrPnX5IU",
			mockOriginal:   "https://practicum.yandex.ru/",
			mockErr:        nil,
			expectedCode:   http.StatusTemporaryRedirect,
			expectedHeader: "https://practicum.yandex.ru/",
		},
		{
			name:           "Not Found",
			shortID:        "unknown",
			mockOriginal:   "",
			mockErr:        service.ErrNotFound,
			expectedCode:   http.StatusNotFound,
			expectedHeader: "",
		},
		{
			name:           "Empty ID",
			shortID:        "",
			mockOriginal:   "",
			mockErr:        nil,
			expectedCode:   http.StatusBadRequest,
			expectedHeader: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &mockService{getResult: tt.mockOriginal, getErr: tt.mockErr}
			h := New(svc)

			req := httptest.NewRequest(http.MethodGet, "/"+tt.shortID, nil)

			req.SetPathValue("shortID", tt.shortID)

			rr := httptest.NewRecorder()

			h.Redirect(rr, req)

			if rr.Code != tt.expectedCode {
				t.Errorf("ожидался статус %d, получен %d", tt.expectedCode, rr.Code)
			}

			if loc := rr.Header().Get("Location"); loc != tt.expectedHeader {
				t.Errorf("ожидался заголовок Location %q, получен %q", tt.expectedHeader, loc)
			}
		})
	}
}
