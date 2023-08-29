package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M356.117 316.117L333.49 293.49L272 354.98V56h-32v298.98l-61.49-61.49l-22.627 22.627L256 416.236l100.117-100.119zM16 464h480v32H16z"/>`),
		g.Group(children),
	)
}