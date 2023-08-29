package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Resize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.59 12l4-4H11V6h7v7h-2V9.41l-4 4V16h8V4H8v8h2.59M22 2v16H12v4H2V12h4V2h16M10 14H4v6h6v-6Z"/>`),
		g.Group(children),
	)
}