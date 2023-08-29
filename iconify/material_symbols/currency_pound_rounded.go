package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyPoundRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.225 21q-.5 0-.862-.363T6 19.776q0-.25.163-.562t.462-.613q.575-.425 1.275-1.325T8.6 15q0-.275-.038-.525T8.476 14H7q-.425 0-.713-.287T6 13q0-.425.288-.713T7 12h.5q-.525-.825-1.012-1.738T6 8q0-2.3 1.6-3.9t3.9-1.6q1.425 0 2.575.637t1.9 1.688q.3.425.238.9T15.7 6.4q-.35.15-.737 0t-.663-.5q-.5-.65-1.213-1.025T11.5 4.5q-1.45 0-2.475 1.025T8 8q0 1.2.6 2t1.225 2H13q.425 0 .713.288T14 13q0 .425-.288.713T13 14h-2.475q.05.225.063.475T10.6 15q0 1.25-.438 2.25T9.1 19H14q.65 0 1.113-.238t.737-.662q.275-.425.65-.563t.725.063q.5.25.537.675t-.212.85q-.525.875-1.475 1.375T14 21H7.225Z"/>`),
		g.Group(children),
	)
}