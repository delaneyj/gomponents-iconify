package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearAllRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15h12q.425 0 .713.288T17 16q0 .425-.288.713T16 17H4Zm2-4q-.425 0-.713-.288T5 12q0-.425.288-.713T6 11h12q.425 0 .713.288T19 12q0 .425-.288.713T18 13H6Zm2-4q-.425 0-.713-.288T7 8q0-.425.288-.713T8 7h12q.425 0 .713.288T21 8q0 .425-.288.713T20 9H8Z"/>`),
		g.Group(children),
	)
}