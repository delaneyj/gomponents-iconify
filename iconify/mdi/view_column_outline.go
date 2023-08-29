package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumnOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5v13h17V5H4m10 2v9h-3V7h3M6 7h3v9H6V7m13 9h-3V7h3v9Z"/>`),
		g.Group(children),
	)
}