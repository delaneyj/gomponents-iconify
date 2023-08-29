package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 3.086l7.414 7.414H2v-2h15.586l-4-4L15 3.086ZM22 13.5v2H6.414l4 4L9 20.914L1.586 13.5H22Z"/>`),
		g.Group(children),
	)
}