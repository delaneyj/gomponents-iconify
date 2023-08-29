package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 5H4V3L0 6.5L4 10V8h10V5zm6 8.5L16 10v2H6v3h10v2l4-3.5z"/>`),
		g.Group(children),
	)
}