package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterFourOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18V2h16v16H6Zm2-2h12V4H8v12Zm-6 6V6h2v14h14v2H2Zm6-6V4v12Zm7-1h2V5h-2v4h-2V5h-2v6h4v4Z"/>`),
		g.Group(children),
	)
}