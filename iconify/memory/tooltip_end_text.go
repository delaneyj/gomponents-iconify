package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipEndText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 8V6H9v2h8m-2 8v-2H9v2h6m2-4v-2H9v2h8m4 8V2h-1V1H6v1H5v5H4v1H3v1H2v1H1v2h1v1h1v1h1v1h1v5h1v1h14v-1h1m-2-1H7v-5H6v-1H5v-1H4v-2h1V9h1V8h1V3h12v16Z"/>`),
		g.Group(children),
	)
}