package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M19 4h-6.5a3 3 0 0 0-3 3v12"/><path d="M6 16.5L9.5 20l3.5-3.5"/></g>`),
		g.Group(children),
	)
}