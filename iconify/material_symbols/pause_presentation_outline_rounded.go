package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PausePresentationOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 16q.425 0 .713-.288T11 15V9q0-.425-.288-.713T10 8q-.425 0-.713.288T9 9v6q0 .425.288.713T10 16Zm4 0q.425 0 .713-.288T15 15V9q0-.425-.288-.713T14 8q-.425 0-.713.288T13 9v6q0 .425.288.713T14 16ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}