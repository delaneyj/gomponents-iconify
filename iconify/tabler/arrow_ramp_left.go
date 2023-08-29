package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRampLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 3v8.707M13 7l4-4l4 4M7 14l-4-4l4-4"/><path d="M17 21A11 11 0 0 0 6 10H3"/></g>`),
		g.Group(children),
	)
}