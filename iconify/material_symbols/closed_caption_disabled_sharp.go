package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptionDisabledSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20V5.825L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425l-3.3-3.3H3Zm18-1.875L17.875 15H18v-2h-1.5v.5h-.125L14.5 11.625V10.5h2v.5H18V9h-5v1.125L6.875 4H21v14.125ZM6 15h5v-1.175L10.175 13H9.5v.5h-2v-3.175L6.375 9.2H6V15Z"/>`),
		g.Group(children),
	)
}