package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntersectSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="27" height="27" x="16" y="16" rx="2"/><rect width="27" height="27" x="5" y="5" rx="2"/><path d="M27 16L16 27m16-6L21 32"/></g>`),
		g.Group(children),
	)
}