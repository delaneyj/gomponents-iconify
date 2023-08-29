package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignmentLeftTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" stroke-linejoin="round" rx="3"/><path stroke="#fff" d="M12 24H16"/><path stroke="#fff" d="M12 18H24"/><path stroke="#fff" d="M12 12H20"/></g>`),
		g.Group(children),
	)
}