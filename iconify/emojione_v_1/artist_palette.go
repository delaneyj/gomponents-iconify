package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtistPalette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#c2986b" d="M5.306 49.865a31.856 31.856 0 0 0 3.743 4.532c12.514 12.512 32.795 12.517 45.31 0c12.515-12.513 12.513-32.794-.002-45.31c-12.512-12.507-32.793-12.512-45.31 0C3.903 14.229.912 20.686-.003 27.375l14.247-2.61l-8.941 25.1"/><path fill="#02897d" d="M29.875 13.75a6.397 6.397 0 0 1-6.396 6.394a6.397 6.397 0 0 1 0-12.792a6.398 6.398 0 0 1 6.396 6.398"/><circle cx="40.875" cy="17.808" r="6.394" fill="#e45264"/><circle cx="49.981" cy="35.613" r="6.395" fill="#405967"/><path fill="#ed2e7c" d="M44.06 52.73a6.394 6.394 0 0 1-12.789 0a6.394 6.394 0 1 1 12.789 0"/><path fill="#e7e6e6" d="M23.482 49.22a6.395 6.395 0 1 1-12.79 0a6.395 6.395 0 0 1 12.79 0"/>`),
		g.Group(children),
	)
}