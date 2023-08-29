package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MindmapList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M26 24L44 24"/><path d="M14 24L18 24"/><path d="M18 38H44"/><path d="M6 38H10"/><path d="M18 10H44"/><path d="M6 10H10"/></g>`),
		g.Group(children),
	)
}