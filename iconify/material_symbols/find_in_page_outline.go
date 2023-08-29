package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FindInPageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15q.825 0 1.413-.588T14 13q0-.825-.588-1.413T12 11q-.825 0-1.413.588T10 13q0 .825.588 1.413T12 15Zm-6 7q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h8l6 6v12q0 .825-.588 1.413T18 22H6Zm6-10Zm-6 8h11.575L14 16.45q-.425.275-.925.413T12 17q-1.65 0-2.825-1.175T8 13q0-1.65 1.175-2.825T12 9q1.65 0 2.825 1.175T16 13q0 .575-.138 1.075T15.45 15L18 17.6V8.85L13.15 4H6v16Z"/>`),
		g.Group(children),
	)
}