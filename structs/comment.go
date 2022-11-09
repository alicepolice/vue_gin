package structs

type Comment struct {
	ID   int    `json:"id"`
	UID  int    `json:"uid"`
	Name string `json:"name"`
	Text string `json:"text"`
	Pid  int    `json:"pid"`
	Date int64  `json:"date"`
}
