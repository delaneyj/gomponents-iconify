package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonPinRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20H5q-.825 0-1.413-.588T3 18V4q0-.825.588-1.413T5 2h14q.825 0 1.413.588T21 4v14q0 .825-.588 1.413T19 20h-4l-2.3 2.3q-.15.15-.325.213t-.375.062q-.2 0-.375-.063T11.3 22.3L9 20Zm3-8q1.45 0 2.475-1.025T15.5 8.5q0-1.45-1.025-2.475T12 5q-1.45 0-2.475 1.025T8.5 8.5q0 1.45 1.025 2.475T12 12Zm-7 6h14v-1.15q-1.35-1.325-3.138-2.087T12 14q-2.075 0-3.863.763T5 16.85V18Z"/>`),
		g.Group(children),
	)
}