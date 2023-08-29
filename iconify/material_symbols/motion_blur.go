package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotionBlur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17q-.425 0-.713-.288T7 16q0-.425.288-.713T8 15h5q-.35-.425-.563-.925T12.1 13H10q-.425 0-.712-.288T9 12q0-.425.288-.713T10 11h2.1q.125-.575.338-1.075T13 9H4q-.425 0-.713-.288T3 8q0-.425.288-.713T4 7h13q2.075 0 3.538 1.463T22 12q0 2.075-1.463 3.538T17 17H8Zm9-2q1.25 0 2.125-.875T20 12q0-1.25-.875-2.125T17 9q-1.25 0-2.125.875T14 12q0 1.25.875 2.125T17 15ZM3 13q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h4q.425 0 .713.288T8 12q0 .425-.288.713T7 13H3Zm1 4q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15h1q.425 0 .713.288T6 16q0 .425-.288.713T5 17H4Z"/>`),
		g.Group(children),
	)
}