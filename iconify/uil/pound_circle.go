package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoundCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 15h-5.18a3 3 0 0 0 .18-1v-1h2.5a1 1 0 0 0 0-2H11v-1a1.95 1.95 0 0 1 3.63-1a1 1 0 0 0 1.74-1A4 4 0 0 0 9 10v1H8a1 1 0 0 0 0 2h1v1a1 1 0 0 1-1 1a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2ZM12 1a11 11 0 1 0 11 11A11 11 0 0 0 12 1Zm0 20a9 9 0 1 1 9-9a9 9 0 0 1-9 9Z"/>`),
		g.Group(children),
	)
}