package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/zelibobik/go_test_1/handlers"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	handlers.Home(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if have, want := string(greeting), "Hello! Your request was processed."; have != want {
		t.Errorf("The greeting is wrong. Have: %s, want: %s.", have, want)
	}
}
