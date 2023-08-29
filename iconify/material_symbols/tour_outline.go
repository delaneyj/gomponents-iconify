package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TourOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22V2h2v2h14l-2 5l2 5H7v8H5ZM7 6v6v-6Zm5.5 5q.825 0 1.413-.588T14.5 9q0-.825-.588-1.413T12.5 7q-.825 0-1.413.588T10.5 9q0 .825.588 1.413T12.5 11ZM7 12h11.05l-1.2-3l1.2-3H7v6Z"/>`),
		g.Group(children),
	)
}