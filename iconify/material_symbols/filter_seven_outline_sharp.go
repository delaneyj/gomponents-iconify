package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterSevenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15h2l4-8V5h-6v2h4l-4 8Zm-5 3V2h16v16H6Zm2-2h12V4H8v12Zm-6 6V6h2v14h14v2H2Zm6-6V4v12Z"/>`),
		g.Group(children),
	)
}