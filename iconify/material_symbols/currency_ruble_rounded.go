package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRubleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21q-.425 0-.713-.288T7 20v-2H6q-.425 0-.713-.288T5 17q0-.425.288-.713T6 16h1v-2H6q-.425 0-.713-.288T5 13q0-.425.288-.713T6 12h1V4q0-.425.288-.713T8 3h5.5q2.3 0 3.9 1.6T19 8.5q0 2.3-1.6 3.9T13.5 14H9v2h3q.425 0 .713.288T13 17q0 .425-.288.713T12 18H9v2q0 .425-.288.713T8 21Zm1-9h4.5q1.45 0 2.475-1.025T17 8.5q0-1.45-1.025-2.475T13.5 5H9v7Z"/>`),
		g.Group(children),
	)
}