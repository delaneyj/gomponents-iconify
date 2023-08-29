package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M21 22H3"/><path stroke-linejoin="round" d="M6 18V2h4v16H6Zm8 0V8h4v10h-4Z"/></g>`),
		g.Group(children),
	)
}