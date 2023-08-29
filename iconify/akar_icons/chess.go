package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.533 3.81L8 2l1 4l-5.37 4.475A1.75 1.75 0 0 0 3 11.82v0c0 .617.537 1.088 1.127.986L9 12l-2.097 7h10.614l1.283-5.745c.913-4.088-1.386-8.21-5.267-9.445ZM4 21a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v1H4v-1Z"/>`),
		g.Group(children),
	)
}