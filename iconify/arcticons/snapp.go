package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="20.419" cy="38.403" r="5.097"/><path d="M32.678 4.5L28.04 22.089c-1.111 4.217-1.8 7.656-6.212 7.698l-4.314.041L24.3 4.508l8.378-.008Z"/></g>`),
		g.Group(children),
	)
}