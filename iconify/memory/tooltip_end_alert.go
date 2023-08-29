package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipEndAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 15h-2v-2h2v2m0-3h-2V7h2v5m7-10v18h-1v1H6v-1H5v-5H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V2h1V1h14v1h1m-2 1H7v5H6v1H5v1H4v2h1v1h1v1h1v5h12V3Z"/>`),
		g.Group(children),
	)
}