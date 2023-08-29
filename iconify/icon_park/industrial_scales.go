package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndustrialScales(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M10 32H38L42 40H6L10 32Z"/><path stroke-linecap="round" d="M16 40V44"/><path stroke-linecap="round" d="M24 12V32"/><rect width="14" height="8" x="17" y="4" fill="#2F88FF"/><path stroke-linecap="round" d="M32 40V44"/></g>`),
		g.Group(children),
	)
}