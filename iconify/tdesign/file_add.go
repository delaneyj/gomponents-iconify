package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V12h-2V9h-6V3H5v18h7v2H3V1Zm12 2.414V7h3.586L15 3.414ZM20 14v4h4v2h-4v4h-2v-4h-4v-2h4v-4h2Z"/>`),
		g.Group(children),
	)
}