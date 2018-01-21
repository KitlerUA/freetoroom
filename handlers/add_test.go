package handlers

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"net/http"
	"strings"
)

var roomJSON = `{
		"room":"77",
		"name":"test_user""
	}`

var mockdb DBMock

func TestHandler_AddRoom(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(roomJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockdb.DB = make(map[string]map[int]string)
	mockdb.DB["rooms"] = make(map[int]string)
	mockdb.DB["rooms"][47] = "hitman"
	h := &Handler{DB: &mockdb}

	testCases := []struct {
		name string,
		username string
		password string
		expected 
	}
}