package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableFurniture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7h20v3h-2l1 9h-2.5l-.56-5H6.06l-.56 5H3l1-9H2V7m15.5 3h-11l-.21 2h11.42l-.21-2Z"/>`),
		g.Group(children),
	)
}