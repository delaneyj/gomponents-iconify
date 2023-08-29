package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BracketsCurly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 6a2 2 0 0 1 2-2a1 1 0 0 0 0-2a4 4 0 0 0-4 4v3a2 2 0 0 1-2 2a1 1 0 0 0 0 2a2 2 0 0 1 2 2v3a4 4 0 0 0 4 4a1 1 0 0 0 0-2a2 2 0 0 1-2-2v-3a4 4 0 0 0-1.38-3A4 4 0 0 0 6 9Zm16 5a2 2 0 0 1-2-2V6a4 4 0 0 0-4-4a1 1 0 0 0 0 2a2 2 0 0 1 2 2v3a4 4 0 0 0 1.38 3A4 4 0 0 0 18 15v3a2 2 0 0 1-2 2a1 1 0 0 0 0 2a4 4 0 0 0 4-4v-3a2 2 0 0 1 2-2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}