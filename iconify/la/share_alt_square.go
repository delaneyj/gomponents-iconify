package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAltSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm13 3a2 2 0 0 0-1.992 2.156L13.32 14.5A2 2 0 0 0 12 14a2 2 0 0 0 0 4a2 2 0 0 0 1.32-.5l4.688 2.344A2 2 0 0 0 20 22a2 2 0 0 0 0-4a2 2 0 0 0-1.32.5l-4.688-2.344a2 2 0 0 0 0-.312L18.68 13.5A2 2 0 0 0 20 14a2 2 0 0 0 0-4z"/>`),
		g.Group(children),
	)
}