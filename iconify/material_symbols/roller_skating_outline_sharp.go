package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RollerSkatingOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16V1h9v4.5q0 .675.413 1.225t1.062.725L20 8.975V16H4Zm2-2h12v-3.475l-4.1-1.15q-.9-.25-1.575-.888T11.3 7H9V6h2.1q-.05-.3-.063-.5T11 5H9V4h2V3H6v11Zm0 0Zm-1 9q-1.25 0-2.125-.875T2 20q0-1.25.875-2.125T5 17q1.25 0 2.125.875T8 20q0 1.25-.875 2.125T5 23Zm0-2q.425 0 .713-.288T6 20q0-.425-.288-.713T5 19q-.425 0-.713.288T4 20q0 .425.288.713T5 21Zm14 2q-1.25 0-2.125-.875T16 20q0-1.25.875-2.125T19 17q1.25 0 2.125.875T22 20q0 1.25-.875 2.125T19 23Zm0-2q.425 0 .713-.288T20 20q0-.425-.288-.713T19 19q-.425 0-.713.288T18 20q0 .425.288.713T19 21Zm-7 2q-1.25 0-2.125-.875T9 20q0-1.25.875-2.125T12 17q1.25 0 2.125.875T15 20q0 1.25-.875 2.125T12 23Zm0-2q.425 0 .713-.288T13 20q0-.425-.288-.713T12 19q-.425 0-.713.288T11 20q0 .425.288.713T12 21Zm-7-1Zm7 0Zm7 0Z"/>`),
		g.Group(children),
	)
}