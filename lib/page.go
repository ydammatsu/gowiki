package lib

import "os"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	file_path := "data/" + p.Title + ".txt"
	return os.WriteFile(file_path, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	file_path := "data/" + title + ".txt"
	body, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
