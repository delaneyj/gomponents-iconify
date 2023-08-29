package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 7h6v2H8v2h3v2H8v4H6V7m8 0h4v10h-2V9h-2V7Z"/>`),
		g.Group(children),
	)
}