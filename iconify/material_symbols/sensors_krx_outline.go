package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SensorsKrxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q-2.1 0-3.55-1.45T7 12q0-2.075 1.45-3.538T12 7q2.075 0 3.538 1.463T17 12q0 2.1-1.463 3.55T12 17Zm0-2q1.25 0 2.125-.875T15 12q0-1.25-.875-2.125T12 9q-1.25 0-2.125.875T9 12q0 1.25.875 2.125T12 15Zm-7.8 1.5q-.575-1-.888-2.125T3 12q0-1.25.313-2.375T4.2 7.5l1.725 1q-.45.775-.687 1.663T5 12q0 .95.238 1.838t.687 1.662l-1.725 1Zm15.6 0l-1.725-1q.45-.775.688-1.663T19 12q0-.95-.238-1.838T18.075 8.5l1.725-1q.575 1 .888 2.125T21 12q0 1.25-.313 2.375T19.8 16.5ZM12 12Z"/>`),
		g.Group(children),
	)
}