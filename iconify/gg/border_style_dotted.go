package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderStyleDotted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11H1v2h2v-2Zm4 0H5v2h2v-2Zm2 0h2v2H9v-2Zm6 0h-2v2h2v-2Zm2 0h2v2h-2v-2Zm6 0h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}