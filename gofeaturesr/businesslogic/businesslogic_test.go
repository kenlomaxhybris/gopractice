package businesslogic

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	myRouter.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestWebServer(t *testing.T) {
	SetUpRouterServer()
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	body := ""

	if body = response.Body.String(); body != "Hi there!" {
		t.Errorf("Expected Hi there!. Got %s", body)
	}

	log.Println("Body ", body)
}
