package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="40" height="40" x="4" y="4" stroke-linejoin="round" rx="2"/><path d="M14 4V44"/><path d="M24 4V44"/><path d="M34 4V44"/><path stroke-linejoin="round" d="M4 14H44"/><path stroke-linejoin="round" d="M4 34H44"/><path stroke-linejoin="round" d="M4 24H44"/></g>`),
		g.Group(children),
	)
}