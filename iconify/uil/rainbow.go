package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 12a1 1 0 0 0 0 2a5 5 0 0 1 5 5a1 1 0 0 0 2 0a7 7 0 0 0-7-7Zm0-8a1 1 0 0 0 0 2a13 13 0 0 1 13 13a1 1 0 0 0 2 0A15 15 0 0 0 5 4Zm0 4a1 1 0 0 0 0 2a9 9 0 0 1 9 9a1 1 0 0 0 2 0A11 11 0 0 0 5 8Z"/>`),
		g.Group(children),
	)
}