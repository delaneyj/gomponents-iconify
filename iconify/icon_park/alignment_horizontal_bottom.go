package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignmentHorizontalBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" stroke-linejoin="round" rx="3"/><path stroke="#fff" d="M22 36H26"/><path stroke="#fff" d="M18 30H30"/><path stroke="#fff" d="M20 24H28"/></g>`),
		g.Group(children),
	)
}