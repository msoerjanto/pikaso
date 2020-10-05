package data

//
import (
	"testing"
)

// Delete all artists from database
func ArtistDeleteAll() (err error) {
	db := db()
	defer db.Close()
	statement := "delete from artists"
	_, err = db.Exec(statement)
	if err != nil {
		return
	}
	return
}

func Test_CreateArtist(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	conv, err := users[0].CreateArtist("My first artist")
	if err != nil {
		t.Error(err, "Cannot create artist")
	}
	if conv.UserId != users[0].Id {
		t.Error("User not linked with artist")
	}
}
