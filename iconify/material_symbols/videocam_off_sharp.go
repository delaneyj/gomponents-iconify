package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocamOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.55 23.35L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4ZM22 17.5l-4-4v1.675L6.825 4H18v6.5l4-4v11ZM4 4l14 14v2H2V4h2Z"/>`),
		g.Group(children),
	)
}