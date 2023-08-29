package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RequestPageOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 18h2v-1h2v-5h-4v-1h4V9h-2V8h-2v1H9v5h4v1H9v2h2v1Zm-7 4V2h10l6 6v14H4Zm2-2h12V8.85L13.15 4H6v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}