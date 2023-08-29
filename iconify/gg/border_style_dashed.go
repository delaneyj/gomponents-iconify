package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderStyleDashed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 11h4v2H4v-2Zm6 0h4v2h-4v-2Zm10 0h-4v2h4v-2Z"/>`),
		g.Group(children),
	)
}