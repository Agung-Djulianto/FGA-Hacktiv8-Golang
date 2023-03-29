package models

type BOOK struct {
	ID        string `json:"id"`
	Titel     string `json:"title"`
	Author    string `json:"author"`
	Deskripsi string `json:"deskripsi"`
}
