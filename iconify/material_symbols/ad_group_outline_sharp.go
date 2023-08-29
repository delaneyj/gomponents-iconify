package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdGroupOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16h12V6H8v10Zm-2 2V2h16v16H6Zm-4 4V6h2v14h14v2H2Zm6-6V4v12Z"/>`),
		g.Group(children),
	)
}