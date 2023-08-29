package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 16h480v32H16zm139.883 179.883l22.627 22.627L240 157.02V456h32V157.02l61.49 61.49l22.627-22.627L256 95.764L155.883 195.883z"/>`),
		g.Group(children),
	)
}