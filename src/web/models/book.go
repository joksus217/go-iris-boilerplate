package models

type Book struct {
	Id       uint   `json:"id" form:"id"`
	AuthorID int    `json:"author_id" form:"author_id"`
	Name     string `json:"name" form:"name"`
	Genre    string `json:"genre" form:"genre"`
}

func (b Book) IsValid() bool {
	return b.Id > 0
}
