package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyFrancRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21q-.425 0-.713-.288T7 20v-2H6q-.425 0-.713-.288T5 17q0-.425.288-.713T6 16h1V4q0-.425.288-.713T8 3h9q.425 0 .713.288T18 4q0 .425-.288.713T17 5H9v6h7q.425 0 .713.288T17 12q0 .425-.288.713T16 13H9v3h3q.425 0 .713.288T13 17q0 .425-.288.713T12 18H9v2q0 .425-.288.713T8 21Z"/>`),
		g.Group(children),
	)
}