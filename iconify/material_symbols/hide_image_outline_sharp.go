package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HideImageOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 18.15l-2-2V5H7.85l-2-2H21v15.15Zm-1.2 4.45L18.2 21H3V5.8L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM6 17l3-4l2.25 3l.825-1.1L5 7.825V19h11.175l-2-2H6Zm7.425-6.425ZM10.6 13.4Z"/>`),
		g.Group(children),
	)
}