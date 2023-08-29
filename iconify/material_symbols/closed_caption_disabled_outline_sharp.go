package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptionDisabledOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.025 9l1.5 1.5H7.5v3h2V13H11v2H6V9h3.025Zm-2.15-5H21v14.125l-2-2V6H8.875l-2-2ZM18 13v2h-.125l-1.5-1.5h.125V13H18Zm-1.5-2v-.5h-2v1.125l-1.5-1.5V9h5v2h-1.5Zm-2.55.05Zm-3.85 1.875V12.9v.025Zm-5.9-8.75L6.025 6H5v12h10.175L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425l-3.3-3.3H3V4.175h1.2Z"/>`),
		g.Group(children),
	)
}