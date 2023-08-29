package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRampRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 3v8.707M11 7L7 3L3 7m14 7l4-4l-4-4"/><path d="M7 21a11 11 0 0 1 11-11h3"/></g>`),
		g.Group(children),
	)
}