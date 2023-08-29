package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastUpButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m17 41l9.111-18.795L33.78 6.388c.897-1.85 3.545-1.85 4.442 0l7.668 15.817L55 41"/><path d="m17 67l9.111-18.882l7.668-15.891c.897-1.86 3.545-1.86 4.442 0l7.668 15.89L55 67"/></g>`),
		g.Group(children),
	)
}