package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatPumpBalanceOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21.6q-.825 0-1.413-.588T1 19.6v-7q0-.425.288-.712T2 11.6h20q.425 0 .713.288T23 12.6v7q0 .825-.588 1.413T21 21.6H3Zm0-2h18v-6H3v6Zm18-6H3h18Zm-14 5q-1.25 0-2.125-.875T4 15.6v-8q0-.425.288-.712T5 6.6q.425 0 .713.288T6 7.6v8q0 .425.288.713T7 16.6q.425 0 .713-.288T8 15.6v-6q0-1.25.875-2.125T11 6.6q1.25 0 2.125.875T14 9.6v6q0 .425.288.713T15 16.6q.425 0 .713-.288T16 15.6v-7q0-1.25.875-2.125T19 5.6h1.175l-.475-.475q-.3-.3-.3-.713t.3-.712q.3-.3.7-.3t.7.3l2.175 2.175q.3.3.3.713t-.3.712L21.1 9.475q-.275.275-.675.263t-.7-.288q-.275-.275-.275-.7t.275-.7l.45-.45H19q-.425 0-.713.288T18 8.6v7q0 1.25-.875 2.125T15 18.6q-1.25 0-2.125-.875T12 15.6v-6q0-.425-.288-.712T11 8.6q-.425 0-.713.288T10 9.6v6q0 1.25-.875 2.125T7 18.6Z"/>`),
		g.Group(children),
	)
}