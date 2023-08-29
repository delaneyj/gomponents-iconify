package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightUpDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 19h6a3 3 0 0 0 3-3V9"/><path d="m10 13l4-4l4 4m-8-5l4-4l4 4"/></g>`),
		g.Group(children),
	)
}