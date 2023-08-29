package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectInverse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3h2v2h2V3h2v2h2V3h2v2h2V3h2v2h2v2h-2v2h2v2h-2v2h2v2h-2v2h2v2h-2v2h-2v-2h-2v2h-2v-2h-2v2H9v-2H7v2H5v-2H3v-2h2v-2H3v-2h2v-2H3V9h2V7H3V5h2V3Z"/>`),
		g.Group(children),
	)
}