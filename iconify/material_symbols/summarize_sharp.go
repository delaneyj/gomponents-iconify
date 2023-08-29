package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SummarizeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9h2V7H7v2Zm0 4h2v-2H7v2Zm0 4h2v-2H7v2Zm-4 4V3h13l5 5v13H3ZM15 5v4h4l-4-4Z"/>`),
		g.Group(children),
	)
}