package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.325 18.475L20 17.15V5H7.85l-2-2H22v15.475h-.675Zm-18.15-15.3L5 5H4v12h10.15L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4l-4.3-4.3H16v2H8v-2H2V3.175h1.175ZM9.1 11.95Zm4.875-.825Z"/>`),
		g.Group(children),
	)
}