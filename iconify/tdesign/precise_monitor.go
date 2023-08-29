package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreciseMonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 2v6h-2V2h2Zm-9 .586L9.914 8.5L8.5 9.914L2.586 4L4 2.586ZM21.414 4L15.5 9.914L14.086 8.5L20 2.586L21.414 4Zm-11.146 7A2 2 0 0 1 14 12a2 2 0 0 1-3.732 1H2v-2h8.268ZM16 11h6v2h-6v-2Zm-6.086 4.5L4 21.414L2.586 20L8.5 14.086L9.914 15.5Zm5.586-1.414L21.414 20L20 21.414L14.086 15.5l1.414-1.414ZM13 16v6h-2v-6h2Z"/>`),
		g.Group(children),
	)
}