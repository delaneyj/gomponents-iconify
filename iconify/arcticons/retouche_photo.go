package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RetouchePhoto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M39.047 5.36a2.479 2.479 0 1 1-.054-.06M15.121 31.478V43.5l-6.77-3.856V18.161"/><path d="M8.357 18.378a14.267 14.267 0 1 1 6.817 13.063"/><path d="M28.354 14.294a7.602 7.602 0 1 1-.167-.187"/></g>`),
		g.Group(children),
	)
}