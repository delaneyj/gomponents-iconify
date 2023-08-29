package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpcomingRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19v-5q0-.825.588-1.413T4 12h4.15q.375 0 .65.25t.325.6q.125.85 1 1.5T12 15q1 0 1.875-.65t1-1.5q.05-.35.325-.6t.65-.25H20q.825 0 1.413.588T22 14v5q0 .825-.588 1.413T20 21H4Zm12.9-10.9q-.275-.275-.275-.7t.275-.7l2.15-2.15q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L18.3 10.1q-.275.275-.7.275t-.7-.275Zm-9.8 0q-.275.275-.7.275t-.7-.275L3.55 7.95q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275L7.1 8.7q.275.275.275.7t-.275.7ZM12 8q-.425 0-.713-.288T11 7V4q0-.425.288-.713T12 3q.425 0 .713.288T13 4v3q0 .425-.288.713T12 8Z"/>`),
		g.Group(children),
	)
}