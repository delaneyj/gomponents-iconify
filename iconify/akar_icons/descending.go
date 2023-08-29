package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Descending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 21V3m0 18l3-3.333M6 21l-3-3.333"/><path stroke-miterlimit="5.759" d="M14 3h8l-1 4h-7V3Zm0 7h6l-1 4h-5v-4Zm0 7h4l-1 4h-3v-4Z"/></g>`),
		g.Group(children),
	)
}