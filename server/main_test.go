package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T){
	ts := httptest.NewServer(App())
	defer ts.Close()

	Convey("HTTP Request E2E test", t, func() {
		Convey("#tracking", func() {
			Convey("Response has 404(not found) status code when invalid path", func(){
				res,err := http.Get(ts.URL + "/track")
				if err != nil {
					t.Fatal(err)
				}
				So(res.StatusCode, ShouldEqual, http.StatusNotFound)
			})

			
			Convey("Response has 400(not found courier code) status code when non-exist courier code", func(){
				res,err := http.Get(ts.URL + "/tracking/test/1")
				if err != nil {
					t.Fatal(err)
				}
				So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			})
		})
	})
}