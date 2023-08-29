package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.4 5H18v10h-7l-.4-2H5v7H3V3h9l.4 2M14 17h8v2h-8v-2Z"/>`),
		g.Group(children),
	)
}