package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shelter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13 2L1 6v2l1-.333V13h10v-2H4V7l9-3V2Z"/>`),
		g.Group(children),
	)
}