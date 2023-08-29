package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><ellipse cx="24" cy="39" stroke-linejoin="round" rx="18" ry="6"/><path stroke-linecap="round" d="M6 39h36"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 39L24 4l18 35"/></g>`),
		g.Group(children),
	)
}