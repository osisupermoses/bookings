package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"reflect"
// 	"strings"
// 	"testing"

// 	"github.com/osisupermoses/bookings/internal/driver"
// 	"github.com/osisupermoses/bookings/internal/models"
// )

// type postData struct {
// 	key   string
// 	value string
// }

// var theTests = []struct {
// 	name               string
// 	url                string
// 	method             string
// 	expectedStatusCode int
// }{
// 	{"home", "/", "GET", http.StatusOK},
// 	{"about", "/about", "GET", http.StatusOK},
// 	{"gq", "/generals-quarters", "GET", http.StatusOK},
// 	{"ms", "/majors-suite", "GET", http.StatusOK},
// 	{"sa", "/search-availability", "GET", http.StatusOK},
// 	{"contact", "/contact", "GET", http.StatusOK},
// 	{"non-existent", "/green/eggs/and/ham", "GET", http.StatusNotFound},
// 	// new routes
// 	{"login", "/user/login", "GET", http.StatusOK},
// 	{"logout", "/user/logout", "GET", http.StatusOK},
// 	{"dashboard", "/admin/dashboard", "GET", http.StatusOK},
// 	{"new res", "/admin/reservations-new", "GET", http.StatusOK},
// 	{"all res", "/admin/reservations-all", "GET", http.StatusOK},
// 	{"show res", "/admin/reservations/new/1/show", "GET", http.StatusOK},
// 	// {"mr", "/make-reservation", "GET", http.StatusOK},
// 	// {"rs", "/reservation-summary", "GET", http.StatusOK},

// 	// {"post-search-avail", "/search-availability", "POST", []postData{
// 	// 	{key: "start", value: "2020-01-01"},
// 	// 	{key: "end", value: "2020-01-01"},
// 	// }, http.StatusOK},
// 	// {"post-search-avail-json", "/search-availability-json", "POST", []postData{
// 	// 	{key: "start", value: "2020-01-01"},
// 	// 	{key: "end", value: "2020-01-01"},
// 	// }, http.StatusOK},
// 	// {"post-make-reservation", "/make-reservation", "POST", []postData{
// 	// 	{key: "first_name", value: "John"},
// 	// 	{key: "last_name", value: "Smith"},
// 	// 	{key: "email", value: "me@here.com"},
// 	// 	{key: "phone", value: "08038223961"},
// 	// }, http.StatusOK},
// }

// func TestHandlers(t *testing.T) {
// 	routes := getRoutes()
// 	ts := httptest.NewTLSServer(routes)
// 	defer ts.Close()

// 	for _, e := range theTests {
// 		if e.method == "GET" {
// 			res, err := ts.Client().Get(ts.URL + e.url)
// 			if err != nil {
// 				t.Log(err)
// 				t.Fatal(err)
// 			}

// 			if res.StatusCode != e.expectedStatusCode {
// 				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
// 			}
// 		}

// 		/* no longer running post request tests here */
// 		// else {
// 		// 	values := url.Values{}
// 		// 	for _, x := range e.params {
// 		// 		values.Add(x.key, x.value)
// 		// 	}
// 		// 	res, err := ts.Client().PostForm(ts.URL + e.url, values)
// 		// 	if err != nil {
// 		// 		t.Log(err)
// 		// 		t.Fatal(err)
// 		// 	}

// 		// 	if res.StatusCode != e.expectedStatusCode {
// 		// 		t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
// 		// 	}
// 		// }
// 	}
// }

// func TestRepository_Reservation(t *testing.T) {
// 	reservation := models.Reservation{
// 		RoomID: 1,
// 		Room: models.Room{
// 			ID:       1,
// 			RoomName: "General's Quarters",
// 		},
// 	}

// 	req, _ := http.NewRequest("GET", "/make-reservation", nil)
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	rr := httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler := http.HandlerFunc(Repo.Reservation)

// 	handler.ServeHTTP(rr, req)
// 	if rr.Code != http.StatusOK {
// 		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
// 	}

// 	// test case where reservation is not in session (reset everything)
// 	req, _ = http.NewRequest("GET", "/make-reservation", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	rr = httptest.NewRecorder()

// 	handler.ServeHTTP(rr, req)
// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test case with non-existent room
// 	req, _ = http.NewRequest("GET", "/make-reservation", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	rr = httptest.NewRecorder()
// 	reservation.RoomID = 100
// 	session.Put(ctx, "reservation", reservation)

// 	handler.ServeHTTP(rr, req)
// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// }

// func TestRepository_PostReservation(t *testing.T) {
// 	reservation := models.Reservation{
// 		RoomID: 1,
// 		Room: models.Room{
// 			ID:       1,
// 			RoomName: "General's Quarters",
// 		},
// 	}

// 	// reqBody := "start_date=2050-01-01"
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=John@smith.com")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	// Alternative way (easier way) instead of using reqBody up here
// 	postedData := url.Values{}
// 	postedData.Add("start_date", "2050-01-01")
// 	postedData.Add("end_date", "2050-01-02")
// 	postedData.Add("first_name", "John")
// 	postedData.Add("last_name", "Smith")
// 	postedData.Add("email", "john@smith.com")
// 	postedData.Add("phone", "555-555-5555")
// 	postedData.Add("room_id", "1")

// 	// test case where post body is present and parsable [everything works as it should]
// 	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr := httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler := http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)
// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test for missing post body
// 	req, _ = http.NewRequest("POST", "/make-reservation", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler = http.HandlerFunc(Repo.PostReservation)
// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test case where reservation is not in session
// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)
// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test for invalid start date [not available for our usecase here anymore]
// 	// reqBody = "start_date=invalid"
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=John@smith.com")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	// req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	// ctx = getCtx(req)
// 	// req = req.WithContext(ctx)
// 	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	// rr = httptest.NewRecorder()

// 	// handler = http.HandlerFunc(Repo.PostReservation)
// 	// handler.ServeHTTP(rr, req)

// 	// if rr.Code != http.StatusSeeOther {
// 	// 	t.Errorf("Reservation handler returned wrong response code for invalid start date: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	// }

// 	// test for invalid end date [not available for our usecase here anymore]
// 	// reqBody = "start_date=2050-01-01"
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=invalid")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=John@smith.com")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	// req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	// ctx = getCtx(req)
// 	// req = req.WithContext(ctx)
// 	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	// rr = httptest.NewRecorder()

// 	// handler = http.HandlerFunc(Repo.PostReservation)
// 	// handler.ServeHTTP(rr, req)

// 	// if rr.Code != http.StatusSeeOther {
// 	// 	t.Errorf("Reservation handler returned wrong response code for invalid end date: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	// }

// 	// test for invalid room id [not available for our usecase here anymore]
// 	// reqBody = "start_date=2050-01-01"
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=John@smith.com")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
// 	// reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=invalid")

// 	// req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	// ctx = getCtx(req)
// 	// req = req.WithContext(ctx)
// 	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	// rr = httptest.NewRecorder()

// 	// handler = http.HandlerFunc(Repo.PostReservation)
// 	// handler.ServeHTTP(rr, req)

// 	// if rr.Code != http.StatusSeeOther {
// 	// 	t.Errorf("Reservation handler returned wrong response code for invalid room id: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	// }

// 	// test for invalid form data
// 	reqBody := "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=J")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=John@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler = http.HandlerFunc(Repo.PostReservation)
// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Reservation handler returned wrong response code for invalid form data: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test for failure to insert reservation into database
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=John@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=2")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()
// 	reservation.RoomID = 2
// 	session.Put(ctx, "reservation", reservation)

// 	handler = http.HandlerFunc(Repo.PostReservation)
// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Reservation handler failed when trying to fail inserting reservation: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test for failure to insert restriction into database
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=John@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()
// 	reservation.RoomID = 1000 // room id changes here
// 	session.Put(ctx, "reservation", reservation)

// 	handler = http.HandlerFunc(Repo.PostReservation)
// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Reservation handler failed when trying to fail inserting restriction: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}
// }

// func TestNewRepo(t *testing.T) {
// 	var db driver.DB
// 	testRepo := NewRepo(&app, &db)

// 	if reflect.TypeOf(testRepo).String() != "*handlers.Repository" {
// 		t.Errorf("Did not get correct type from NewRepo: got %s, wanted *Repository", reflect.TypeOf(testRepo).String())
// 	}
// }

// func TestRepository_PostAvailability(t *testing.T) {
// 	/*****************************************
// 	// first case -- rooms are not available
// 	*****************************************/
// 	// create our request body
// 	reqBody := "start=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2050-01-02")

// 	// create our request
// 	req, _ := http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr := httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler := http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have no rooms available, we expect to get status http.StatusSeeOther
// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Post availability when no rooms available gave wrong status code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	/*****************************************
// 	// second case -- rooms are available
// 	*****************************************/
// 	// this time, we specify a start date before 2040-01-01, which will give us
// 	// a non-empty slice, indicating that rooms are available
// 	reqBody = "start=2040-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2040-01-02")

// 	// create our request
// 	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusOK
// 	if rr.Code != http.StatusOK {
// 		t.Errorf("Post availability when rooms are available gave wrong status code: got %d, wanted %d", rr.Code, http.StatusOK)
// 	}

// 	/*****************************************
// 	// third case -- empty post body
// 	*****************************************/
// 	// create our request with a nil body, so parsing form fails
// 	req, _ = http.NewRequest("POST", "/search-availability", nil)

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusSeeOther
// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Post availability with empty request body (nil) gave wrong status code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	/*****************************************
// 	// fourth case -- start date in wrong format
// 	*****************************************/
// 	// this time, we specify a start date in the wrong format
// 	reqBody = "start=invalid"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2040-01-02")
// 	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusSeeOther
// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Post availability with invalid start date gave wrong status code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	/*****************************************
// 	// fifth case -- end date in wrong format
// 	*****************************************/
// 	// this time, we specify a start date in the wrong format
// 	reqBody = "start=2040-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "invalid")
// 	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusSeeOther
// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Post availability with invalid end date gave wrong status code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	/*****************************************
// 	// sixth case -- database query fails
// 	*****************************************/
// 	// this time, we specify a start date of 2060-01-01, which will cause
// 	// our testdb repo to return an error
// 	reqBody = "start=2060-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2060-01-02")
// 	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.PostAvailability)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have rooms available, we expect to get status http.StatusSeeOther
// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("Post availability when database query fails gave wrong status code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}
// }

// func TestRepository_AvailabilityJSON(t *testing.T) {
// 	/*****************************************
// 	// first case -- rooms are not available
// 	*****************************************/
// 	// create our request body
// 	reqBody := "start=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	// create our request
// 	req, _ := http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr := httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler := http.HandlerFunc(Repo.AvailabilityJSON)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// since we have no rooms available, we expect to get status http.StatusSeeOther
// 	// this time we want to parse JSON and get the expected response
// 	var j jsonResponse
// 	err := json.Unmarshal([]byte(rr.Body.String()), &j)
// 	if err != nil {
// 		t.Error("failed to parse json!")
// 	}

// 	// since we specified a start date > 2049-12-31, we expect no availability
// 	if j.OK {
// 		t.Error("Got availability when none was expected in AvailabilityJSON")
// 	}

// 	/*****************************************
// 	// second case -- rooms not available
// 	*****************************************/
// 	// create our request body
// 	reqBody = "start=2040-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2040-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	// create our request
// 	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.AvailabilityJSON)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// this time we want to parse JSON and get the expected response
// 	err = json.Unmarshal([]byte(rr.Body.String()), &j)
// 	if err != nil {
// 		t.Error("failed to parse json!")
// 	}

// 	// since we specified a start date < 2049-12-31, we expect availability
// 	if !j.OK {
// 		t.Error("Got no availability when some was expected in AvailabilityJSON")
// 	}

// 	/*****************************************
// 	// third case -- no request body
// 	*****************************************/
// 	// create our request
// 	req, _ = http.NewRequest("POST", "/search-availability-json", nil)

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.AvailabilityJSON)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// this time we want to parse JSON and get the expected response
// 	err = json.Unmarshal([]byte(rr.Body.String()), &j)
// 	if err != nil {
// 		t.Error("failed to parse json!")
// 	}

// 	// since we specified a start date < 2049-12-31, we expect availability
// 	if j.OK || j.Message != "Internal server error" {
// 		t.Error("Got availability when request body was empty")
// 	}

// 	/*****************************************
// 	// fourth case -- database error
// 	*****************************************/
// 	// create our request body
// 	reqBody = "start=2060-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2060-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")
// 	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

// 	// get the context with session
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	// set the request header
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// create our response recorder, which satisfies the requirements
// 	// for http.ResponseWriter
// 	rr = httptest.NewRecorder()

// 	// make our handler a http.HandlerFunc
// 	handler = http.HandlerFunc(Repo.AvailabilityJSON)

// 	// make the request to our handler
// 	handler.ServeHTTP(rr, req)

// 	// this time we want to parse JSON and get the expected response
// 	err = json.Unmarshal([]byte(rr.Body.String()), &j)
// 	if err != nil {
// 		t.Error("failed to parse json!")
// 	}

// 	// since we specified a start date < 2049-12-31, we expect availability
// 	if j.OK || j.Message != "Error querying database" {
// 		t.Error("Got availability when simulating database error")
// 	}
// }

// func TestRepository_ReservationSummary(t *testing.T) {
// 	/*****************************************
// 	// first case -- reservation in session
// 	*****************************************/
// 	reservation := models.Reservation{
// 		RoomID: 1,
// 		Room: models.Room{
// 			ID:       1,
// 			RoomName: "General's Quarters",
// 		},
// 	}

// 	req, _ := http.NewRequest("GET", "/reservation-summary", nil)
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	rr := httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler := http.HandlerFunc(Repo.ReservationSummary)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusOK {
// 		t.Errorf("ReservationSummary handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
// 	}

// 	/*****************************************
// 	// second case -- reservation not in session
// 	*****************************************/
// 	req, _ = http.NewRequest("GET", "/reservation-summary", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.ReservationSummary)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("ReservationSummary handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
// 	}
// }

// func TestRepository_ChooseRoom(t *testing.T) {
// 	/*****************************************
// 	// first case -- reservation in session
// 	*****************************************/
// 	reservation := models.Reservation{
// 		RoomID: 1,
// 		Room: models.Room{
// 			ID:       1,
// 			RoomName: "General's Quarters",
// 		},
// 	}

// 	req, _ := http.NewRequest("GET", "/choose-room/1", nil)
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)
// 	// set the RequestURI on the request so that we can grab the ID
// 	// from the URL
// 	req.RequestURI = "/choose-room/1"

// 	rr := httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler := http.HandlerFunc(Repo.ChooseRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	///*****************************************
// 	//// second case -- reservation not in session
// 	//*****************************************/
// 	req, _ = http.NewRequest("GET", "/choose-room/1", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.RequestURI = "/choose-room/1"

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.ChooseRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	///*****************************************
// 	//// third case -- missing url parameter, or malformed parameter
// 	//*****************************************/
// 	req, _ = http.NewRequest("GET", "/choose-room/fish", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.RequestURI = "/choose-room/fish"

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.ChooseRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}
// }

// func TestRepository_BookRoom(t *testing.T) {
// 	/*****************************************
// 	// first case -- database works
// 	*****************************************/
// 	reservation := models.Reservation{
// 		RoomID: 1,
// 		Room: models.Room{
// 			ID:       1,
// 			RoomName: "General's Quarters",
// 		},
// 	}

// 	req, _ := http.NewRequest("GET", "/book-room?s=2050-01-01&e=2050-01-02&id=1", nil)
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	rr := httptest.NewRecorder()
// 	session.Put(ctx, "reservation", reservation)

// 	handler := http.HandlerFunc(Repo.BookRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	/*****************************************
// 	// second case -- database failed
// 	*****************************************/
// 	req, _ = http.NewRequest("GET", "/book-room?s=2040-01-01&e=2040-01-02&id=4", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)

// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.BookRoom)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
// 	}
// }

// var loginTests = []struct {
// 	name string
// 	email string
// 	expectedStatusCode int
// 	expectedHTML string
// 	expectedLocation string
// } {
// 	{
// 		"valid-credentials",
// 		"me@here.com",
// 		http.StatusSeeOther,
// 		"",
// 		"/",
// 	},
// 	{
// 		"invalid-credentials",
// 		"jack@nimble.com",
// 		http.StatusSeeOther,
// 		"",
// 		"/user/login",
// 	},
// 	{
// 		"invalid-data",
// 		"j",
// 		http.StatusOK, // because we are not doing a redirect or leaving the page at all
// 		`action="/user/login`,
// 		"",
// 	},
// }

// func TestLogin(t *testing.T) {
// 	// range through all tests
// 	for _, e := range loginTests {
// 		postedData := url.Values{}
// 		postedData.Add("email", e.email)
// 		postedData.Add("password", "password")

// 		// create the request
// 		req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(postedData.Encode()))
// 		ctx := getCtx(req)
// 		req = req.WithContext(ctx)

// 		// set the header
// 		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 		// set the recorder
// 		rr := httptest.NewRecorder()
		
// 		// set the handler
// 		handler := http.HandlerFunc(Repo.PostShowLogin)
// 		handler.ServeHTTP(rr, req)

// 		// test expected code
// 		if rr.Code != e.expectedStatusCode {
// 			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
// 		}

// 		// test expected location
// 		if e.expectedLocation != "" {
// 			// get the url from test
// 			actualLoc, _ := rr.Result().Location()
// 			if actualLoc.String() != e.expectedLocation {
// 				t.Errorf("failed %s: expected location %s, but got %s", e.name, e.expectedLocation, actualLoc.String())
// 			}
// 		}

// 		// test checking for expected values in HTML
// 		if e.expectedHTML != "" {
// 			// read the response body into a string
// 			html := rr.Body.String()
// 			if !strings.Contains(html,e.expectedHTML) {
// 				t.Errorf("failed %s: expected to find %s but did not", e.name, e.expectedHTML)
// 			}
// 		}
// 	}
// }

// func getCtx(req *http.Request) context.Context {
// 	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return ctx
// }
