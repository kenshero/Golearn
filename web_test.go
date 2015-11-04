package main

import "io/ioutil"
import "net/http/httptest"
import "net/http"
import "testing"
import "contact"

func TestGetContact(t *testing.T) {
		contact.Add("kenshero",contact.Contact{
			Tel		:"123456",
			Email	:"ken@mail.com",
		})	

		ts := httptest.NewServer(http.HandlerFunc(getContact))
		defer ts.Close()

		resp, _ := http.Get(ts.URL + "/contacts/get?name=kenshero")
		b, 	_	:= ioutil.ReadAll(resp.Body)

		content := string(b)
		expectContent := `{"tel":"123456","mail":"ken@mail.com"}`+"\n"

		if content != expectContent{
			t.Error("Expect %s but got %s",expectContent,content)
		}
}

func TestPostContact(t *testing.T) {
		contact.Add("kenshero",contact.Contact{
			Tel		:"123456",
			Email	:"ken@mail.com",
		})	

		ts := httptest.NewServer(http.HandlerFunc(getContact))
		defer ts.Close()

		resp, _ := http.POST(ts.URL + "/contacts/get?name=kenshero","application/json",nil)

		if resp.StatusCode != http.StatusMethodNotAllowed{
			t.Error("Expect 405 but got %d",resp.StatusCode)
		}
		
		expectContent := `{"tel":"123456","mail":"ken@mail.com"}`+"\n"

		if content != expectContent{
			t.Error("Expect %s but got %s",expectContent,content)
		}
}