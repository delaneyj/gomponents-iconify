package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pentagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.636 5.636a9 9 0 1 1 12.728 12.728A9 9 0 0 1 5.636 5.636z"/><path d="m15.236 11l5.264 4H14l-2 6l-2-6H3.5l5.276-4L6.72 4.72L12 8.5l5.28-3.78z"/></g>`),
		g.Group(children),
	)
}