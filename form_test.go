package gosurf

import (
	"fmt"
	"github.com/headzoo/gosurf/unittest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFormAttribs(t *testing.T) {
	unittest.Run(t)
	unittest.Run(t)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, htmlForm)
	}))
	defer ts.Close()

	b := NewBrowser()
	b.Get(ts.URL)
	f, err := b.Form("[name='default']")
	unittest.AssertEquals(nil, err)
	unittest.AssertEquals("POST", f.Method())
	unittest.AssertEquals(ts.URL+"/", f.Action())

}