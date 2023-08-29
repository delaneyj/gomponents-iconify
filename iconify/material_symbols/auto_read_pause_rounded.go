package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoReadPauseRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 14q.425 0 .713-.288T11 13V7q0-.425-.288-.713T10 6q-.425 0-.713.288T9 7v6q0 .425.288.713T10 14Zm4 0q.425 0 .713-.288T15 13V7q0-.425-.288-.713T14 6q-.425 0-.713.288T13 7v6q0 .425.288.713T14 14Zm-8 4l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Z"/>`),
		g.Group(children),
	)
}