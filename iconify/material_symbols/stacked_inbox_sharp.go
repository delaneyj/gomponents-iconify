package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedInboxSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17V3h18v14H5Zm9-5q.825 0 1.413-.588T16 10h5V5H7v5h5q0 .825.588 1.413T14 12Zm5 9H1V7h2v12h16v2Z"/>`),
		g.Group(children),
	)
}