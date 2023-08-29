package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VibrationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 15V9h2v6H0Zm3 2V7h2v10H3Zm19-2V9h2v6h-2Zm-3 2V7h2v10h-2ZM6 21V3h12v18H6Z"/>`),
		g.Group(children),
	)
}