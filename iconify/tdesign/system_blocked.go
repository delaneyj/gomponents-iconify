package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemBlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v9.5h-2V3H3v13h8.5v2H1V1Zm17.5 13a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.483 3.483 0 0 0 18.5 14Zm3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745ZM13 17.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0ZM2.25 21h9.25v2H2.25v-2Z"/>`),
		g.Group(children),
	)
}