package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Straighten(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V8q0-.825.588-1.413T4 6h3v6h2V6h2v6h2V6h2v6h2V6h3q.825 0 1.413.588T22 8v8q0 .825-.588 1.413T20 18H4Z"/>`),
		g.Group(children),
	)
}