package handlers

import (
	"testing"
	"net/http"
	"github.com/labstack/echo"
	"net/http/httptest"
)

func TestHandler_Login(t *testing.T) {
	var mockdb = DBMock{}
	mockdb.DBAccounts = make(map[string]map[string]string)
	mockdb.DBAccounts["accounts"] = make(map[string]string)
	mockdb.DBAccounts["accounts"]["kitler"] = "secret"
	h := &Handler{DB: &mockdb}

	testCases := []struct {
		name string
		username string
		password string
		expected int
	}{
		{"normal case", "kitler", "secret", http.StatusOK},
		{"client not exist", "not kitler", "no matter", http.StatusUnauthorized},
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
			req.Form.Add("username", testLocal.username)
			req.Form.Add("password", testLocal.password)
			c := e.NewContext(req, rec)
			if _ = h.Login(c); c.Response().Status != testLocal.expected{
				t.Errorf("Expected %s , got %s", testLocal.expected, c.Response().Status)
			}
		})
	}
}