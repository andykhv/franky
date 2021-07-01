package franky

type RecordRequest struct {
	UserId    string
	Song      string
	Artist    string
	Album     string
	Playlist  string
	Category  string
	StartDate int64
	EndDate   int64
}
