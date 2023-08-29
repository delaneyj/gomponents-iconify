package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAttachment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V12h-2V9h-6V3H5v18h7v2H3V1Zm12 2.414V7h3.586L15 3.414ZM14 15.5a2.5 2.5 0 0 1 5 0V20h-2v-4.5a.5.5 0 0 0-1 0V20a2 2 0 1 0 4 0v-4h2v4a4 4 0 0 1-8 0v-4.5Z"/>`),
		g.Group(children),
	)
}