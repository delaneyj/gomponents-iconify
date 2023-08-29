package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicExternalOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 6.65L5.35 2.5q.35-.275.775-.388T7 2q1.25 0 2.125.862T10 5q0 .45-.138.863T9.5 6.65ZM20 17.15l-2-2V6q0-.825-.588-1.413T16 4q-.825 0-1.413.588T14 6v5.15l-2-2V6q0-1.7 1.175-2.85T16 2q1.65 0 2.825 1.15T20 6v11.15Zm-.2 5.45L14 16.8V18q0 1.65-1.175 2.825T10 22q-1.65 0-2.825-1.175T6 18h-.55q-.2 0-.338-.125t-.162-.325L4.1 9.1q-.05-.45.25-.775T5.1 8h.05L1.4 4.2q-.275-.275-.287-.688T1.4 2.8q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275ZM10 20q.825 0 1.413-.588T12 18v-3.2l-2.45-2.45l-.5 5.2q-.025.2-.163.325T8.55 18H8q0 .825.588 1.413T10 20Z"/>`),
		g.Group(children),
	)
}