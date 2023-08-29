package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterNinePlusSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18V2h16v16H6Zm-4 4V6h2v14h14v2H2Zm8-8h4V6H9v5h3v1h-2v2Zm2-5h-1V8h1v1Zm4.5 4h2v-2h2V9h-2V7h-2v2h-2v2h2v2Z"/>`),
		g.Group(children),
	)
}