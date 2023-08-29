package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RangeHood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.05 12H21l-4-4V3H7v5l-3.95 4ZM4 20h16q.825 0 1.413-.588T22 18v-4H2v4q0 .825.588 1.413T4 20Zm6-3.3v-1.5h4v1.5h-4Z"/>`),
		g.Group(children),
	)
}