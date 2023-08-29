package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterEightSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18V2h16v16H6Zm-4 4V6h2v14h14v2H2Zm10-12l-1 1v4h6v-4l-1-1l1-1V5h-6v4l1 1Zm3-3v2h-2V7h2Zm0 4v2h-2v-2h2Z"/>`),
		g.Group(children),
	)
}