package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 32V16m18 11v-6M6 27v-6m8-15H8a2 2 0 0 0-2 2v6m28-8h6a2 2 0 0 1 2 2v6m-8 28h6a2 2 0 0 0 2-2v-6m-28 8H8a2 2 0 0 1-2-2v-6M27 6h-6m11 18H16m11 18h-6"/>`),
		g.Group(children),
	)
}