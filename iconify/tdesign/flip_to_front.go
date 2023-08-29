package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipToFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3h14v14H7V3Zm2 2v10h10V5H9ZM5 7.5v3H3v-3h2Zm0 5v3H3v-3h2Zm0 5V19h1.5v2H3v-3.5h2Zm6.5 1.5v2h-3v-2h3Zm2 0h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}