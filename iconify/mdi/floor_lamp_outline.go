package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorLampOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 2l2 7H7l2-7m4.6 2h-3.2l-.85 3h4.9M11 10h2v10h3v2H8v-2h3Z"/>`),
		g.Group(children),
	)
}