package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7h6v2H7v2h3v2H7v4H5V7m10 10h-2l4-8h-4V7h6v2l-4 8Z"/>`),
		g.Group(children),
	)
}