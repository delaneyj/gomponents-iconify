package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarthquakeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.275 22L5.5 13H2v-2h4.975l1.875 6.075L12.2 2h1.6l2.35 10.175L17.775 7H19.2l1.5 4H22v2h-2.7l-.725-1.925L16.725 17H15.2L13 7.525L9.8 22H8.275Z"/>`),
		g.Group(children),
	)
}