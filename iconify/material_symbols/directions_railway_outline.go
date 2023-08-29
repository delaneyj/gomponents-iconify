package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsRailwayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-1l1.5-1q-1.475 0-2.488-1.012T4 15.5V6q0-2.075 1.925-3.038T12 2q4.3 0 6.15.925T20 6v9.5q0 1.475-1.012 2.488T16.5 19l1.5 1v1H6Zm0-11h12V7H6v3Zm10.5 2H6h12h-1.5ZM12 16q.625 0 1.063-.438T13.5 14.5q0-.625-.438-1.063T12 13q-.625 0-1.063.438T10.5 14.5q0 .625.438 1.063T12 16Zm-4.5 1h9q.65 0 1.075-.425T18 15.5V12H6v3.5q0 .65.425 1.075T7.5 17ZM12 4q-2.15 0-3.563.25T6.45 5h11.2q-.45-.5-1.862-.75T12 4Zm0 1h5.65h-11.2H12Z"/>`),
		g.Group(children),
	)
}