package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataExplorationOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12v10H12Zm0-2q3.35 0 5.675-2.325T20 12q0-3.35-2.325-5.675T12 4Q8.65 4 6.325 6.325T4 12q0 .575.075 1.125T4.3 14.2L8 10.5l3.3 2.775L14.575 10H13V8h5v5h-2v-1.575L11.4 16l-3.275-2.8l-2.95 2.95q1.05 1.725 2.838 2.788T12 20Zm7.5.5q.425 0 .713-.288t.287-.712q0-.425-.288-.713T19.5 18.5q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287ZM11.375 12Z"/>`),
		g.Group(children),
	)
}