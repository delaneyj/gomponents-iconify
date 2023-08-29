package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1 1.754l11 9.9l11-9.9V21h-2v-7h-2v7h-2v-7H7v7H5v-7H3v7H1V1.754ZM3 12h6.394L3 6.245V12Zm11.606 0h6.393L21 6.245L14.606 12Z"/>`),
		g.Group(children),
	)
}