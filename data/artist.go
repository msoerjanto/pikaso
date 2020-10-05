package data

import (
	"log"
	"time"
)

type Artist struct {
	Id          int
	Uuid        string
	FirstName   string
	LastName    string
	Description string
	ProfilePic  string
	UserId      int
	CreatedAt   time.Time
}

// format the CreatedAt date to display nicely on the screen
func (artist *Artist) CreatedAtDate() string {
	return artist.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of pieces in a artist
func (artist *Artist) NumPieces() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM pieces where artist_id = $1", artist.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// get pieces to a artist
func (artist *Artist) Pieces() (pieces []Piece, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, artist_id, created_at FROM pieces where artist_id = $1", artist.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		piece := Piece{}
		if err = rows.Scan(&piece.Id, &piece.Uuid, &piece.Body, &piece.UserId, &piece.ArtistId, &piece.CreatedAt); err != nil {
			return
		}
		pieces = append(pieces, piece)
	}
	rows.Close()
	return
}

// Create a new artist
func (user *User) CreateArtist(firstName string, lastName string, description string, profilePic string) (conv Artist, err error) {
	statement := "insert into artists (uuid, first_name, last_name, description, profile_pic, user_id, created_at) values ($1, $2, $3, $4, $5, $6, $7) returning id, uuid, first_name, last_name, description, profile_pic, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), firstName, lastName, description, profilePic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.FirstName, &conv.LastName, &conv.Description, &conv.ProfilePic, &conv.UserId, &conv.CreatedAt)
	return
}

// Create a new piece to a artist
func (user *User) CreatePiece(conv Artist, body string) (piece Piece, err error) {
	statement := "insert into pieces (uuid, body, user_id, artist_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, artist_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&piece.Id, &piece.Uuid, &piece.Body, &piece.UserId, &piece.ArtistId, &piece.CreatedAt)
	return
}

// Get all artists in the database and returns it
func Artists() (artists []Artist, err error) {
	rows, err := Db.Query("SELECT id, uuid, first_name, last_name, description, profile_pic, user_id, created_at FROM artists ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Artist{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.FirstName, &conv.LastName, &conv.Description, &conv.ProfilePic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		artists = append(artists, conv)
	}
	rows.Close()
	return
}

// Get a artist by the UUID
func ArtistByUUID(uuid string) (conv Artist, err error) {
	conv = Artist{}
	err = Db.QueryRow("SELECT id, uuid, first_name, last_name, description, profile_pic, user_id, created_at FROM artists WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.FirstName, &conv.LastName, &conv.Description, &conv.ProfilePic, &conv.UserId, &conv.CreatedAt)
	return
}

// Get the user who started this artist
func (artist *Artist) User() (user User) {
	user = User{}
	err := Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", artist.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		log.Fatal("Problem reading user")
	}
	return
}
