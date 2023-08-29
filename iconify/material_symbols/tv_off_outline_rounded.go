package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.325 18.475L20 17.15V5H7.85l-2-2H20q.825 0 1.413.588T22 5v12q0 .45-.163.813t-.512.662Zm-18.15-15.3L5 5H4v12h10.15L1.4 4.2q-.275-.275-.287-.688T1.4 2.8q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L16.2 19H16v1q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20v-1H4q-.825 0-1.413-.588T2 17V5q0-.925.588-1.375l.587-.45ZM9.1 11.95Zm4.875-.825Z"/>`),
		g.Group(children),
	)
}