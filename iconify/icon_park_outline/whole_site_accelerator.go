package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WholeSiteAccelerator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="22" cy="40" r="4"/><circle cx="26" cy="8" r="4"/><circle cx="36" cy="24" r="4"/><circle cx="12" cy="24" r="4"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 24H16m7-13l-8 10"/><path d="m33 27l-8.001 10"/></g>`),
		g.Group(children),
	)
}