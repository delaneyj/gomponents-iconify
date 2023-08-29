package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollerShadesClosedOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22.75q-.725 0-1.238-.513T10.25 21H3q-.425 0-.713-.288T2 20q0-.425.288-.713T3 19h1V5q0-.825.588-1.413T6 3h12q.825 0 1.413.588T20 5v14h1q.425 0 .713.288T22 20q0 .425-.288.713T21 21h-7.25q0 .725-.513 1.238T12 22.75ZM6 15h12V5H6v10Zm0 4h5v-2H6v2Zm7 0h5v-2h-5v2ZM6 5h12H6Z"/>`),
		g.Group(children),
	)
}