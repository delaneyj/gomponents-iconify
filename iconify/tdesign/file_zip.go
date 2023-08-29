package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileZip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm2 2v18h14V9h-6V3h-1.996v2.004h-2V7h2v2.004h-2V11h2v2.004h-2v2H7V13h2v-1.996H7V9h2V7.004H7V5h2V3H5Zm10 .414V7h3.586L15 3.414Z"/>`),
		g.Group(children),
	)
}