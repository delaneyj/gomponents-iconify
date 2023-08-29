package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatisticDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-miterlimit="5.759" d="M3 3v16a2 2 0 0 0 2 2h16"/><path stroke-miterlimit="5.759" d="m7 9l4 4l4-4l6 6"/><path d="M18 15h3v-3"/></g>`),
		g.Group(children),
	)
}