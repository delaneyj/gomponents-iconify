package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 10c6-6.667 14-6.667 20 0M6 14c3.6-4 8.4-4 12 0"/><circle cx="12" cy="18" r="1"/></g>`),
		g.Group(children),
	)
}