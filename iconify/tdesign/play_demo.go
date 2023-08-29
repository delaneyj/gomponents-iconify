package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayDemo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v2h-1v14h-6.586l4 4L17 23.414l-5-5l-5 5L5.586 22l4-4H3V4H2V2Zm3 2v12h14V4H5Zm5 2.5l4.667 3.5L10 13.5v-7Z"/>`),
		g.Group(children),
	)
}