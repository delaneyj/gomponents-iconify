package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SensorDoorSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h16v20H4Zm11.5-8.5q.625 0 1.063-.438T17 12q0-.625-.438-1.063T15.5 10.5q-.625 0-1.063.438T14 12q0 .625.438 1.063t1.062.437Z"/>`),
		g.Group(children),
	)
}