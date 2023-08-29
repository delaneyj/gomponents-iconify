package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuntingGear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M44 29H4v13h40V29ZM4 29L9.038 4.999H39.02l4.98 24"/><path stroke-linecap="round" d="M19 12a5 5 0 0 0 0 10m10 0a5 5 0 0 0 0-10m-9 5h8"/></g>`),
		g.Group(children),
	)
}