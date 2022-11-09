package structs

type Gamelist struct {
	// ID    int64    `db:"id" json:"id"`
	Title string   `db:"title" json:"title"`
	Text  string   `db:"text" json:"text"`
	Img   string   `db:"img" json:"img"`
	Price float64  `db:"price" json:"price"`
	Tag   []string `db:"tag" json:"tag"` // 新添加
	Web   bool     `db:"Web" json:"web"`
}
