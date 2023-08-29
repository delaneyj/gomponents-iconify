package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M41 44V20H7v24h34Zm-17 0V20m17 24H7"/><path d="M4 12h40v8H4z"/><path stroke-linecap="round" d="m16 4l8 8l8-8"/></g>`),
		g.Group(children),
	)
}