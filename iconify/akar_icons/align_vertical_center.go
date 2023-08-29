package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M21 12H3"/><path stroke-linejoin="round" d="M6 16v4h4v-4m4 0v2h4v-2m-4-8V6h4v2m-8 0V4H6v4"/></g>`),
		g.Group(children),
	)
}