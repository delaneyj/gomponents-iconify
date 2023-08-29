package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntersectSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="27" height="27" x="16" y="16" rx="2"/><rect width="27" height="27" x="5" y="5" rx="2"/><path d="M27 16L16 27"/><path d="M32 21L21 32"/></g>`),
		g.Group(children),
	)
}