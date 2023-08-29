package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taurus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="31" r="9"/><path stroke-linecap="round" d="M44 8c0 7.732-8.954 14-20 14S4 15.732 4 8"/></g>`),
		g.Group(children),
	)
}