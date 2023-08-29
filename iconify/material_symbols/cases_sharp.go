package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CasesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 22V9h2v11h17v2H1Zm4-4V5h5V1h8v4h5v13H5Zm7-13h4V3h-4v2Z"/>`),
		g.Group(children),
	)
}