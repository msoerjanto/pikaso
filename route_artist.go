package main

import (
	"fmt"
	"net/http"

	"github.com/msoerjanto/pikaso/application"
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

// Create the artist
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
		ppFile, ppHeader, err := request.FormFile("profilePic")

		// Upload the image to S3
		profilePic := application.UploadMultiPartFileToS3(ppFile, ppHeader, err, "artist")

		// Create the artist
		firstName := request.PostFormValue("firstName")
		lastName := request.PostFormValue("lastName")
		description := request.PostFormValue("description")
		fmt.Println(ppFile)
		fmt.Println(ppHeader)
		fmt.Println(err)

		if _, err := user.CreateArtist(firstName, lastName, description, profilePic); err != nil {
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
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
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
