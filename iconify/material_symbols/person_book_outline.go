package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonBookOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q-1.4 0-2.675.438T7 19.75V20h10v-.25q-1.05-.875-2.325-1.313T12 18Zm0-2q1.725 0 3.225.525T18 18V4H6v14q1.275-.95 2.775-1.475T12 16Zm0-4q-.625 0-1.063-.438T10.5 10.5q0-.625.438-1.063T12 9q.625 0 1.063.438T13.5 10.5q0 .625-.438 1.063T12 12ZM6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H6Zm6-8q1.45 0 2.475-1.025T15.5 10.5q0-1.45-1.025-2.475T12 7q-1.45 0-2.475 1.025T8.5 10.5q0 1.45 1.025 2.475T12 14Zm0-3.5Z"/>`),
		g.Group(children),
	)
}