package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="40" height="28" x="4" y="8" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><rect width="16" height="12" x="12" y="16" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="1"/><circle cx="37" cy="15" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 22h2m-2 7h2m-26 7v6m24-6v6"/></g>`),
		g.Group(children),
	)
}