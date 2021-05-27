package franky

/*
A Record struct is tied to a User struct by its UserId field
*/
type Record struct {
	UserId   string
	Song     string
	Artist   string
	Album    string
	Playlist string
	Category string
	Duration uint
	Date     int64
}
