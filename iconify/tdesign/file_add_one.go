package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAddOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm15.586 6L15 3.414V7h3.586ZM13 3H5v18h14V9h-6V3Zm0 8v3h3v2h-3v3h-2v-3H8v-2h3v-3h2Z"/>`),
		g.Group(children),
	)
}