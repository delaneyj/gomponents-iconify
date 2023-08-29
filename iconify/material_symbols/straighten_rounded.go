package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StraightenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V8q0-.825.588-1.413T4 6h3v5q0 .425.288.713T8 12q.425 0 .713-.288T9 11V6h2v5q0 .425.288.713T12 12q.425 0 .713-.288T13 11V6h2v5q0 .425.288.713T16 12q.425 0 .713-.288T17 11V6h3q.825 0 1.413.588T22 8v8q0 .825-.588 1.413T20 18H4Z"/>`),
		g.Group(children),
	)
}