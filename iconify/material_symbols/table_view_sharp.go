package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableViewSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21V7h14v14H7Zm2-10h10V9H9v2Zm4 4h2v-2h-2v2Zm0 4h2v-2h-2v2Zm-4-4h2v-2H9v2Zm8 0h2v-2h-2v2Zm-8 4h2v-2H9v2Zm8 0h2v-2h-2v2ZM6 17H3V3h14v3h-2V5H5v10h1v2Z"/>`),
		g.Group(children),
	)
}