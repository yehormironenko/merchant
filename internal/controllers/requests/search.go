package requests

import "github.com/rs/zerolog"

type SearchBook struct {
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
}

func (s SearchBook) MarshalZerologObject(e *zerolog.Event) {
	e.Str("Title", s.Title).
		Str("Genre", s.Genre).
		Str("Author", s.Author)
}
