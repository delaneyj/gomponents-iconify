package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransportationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 16q-1.8 0-3.175-1.125T14.1 12h-2.975l-1.2-2H14.1q.125-.55.338-1.05T15 8H8.7L7.5 6h8.55l-1.1-3H11V1h5.35l1.825 5H19q2.075 0 3.538 1.463T24 11q0 2.075-1.463 3.538T19 16Zm0-2q1.25 0 2.125-.875T22 11q0-1.25-.875-2.125T19 8h-.075l.975 2.675l-1.9.675l-.95-2.625q-.5.425-.775 1.025T16 11q0 1.25.875 2.125T19 14ZM7 23q-1.25 0-2.125-.875T4 20H0v-6h2v-3H0V9h7l3 5h4v6h-4q0 1.25-.875 2.125T7 23Zm-3-9h3.675l-1.8-3H4v3Zm3 7q.425 0 .713-.288T8 20q0-.425-.288-.713T7 19q-.425 0-.713.288T6 20q0 .425.288.713T7 21Z"/>`),
		g.Group(children),
	)
}