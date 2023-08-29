package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h6v2H5v2h3v2H5v4H3V7m8 0h4v10h-2V9h-2V7m6 0h4v10h-2V9h-2V7Z"/>`),
		g.Group(children),
	)
}