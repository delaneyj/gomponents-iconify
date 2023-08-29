package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 16h3q.425 0 .713-.288T19 15V9q0-.425-.288-.713T18 8h-3q-.425 0-.713.288T14 9v6q0 .425.288.713T15 16Zm1-2v-4h1v4h-1Zm-7 2h3q.425 0 .713-.288T13 15V9q0-.425-.288-.713T12 8H9q-.425 0-.713.288T8 9v6q0 .425.288.713T9 16Zm1-2v-4h1v4h-1Zm-4 2q.425 0 .713-.288T7 15V9q0-.425-.288-.713T6 8q-.425 0-.713.288T5 9v6q0 .425.288.713T6 16Zm-2 4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}