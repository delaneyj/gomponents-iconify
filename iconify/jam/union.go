package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Union(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.858 2.48L10 2.888l-.858-.408a5 5 0 1 0 0 9.04l.858-.408l.858.408a5 5 0 1 0 0-9.04ZM13 0a7 7 0 1 1-3 13.326A7 7 0 1 1 10 .673A6.973 6.973 0 0 1 13 0Z"/>`),
		g.Group(children),
	)
}