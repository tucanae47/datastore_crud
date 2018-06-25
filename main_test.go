package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	var jsonStr = []byte(`{"Brand":"gw","Color":"White","Serial":"4701"}`)
	res, err := http.Post("https://merlin-208202.appspot.com/merlin/new", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUpdate(t *testing.T) {
	var expected = []byte(`{"Brand":"gw","Color":"White","Serial":"4702"}`)
	res, err := http.Post("https://merlin-208202.appspot.com/merlin/update/5631986051842048", "application/json", bytes.NewBuffer(expected))
	if err != nil {
		t.Fatal(err)
	}
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	update, err := http.Get("https://merlin-208202.appspot.com/merlin/5631986051842048")
	defer update.Body.Close()
	data, err := ioutil.ReadAll(update.Body)
	data = data[:len(data)-1]
	if string(data) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", data, expected)
	}

}

func TestDelete(t *testing.T) {
	res, err := http.Get("https://merlin-208202.appspot.com/merlin/delete/5631986051842048")
	if err != nil {
		t.Fatal(err)
	}
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestRetrieve(t *testing.T) {
	res, err := http.Get("https://merlin-208202.appspot.com/merlin/5631986051842048")
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	data = data[:len(data)-1]
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := []byte(`{"Brand":"gw","Color":"White","Serial":"4701"}`)
	if string(data) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", data, expected)
	}
}
