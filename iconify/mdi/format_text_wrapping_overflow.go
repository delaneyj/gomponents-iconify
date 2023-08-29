package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextWrappingOverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21H5V3h2v18m7-18h-2v6h2V3m0 12h-2v6h2v-6m5-3l-3-3v2H9v2h7v2l3-3Z"/>`),
		g.Group(children),
	)
}