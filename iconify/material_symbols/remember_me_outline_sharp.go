package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RememberMeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 23V1h14v22H5Zm2-3v1h10v-1H7ZM7 4h10V3H7v1Zm5 12q-1.35 0-2.625.388T7 17.5v.5h10v-.5q-1.1-.725-2.375-1.113T12 16Zm0-2q1.35 0 2.613.313T17 15.2V6H7v9.2q1.125-.575 2.388-.887T12 14Zm0-1q1.25 0 2.125-.875T15 10q0-1.25-.875-2.125T12 7q-1.25 0-2.125.875T9 10q0 1.25.875 2.125T12 13Zm0-2q-.425 0-.713-.288T11 10q0-.425.288-.713T12 9q.425 0 .713.288T13 10q0 .425-.288.713T12 11Zm0 7h5H7h5Zm0-8Zm0-6Zm0 16Z"/>`),
		g.Group(children),
	)
}