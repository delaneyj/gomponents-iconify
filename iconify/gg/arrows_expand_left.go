package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.1 4.1v-2h-8v8h2V5.516l5.779 5.778l1.414-1.414L5.515 4.1H10.1Zm9.8 9.8h2v8h-8v-2h4.585l-5.778-5.779l1.414-1.414l5.778 5.778V13.9Z"/>`),
		g.Group(children),
	)
}