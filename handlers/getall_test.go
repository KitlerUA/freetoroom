package handlers

import (
	"testing"
	"net/http"
	"github.com/labstack/echo"
	"net/http/httptest"
)

func TestHandler_GetAll(t *testing.T) {
	var mockdb = DBMock{}
	mockdb.DB = make(map[string]map[int]string)
	mockdb.DB["rooms"] = make(map[int]string)
	mockdb.DB["rooms"][47] = "hitman"
	h := &Handler{DB: &mockdb}

	testCases := []struct {
		name string
		expected int
	}{
		{"normal case", http.StatusOK},
	}
	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			e := echo.New()
			req := httptest.NewRequest(echo.POST, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if err := h.GetAll(c); c.Response().Status != testLocal.expected{
				t.Errorf("Expected %s , got %s. Err %s", testLocal.expected, c.Response().Status, err)
			}
		})
	}
}
