package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalBottomRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21q0-.425.288-.713T3 20h18q.425 0 .713.288T22 21q0 .425-.288.713T21 22H3Zm5.5-4q-.625 0-1.063-.438T7 16.5v-13q0-.625.438-1.063T8.5 2q.625 0 1.063.438T10 3.5v13q0 .625-.438 1.063T8.5 18Zm7 0q-.625 0-1.063-.438T14 16.5v-7q0-.625.438-1.063T15.5 8q.625 0 1.063.438T17 9.5v7q0 .625-.438 1.063T15.5 18Z"/>`),
		g.Group(children),
	)
}