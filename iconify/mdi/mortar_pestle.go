package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MortarPestle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 7l-2 6l2 6v2H3v-2l2-6l-2-6V5h12.7l1.5-4l2.3.8L18.3 5H21v2Z"/>`),
		g.Group(children),
	)
}