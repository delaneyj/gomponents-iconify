package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shelves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23V1h2v2h14V1h2v22h-2v-2H5v2H3Zm2-12h2V7h6v4h6V5H5v6Zm0 8h6v-4h6v4h2v-6H5v6Z"/>`),
		g.Group(children),
	)
}