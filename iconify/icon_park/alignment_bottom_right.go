package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignmentBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" stroke-linejoin="round" rx="3"/><path stroke="#fff" d="M24 32V36"/><path stroke="#fff" d="M30 24V36"/><path stroke="#fff" d="M36 28V36"/></g>`),
		g.Group(children),
	)
}