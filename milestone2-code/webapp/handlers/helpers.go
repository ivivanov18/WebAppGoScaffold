package handlers

import "os"

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {
	filename := "./static/" + title
	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
