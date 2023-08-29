package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LegendToggle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-2h16v2H4Zm0-4v-2h16v2H4Zm0-4V8.65L10 5l5 3.55L20 5v2.45L15 11L9.925 7.4L4 11Z"/>`),
		g.Group(children),
	)
}