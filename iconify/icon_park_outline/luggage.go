package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="32" height="26" x="8" y="14" stroke-linejoin="round" rx="2"/><path d="M20 23v8"/><path stroke-linejoin="round" d="M15 40v4m18-4v4"/><path d="M28 23v8"/><path stroke-linejoin="round" d="M19 4h10m-5 0v10"/></g>`),
		g.Group(children),
	)
}