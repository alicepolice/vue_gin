package structs

type Gameblog struct {
	// ID int64 `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Text  string `db:"text" json:"text"`
	Img   string `db:"img" json:"img"`
}
