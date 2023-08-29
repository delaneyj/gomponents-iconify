package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.95 14.395l1.414-1.415L12 6.617L5.636 12.98l1.414 1.415L12 9.445l4.95 4.95ZM6 17.384h12v-2H6v2Z"/>`),
		g.Group(children),
	)
}