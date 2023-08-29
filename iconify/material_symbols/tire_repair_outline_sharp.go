package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TireRepairOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 8q-.425 0-.713-.288T18 7q0-.2.088-.388T18.3 6.3q.3-.3 1.375-.675l1.075-.375l-.375 1.075q-.15.425-.338.825t-.337.55q-.125.125-.312.213T19 8ZM4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h6q.825 0 1.413.588T12 5v8h4v6h2v-6h-1v-1.425q-1.35-.575-2.175-1.8T14 7q0-2.075 1.463-3.538T19 2q2.075 0 3.538 1.463T24 7q0 1.55-.825 2.775T21 11.575V13h-1v8h-6v-6h-2v4q0 .825-.588 1.413T10 21H4Zm15-11q1.25 0 2.125-.875T22 7q0-1.25-.875-2.125T19 4q-1.25 0-2.125.875T16 7q0 1.25.875 2.125T19 10ZM4 19h6v-2l-2 2v-2.825l2-2V12l-2 2v-2.825l2-2V7L8 9V6.175L9.175 5h-4.35L6 6.175V9L4 7v2.175l2 2V14l-2-2v2.175l2 2V19l-2-2v2Zm3-7Z"/>`),
		g.Group(children),
	)
}