package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VignetteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm10-4q2.45 0 4.225-1.188T18 12q0-1.625-1.775-2.813T12 8Q9.55 8 7.775 9.188T6 12q0 1.625 1.775 2.813T12 16Z"/>`),
		g.Group(children),
	)
}