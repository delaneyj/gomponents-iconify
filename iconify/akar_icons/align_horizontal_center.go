package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M12 3v18"/><path stroke-linejoin="round" d="M16 6h4v4h-4m-8 0H4V6h4m8 8h2v4h-2m-8-4H6v4h2"/></g>`),
		g.Group(children),
	)
}