package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicAccidentalNatural(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 8.75V3.5H8v14l6-2.25v5.25h2v-14l-6 2.25m4 4.5l-4 1.5v-4l4-1.5v4Z"/>`),
		g.Group(children),
	)
}