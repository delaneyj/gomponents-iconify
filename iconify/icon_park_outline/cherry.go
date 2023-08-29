package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cherry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><circle cx="14" cy="34" r="8" stroke-linejoin="round"/><circle cx="35" cy="37" r="7" stroke-linejoin="round"/><path d="M37 10c-2.651.812-8.372 3.014-11.72 6.26C20.255 21.13 19 24.5 18 27m19-17c-1.117 1.318-3.285 4.596-3.956 8.39C32.035 24.078 33 27.5 34 30M30 4l14 12"/></g>`),
		g.Group(children),
	)
}