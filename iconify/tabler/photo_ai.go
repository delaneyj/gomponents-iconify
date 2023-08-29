package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoAi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 8h.01M10 21H6a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v5"/><path d="m3 16l5-5c.928-.893 2.072-.893 3 0l1 1m2 9v-4a2 2 0 1 1 4 0v4m-4-2h4m3-4v6"/></g>`),
		g.Group(children),
	)
}