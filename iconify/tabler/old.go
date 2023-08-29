package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Old(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m11 21l-1-4l-2-3V8"/><path d="m5 14l-1-3l4-3l3 2l3 .5M7 4a1 1 0 1 0 2 0a1 1 0 1 0-2 0m0 13l-2 4m11 0v-8.5a1.5 1.5 0 0 1 3 0v.5"/></g>`),
		g.Group(children),
	)
}