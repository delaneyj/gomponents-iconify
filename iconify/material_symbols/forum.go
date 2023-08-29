package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18q-.425 0-.713-.288T6 17v-2h13V6h2q.425 0 .713.288T22 7v15l-4-4H7Zm-5-1V3q0-.425.288-.713T3 2h13q.425 0 .713.288T17 3v9q0 .425-.288.713T16 13H6l-4 4Z"/>`),
		g.Group(children),
	)
}