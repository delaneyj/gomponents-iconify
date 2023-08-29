package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisabledTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 6a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-6 5a5 5 0 1 0 3.95 7.95"/><path d="m19 20l-4-5h-4l3-5l-4-3l-4 1"/></g>`),
		g.Group(children),
	)
}