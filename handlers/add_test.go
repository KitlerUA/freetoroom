package handlers

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"net/http"
)

var roomJSON = `{
		"room":"47",
		"name":"hitman""
	}`

var mockdb DBMock

func TestHandler_AddRoom(t *testing.T) {

	mockdb.DB = make(map[string]map[int]string)
	mockdb.DB["rooms"] = make(map[int]string)
	mockdb.DB["rooms"][47] = "hitman"
	h := &Handler{DB: &mockdb}

	testCases := []struct {
		name string
		room string
		client string
		expected int
	}{
		{"normal case", "48", "hitman8", http.StatusOK},
		{"alreafy exist", "47", "no matter", http.StatusInternalServerError},
		{"empty room field", "", "something", http.StatusBadRequest},
	}
	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			e := echo.New()
			req := httptest.NewRequest(echo.POST, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			req.Form = make(map[string][]string)
			req.Form.Add("room", testLocal.room)
			req.Form.Add("client", testLocal.client)
			c := e.NewContext(req, rec)
			if err := h.AddRoom(c); c.Response().Status != testLocal.expected{
				t.Errorf("Expected %s , got %s. Err %s", testLocal.expected, c.Response().Status, err)
			}
		})
	}
}