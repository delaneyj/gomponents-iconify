package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="30" x="4" y="6" rx="3"/><path d="M24 36v7m8-29L16 28m-6 15h28"/><circle cx="15" cy="17" r="3"/><circle cx="33" cy="25" r="3"/></g>`),
		g.Group(children),
	)
}