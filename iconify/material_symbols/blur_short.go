package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlurShort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 15q1.25 0 2.125-.875T18 12q0-1.25-.875-2.125T15 9q-1.25 0-2.125.875T12 12q0 1.25.875 2.125T15 15Zm0 2q-1.8 0-3.175-1.125T10.1 13H7q-.425 0-.713-.288T6 12q0-.425.288-.713T7 11h3.1q.125-.55.338-1.063T11 9H5q-.425 0-.713-.288T4 8q0-.425.288-.713T5 7h10q2.075 0 3.538 1.463T20 12q0 2.075-1.463 3.538T15 17Z"/>`),
		g.Group(children),
	)
}