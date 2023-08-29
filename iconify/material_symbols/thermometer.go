package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-2.075 0-3.538-1.463T7 16q0-1.2.525-2.238T9 12V6q0-1.25.875-2.125T12 3q1.25 0 2.125.875T15 6v6q.95.725 1.475 1.763T17 16q0 2.075-1.463 3.538T12 21Zm-1-11h2V6q0-.425-.288-.713T12 5q-.425 0-.713.288T11 6v4Z"/>`),
		g.Group(children),
	)
}