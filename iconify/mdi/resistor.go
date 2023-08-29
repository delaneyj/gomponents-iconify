package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Resistor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11h5l3.07 4.35L13.11 4L18 11h4v2h-5l-3.07-4.35L10.89 20L6 13H2v-2Z"/>`),
		g.Group(children),
	)
}