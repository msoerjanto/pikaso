package data

import "time"

type Piece struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ArtistId  int
	CreatedAt time.Time
}

func (piece *Piece) CreatedAtDate() string {
	return piece.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// Get the user who wrote the piece
func (piece *Piece) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", piece.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
