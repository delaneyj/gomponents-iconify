package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsLineDashed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.42 9.42L8 12"/><circle cx="4" cy="8" r="2"/><path d="m14 6l-8.58 8.58"/><circle cx="4" cy="16" r="2"/><path d="M10.8 14.8L14 18m2-6h-2m8 0h-2"/></g>`),
		g.Group(children),
	)
}