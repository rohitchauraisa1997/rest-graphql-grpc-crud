package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Video struct {
	ID     string `json:"id" bson:"_id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author *User  `json:"author"`
}
