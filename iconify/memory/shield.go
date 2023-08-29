package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 4h2V3h2V2h2V1h4v1h2v1h2v1h2v10h-1v2h-1v2h-1v1h-1v1h-2v1H9v-1H7v-1H6v-1H5v-2H4v-2H3V4m7-1v1H8v1H6v1H5v7h1v2h1v2h1v1h2v1h2v-1h2v-1h1v-2h1v-2h1V6h-1V5h-2V4h-2V3h-2Z"/>`),
		g.Group(children),
	)
}