package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="m15 12l3-6h12l3 6H15Z"/><rect width="40" height="30" x="4" y="12" rx="3"/><path d="M24 35a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/></g>`),
		g.Group(children),
	)
}