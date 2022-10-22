package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonTest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/json", JsonTest)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/json", nil)
	mux.ServeHTTP(recorder, request)

	if recorder.Code != 200 {
		t.Errorf("Response code is %v", recorder.Code)
	}

	var post Post
	json.Unmarshal(recorder.Body.Bytes(), &post)

	if post.Name != "急速小子" {
		t.Error("未能完成数据读取")
	}

}
