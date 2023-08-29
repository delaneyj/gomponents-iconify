package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yoga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 4a1 1 0 1 0 2 0a1 1 0 1 0-2 0M4 20h4l1.5-3m7.5 3l-1-5h-5l1-7"/><path d="m4 10l4-1l4-1l4 1.5l4 1.5"/></g>`),
		g.Group(children),
	)
}