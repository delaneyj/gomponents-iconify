package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectAddressOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 24c0 9.941 8.059 18 18 18s18-8.059 18-18M24 14v28"/><circle cx="24" cy="10" r="4"/></g>`),
		g.Group(children),
	)
}