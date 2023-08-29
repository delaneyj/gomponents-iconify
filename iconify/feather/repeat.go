package feather

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 1l4 4l-4 4"/><path d="M3 11V9a4 4 0 0 1 4-4h14M7 23l-4-4l4-4"/><path d="M21 13v2a4 4 0 0 1-4 4H3"/></g>`),
		g.Group(children),
	)
}