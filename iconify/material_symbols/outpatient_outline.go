package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutpatientOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V3h14v18H9v-4H7v4H1Zm2-2h2v-4h6v4h2V5H3v14Zm2-6h2v-2H5v2Zm0-4h2V7H5v2Zm4 4h2v-2H9v2Zm0-4h2V7H9v2Zm10.5 6.5l-1.4-1.4l1.075-1.1H16v-2h3.175L18.1 9.9l1.4-1.4L23 12l-3.5 3.5ZM5 19v-4h6v4v-4H5v4Z"/>`),
		g.Group(children),
	)
}