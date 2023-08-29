package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightClassOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 13q-.825 0-1.413-.588T12 11V6q0-.825.588-1.413T14 4h2q.825 0 1.413.588T18 6v5q0 .825-.588 1.413T16 13h-2Zm0-2h2V6h-2v5Zm-4.5 7q-.675 0-1.2-.388t-.725-1.037L5 8V4h2v4l2.5 8H18v2H9.5ZM8 21v-2h10v2H8Zm6-15h2h-2Z"/>`),
		g.Group(children),
	)
}