package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FutureBuildTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 44h40"/><ellipse cx="24.5" cy="7" rx="13.5" ry="3"/><path d="M16 9s4.16 8.883 5 15c1.069 7.776-1 20-1 20M32.227 9s-4.16 8.883-5 15C26.157 31.776 28 44 28 44"/></g>`),
		g.Group(children),
	)
}