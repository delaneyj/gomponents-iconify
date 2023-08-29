package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarRepairSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-3H4v-2h16v2h-7v3h-2ZM9 11.5q.425 0 .713-.288T10 10.5q0-.425-.288-.713T9 9.5q-.425 0-.713.288T8 10.5q0 .425.288.713T9 11.5Zm6 0q.425 0 .713-.288T16 10.5q0-.425-.288-.713T15 9.5q-.425 0-.713.288T14 10.5q0 .425.288.713T15 11.5ZM5 8.6L6.925 3h10.15L19 8.6V16h-2v-2H7v2H5V8.6ZM7.65 7h8.7l-.7-2h-7.3l-.7 2Z"/>`),
		g.Group(children),
	)
}