package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.29 12.29A1 1 0 0 0 12 14h.19a.6.6 0 0 0 .19-.06a.56.56 0 0 0 .17-.09l.15-.12a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41-.02ZM12 15a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1Zm6.42-6.78a7 7 0 0 0-13.36 1.89A4 4 0 0 0 6 18h2a1 1 0 0 0 0-2H6a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 13a3 3 0 0 1-3 3h-1a1 1 0 0 0 0 2h1a5 5 0 0 0 1.42-9.78Z"/>`),
		g.Group(children),
	)
}