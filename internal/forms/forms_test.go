package forms

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedValues := url.Values{}
	postedValues.Add("a", "a")
	postedValues.Add("b", "a")
	postedValues.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedValues
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}

}

func TestForm_Has(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	has := form.Has("a")
	if has {
		t.Error("form shows field is present when it's in fact not")
	}

	postedValues = url.Values{}
	postedValues.Add("a", "a")
	form = New(postedValues)

	has = form.Has("a")
	if !has {
		t.Error("shows field is abasent when it's not")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.MinLength("a", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	err := form.Errors.Get("a")
	if err == "" {
		t.Error("should have an error but did not get one")
	}

	postedValues = url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 met when data is shorter")
	}
    
	postedValues = url.Values{}
	postedValues.Add("name", "Abiodun")
	form = New(postedValues)

	form.MinLength("name", 3)
	if !form.Valid() {
		t.Error("form shows field length is less than min length when it's not")
	}

	err = form.Errors.Get("name")
	if err != "" {
		t.Error("should have no error but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	sampleInvalidEmails := []string {
		"", "osisupermoses@", "abc.com", "@gmail.com",
	}
	sampleValidEmails := []string {
		"example@gmail.com", "osisupermoses@gmail.com", "techiehub@sample.com",
	}
	
	postedValues = url.Values{}
	for i, email := range sampleInvalidEmails {
		field := fmt.Sprintf("email%d", i)
		postedValues.Add(field, email)
		form = New(postedValues)

		form.IsEmail(field)
		if form.Valid() {
			t.Error("invalid emails were passed as valid")
		}
	}

	postedValues = url.Values{}
	for i, email := range sampleValidEmails {
		field := fmt.Sprintf("email%d", i)
		postedValues.Add(field, email)
		form := New(postedValues)

		form.IsEmail(field)
		if !form.Valid() {
			t.Error("valid emails were detected to be invalid")
		}
	}
}
