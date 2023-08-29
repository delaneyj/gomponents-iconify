package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MangaPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.965 41.862l8.413 2.362l-2.575-8.171M9.5 24h29M24 9.5v29m11.965 3.362a21.5 21.5 0 1 1 5.838-5.81"/>`),
		g.Group(children),
	)
}