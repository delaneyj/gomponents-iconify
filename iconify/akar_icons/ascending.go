package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ascending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3v18M6 3l3 3.333M6 3L3 6.333"/><path stroke-miterlimit="5.759" d="M14 21h8l-1-4h-7v4Zm0-7h6l-1-4h-5v4Zm0-7h4l-1-4h-3v4Z"/></g>`),
		g.Group(children),
	)
}