package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrScanLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16v5H3v-5h2v3h14v-3h2ZM3 11h18v2H3v-2Zm18-3h-2V5H5v3H3V3h18v5Z"/>`),
		g.Group(children),
	)
}