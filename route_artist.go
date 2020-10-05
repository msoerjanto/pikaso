package main

import (
	"fmt"
	"net/http"

	"github.com/msoerjanto/pikaso/data"
)

// GET /artist/new
// Show the new artist form page
func newArtist(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.artist")
	}
}

// POST /signup
// Create the user account
func createArtist(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		firstName := request.FormValue("firstName")
		lastName := request.FormValue("lastName")
		description := request.FormValue("description")
		if _, err := user.CreateArtist(firstName, lastName, description); err != nil {
			danger(err, "Cannot create artist")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /artist/read
// Show the details of the artist, including the posts and the form to write a post
func readArtist(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	artist, err := data.ArtistByUUID(uuid)
	if err != nil {
		error_message(writer, request, "Cannot read artist")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &artist, "layout", "public.navbar", "public.artist")
		} else {
			generateHTML(writer, &artist, "layout", "private.navbar", "private.artist")
		}
	}
}

// POST /artist/post
// Create the post
func postPiece(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		body := request.FormValue("body")
		uuid := request.FormValue("uuid")
		artist, err := data.ArtistByUUID(uuid)
		if err != nil {
			error_message(writer, request, "Cannot read artist")
		}
		if _, err := user.CreatePiece(artist, body); err != nil {
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/artist/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
