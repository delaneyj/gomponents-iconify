package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExportThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 112v96a12 12 0 0 1-12 12H56a12 12 0 0 1-12-12v-96a12 12 0 0 1 12-12h24a4 4 0 0 1 0 8H56a4 4 0 0 0-4 4v96a4 4 0 0 0 4 4h144a4 4 0 0 0 4-4v-96a4 4 0 0 0-4-4h-24a4 4 0 0 1 0-8h24a12 12 0 0 1 12 12ZM90.83 66.83L124 33.66V136a4 4 0 0 0 8 0V33.66l33.17 33.17a4 4 0 1 0 5.66-5.66l-40-40a4 4 0 0 0-5.66 0l-40 40a4 4 0 0 0 5.66 5.66Z"/>`),
		g.Group(children),
	)
}