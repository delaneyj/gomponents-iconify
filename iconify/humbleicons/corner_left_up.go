package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M19 20h-6.5a3 3 0 0 1-3-3V5"/><path d="M6 7.5L9.5 4L13 7.5"/></g>`),
		g.Group(children),
	)
}