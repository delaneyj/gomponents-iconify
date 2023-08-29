package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapCallsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15.15V8q0-1.65 1.175-2.825T9 4q1.65 0 2.825 1.175T13 8v7q0 .825.588 1.413T15 17q.825 0 1.413-.588T17 15V7.85l-.875.875q-.3.3-.713.288T14.7 8.7q-.275-.3-.287-.7t.287-.7l2.6-2.6q.15-.15.325-.212T18 4.425q.2 0 .375.063t.325.212l2.6 2.6q.3.3.288.7t-.288.7q-.3.3-.713.313t-.712-.288L19 7.85V15q0 1.65-1.175 2.825T15 19q-1.65 0-2.825-1.175T11 15V8q0-.825-.588-1.413T9 6q-.825 0-1.413.588T7 8v7.15l.875-.875q.3-.3.713-.288t.712.313q.275.3.288.7t-.288.7l-2.6 2.6q-.15.15-.325.213T6 18.575q-.2 0-.375-.063T5.3 18.3l-2.6-2.6q-.3-.3-.288-.7t.288-.7q.3-.3.713-.312t.712.287L5 15.15Z"/>`),
		g.Group(children),
	)
}