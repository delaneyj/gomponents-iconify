package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrScanTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 3h6v5h-2V5h-4V3ZM9 3v2H5v3H3V3h6Zm6 18v-2h4v-3h2v5h-6Zm-6 0H3v-5h2v3h4v2ZM3 11h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}