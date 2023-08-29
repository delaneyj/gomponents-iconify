package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1V3Zm2 2v4h18V5H3Zm18 6h-2v5h-2v-5h-2v5h-2v-5h-2v5H9v-5H7v5H5v-5H3v8h18v-8Z"/>`),
		g.Group(children),
	)
}