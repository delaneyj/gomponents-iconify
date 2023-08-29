package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M21 2H3"/><path stroke-linejoin="round" d="M6 22V6h4v16H6Zm8-6V6h4v10h-4Z"/></g>`),
		g.Group(children),
	)
}