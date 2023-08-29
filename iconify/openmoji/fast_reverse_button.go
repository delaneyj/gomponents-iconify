package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastReverseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m41 55l-18.795-9.111L6.388 38.22c-1.85-.897-1.85-3.545 0-4.442l15.817-7.668L41 17"/><path d="m67 55l-18.882-9.111l-15.891-7.668c-1.86-.897-1.86-3.545 0-4.442l15.89-7.668L67 17"/></g>`),
		g.Group(children),
	)
}