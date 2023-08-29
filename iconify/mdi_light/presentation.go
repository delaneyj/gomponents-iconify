package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Presentation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h8a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1h8v1h-1v11h-5.73l1.61 6h-1.04l-1.61-6H9.77l-1.61 6H7.12l1.61-6H3V5H2V4m17 11V5H4v10h15Z"/>`),
		g.Group(children),
	)
}