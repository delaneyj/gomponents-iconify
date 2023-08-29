package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalTopRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 22q-.625 0-1.063-.438T7 20.5v-13q0-.625.438-1.063T8.5 6q.625 0 1.063.438T10 7.5v13q0 .625-.438 1.063T8.5 22Zm7-6q-.625 0-1.063-.438T14 14.5v-7q0-.625.438-1.063T15.5 6q.625 0 1.063.438T17 7.5v7q0 .625-.438 1.063T15.5 16ZM3 4q-.425 0-.713-.288T2 3q0-.425.288-.713T3 2h18q.425 0 .713.288T22 3q0 .425-.288.713T21 4H3Z"/>`),
		g.Group(children),
	)
}