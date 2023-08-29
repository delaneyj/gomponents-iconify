package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.22 7.091A7.1 7.1 0 0 0 14.127 0H2.181a2.182 2.182 0 0 0 0 4.364h.927v15.272h-.927a2.182 2.182 0 0 0 0 4.364H14.13a7.087 7.087 0 0 0 5.106-12.002l.002.002a7.04 7.04 0 0 0 1.981-4.906v-.003zm-7.09 2.727H7.473V4.363h6.655a2.727 2.727 0 0 1 0 5.454zm0 9.818H7.473v-5.454h6.655a2.727 2.727 0 0 1 0 5.454z"/>`),
		g.Group(children),
	)
}