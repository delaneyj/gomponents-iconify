package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Church(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 1v2H9v2h2v.98l-5 3.5V12H2v10h20V12h-4V9.48l-5-3.5V5h2V3h-2V1h-2Zm1 6.72l4 2.8V20h-1v-4a3 3 0 1 0-6 0v4H8v-9.48l4-2.8ZM6 20H4v-6h2v6Zm5 0v-4a1 1 0 1 1 2 0v4h-2Zm7 0v-6h2v6h-2Z"/>`),
		g.Group(children),
	)
}