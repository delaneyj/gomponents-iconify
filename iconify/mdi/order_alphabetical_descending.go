package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrderAlphabeticalDescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 13H5c-1.1 0-2 .9-2 2v6h2v-2h2v2h2v-6a2 2 0 0 0-2-2m0 4H5v-2h2M9 3v2L5.67 9H9v2H3V9l3.33-4H3V3m9 2h10v2H12m0 12v-2h10v2m-10-8h10v2H12Z"/>`),
		g.Group(children),
	)
}