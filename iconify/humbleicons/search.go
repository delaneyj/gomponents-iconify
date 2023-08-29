package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Search(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="m20 20l-6-6"/><path d="M15 9.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Z"/></g>`),
		g.Group(children),
	)
}