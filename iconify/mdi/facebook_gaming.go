package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookGaming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 14.5v-5H21V21h-5.5v-6.5h-6M3 3h18v5.5H8.5v7h6V21H3V3Z"/>`),
		g.Group(children),
	)
}