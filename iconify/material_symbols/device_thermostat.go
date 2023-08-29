package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceThermostat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.538-1.463T7 17q0-1.2.525-2.238T9 13V5q0-1.25.875-2.125T12 2q1.25 0 2.125.875T15 5v8q.95.725 1.475 1.763T17 17q0 2.075-1.463 3.538T12 22Zm-1-11h2v-1h-1V9h1V7h-1V6h1V5q0-.425-.288-.713T12 4q-.425 0-.713.288T11 5v6Z"/>`),
		g.Group(children),
	)
}