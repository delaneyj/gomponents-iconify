package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Institution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .856l10 5.556V9H2V6.412L12 .856ZM5.06 7h13.88L12 3.144L5.06 7ZM7 11v8H5v-8h2Zm6 0v8h-2v-8h2Zm6 0v8h-2v-8h2ZM2 21h20v2H2v-2Z"/>`),
		g.Group(children),
	)
}