package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermContactCalendarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q-1.4 0-2.675.438T7 19.75V20h10v-.25q-1.05-.875-2.325-1.313T12 18Zm-7 .85q1.35-1.325 3.138-2.087T12 16q2.075 0 3.863.763T19 18.85V6H5v12.85ZM12 14q-1.45 0-2.475-1.025T8.5 10.5q0-1.45 1.025-2.475T12 7q1.45 0 2.475 1.025T15.5 10.5q0 1.45-1.025 2.475T12 14Zm0-2q.625 0 1.063-.438T13.5 10.5q0-.625-.438-1.063T12 9q-.625 0-1.063.438T10.5 10.5q0 .625.438 1.063T12 12ZM5 22q-.825 0-1.413-.588T3 20V6q0-.825.588-1.413T5 4h1V2h2v2h8V2h2v2h1q.825 0 1.413.588T21 6v14q0 .825-.588 1.413T19 22H5Zm7-11.5Zm0 9.5h5H7h5Z"/>`),
		g.Group(children),
	)
}