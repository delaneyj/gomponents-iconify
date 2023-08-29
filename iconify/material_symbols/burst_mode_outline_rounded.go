package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BurstModeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 18V6q0-.425.288-.713T2 5q.425 0 .713.288T3 6v12q0 .425-.288.713T2 19q-.425 0-.713-.288T1 18Zm4 0V6q0-.425.288-.713T6 5q.425 0 .713.288T7 6v12q0 .425-.288.713T6 19q-.425 0-.713-.288T5 18Zm6 1q-.825 0-1.413-.588T9 17V7q0-.825.588-1.413T11 5h10q.825 0 1.413.588T23 7v10q0 .825-.588 1.413T21 19H11Zm0-2h10V7H11v10Zm4.5-3l-1-1.325q-.15-.2-.4-.188t-.4.213l-1.125 1.5q-.2.25-.038.525T13 15h6q.3 0 .45-.275t-.05-.525l-1.6-2.175q-.15-.2-.4-.2t-.4.2L15.5 14ZM11 17V7v10Z"/>`),
		g.Group(children),
	)
}