package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MangaOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm6.1-2l-1.075-1.5l1.9-2.625v-3.25l3.075-1L13.925 7L17 8l3-.975V6H4v12h4.1Zm2.825-5.925ZM10.575 18H20v-7.725l-.675-.925L17 10.1l-2.3-.75l-1.45 1.975l-2.325.75v2.45L9.5 16.5l1.075 1.5Z"/>`),
		g.Group(children),
	)
}