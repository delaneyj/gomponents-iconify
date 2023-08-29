package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm15.586 6L15 3.414V7h3.586ZM13 3H5v18h14V9h-6V3Zm-6 9h10v2H7v-2Zm0 4h10v2H7v-2Z"/>`),
		g.Group(children),
	)
}