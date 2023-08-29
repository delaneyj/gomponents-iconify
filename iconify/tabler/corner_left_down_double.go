package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftDownDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 4h-6a3 3 0 0 0-3 3v7"/><path d="m13 10l-4 4l-4-4m8 5l-4 4l-4-4"/></g>`),
		g.Group(children),
	)
}