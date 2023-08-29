package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItalicK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 19L6 1h3L7 9.5L15 1h3l-8.5 9l5.5 9h-3l-5-9l-2 9H2Z"/>`),
		g.Group(children),
	)
}