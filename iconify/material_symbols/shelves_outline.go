package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShelvesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23V1h2v2h14V1h2v22h-2v-2H5v2H3Zm2-12h2V7h6v4h6V5H5v6Zm0 8h6v-4h6v4h2v-6H5v6Zm4-8h2V9H9v2Zm4 8h2v-2h-2v2Zm-4-8h2h-2Zm4 8h2h-2Z"/>`),
		g.Group(children),
	)
}