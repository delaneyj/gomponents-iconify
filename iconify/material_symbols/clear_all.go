package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17v-2h14v2H3Zm2-4v-2h14v2H5Zm2-4V7h14v2H7Z"/>`),
		g.Group(children),
	)
}