package structs

type Post struct {
	ID        int64  `db:"id" json:"id"`
	Bgcolor   string `db:"bgcolor" json:"bgcolor"`
	Textcolor string `db:"textcolor" json:"textcolor"`
	Headimg   string `db:"headimg" json:"headimg"`
	Videosrc  string `db:"videosrc" json:"videosrc"`
	Imgs      string `db:"imgs" json:"imgs"`
	Html      string `db:"html" json:"html"`
}
