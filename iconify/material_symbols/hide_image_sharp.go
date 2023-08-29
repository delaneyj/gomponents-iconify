package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HideImageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18.15L5.85 3H21v15.15Zm-1.2 4.45L18.2 21H3V5.8L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM6 17h8.175l-2.1-2.1l-.825 1.1L9 13l-3 4Z"/>`),
		g.Group(children),
	)
}