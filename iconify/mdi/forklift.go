package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forklift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4v7H4c-1.11 0-2 .89-2 2v4a3 3 0 0 0 3 3a3 3 0 0 0 3-3h2a3 3 0 0 0 3 3a3 3 0 0 0 3-3v-4l-4-9H6m11 1v14h5v-1.5h-3.5V5H17m-9.5.5h3.7l3.3 7.5h-7V5.5M5 15.5A1.5 1.5 0 0 1 6.5 17A1.5 1.5 0 0 1 5 18.5A1.5 1.5 0 0 1 3.5 17A1.5 1.5 0 0 1 5 15.5m8 0a1.5 1.5 0 0 1 1.5 1.5a1.5 1.5 0 0 1-1.5 1.5a1.5 1.5 0 0 1-1.5-1.5a1.5 1.5 0 0 1 1.5-1.5Z"/>`),
		g.Group(children),
	)
}