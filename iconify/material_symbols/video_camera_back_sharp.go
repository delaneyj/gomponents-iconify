package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraBackSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h16v6.5l4-4v11l-4-4V20H2Zm3-4h10l-3.45-4.5l-2.3 3l-1.55-2L5 16Z"/>`),
		g.Group(children),
	)
}