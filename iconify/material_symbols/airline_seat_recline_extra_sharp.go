package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatReclineExtraSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 6q-.825 0-1.413-.588T6.5 4q0-.825.588-1.413T8.5 2q.825 0 1.413.588T10.5 4q0 .825-.588 1.413T8.5 6ZM5.575 20L3 7h2.05l2.2 11H14v2H5.575ZM19.5 22l-2.9-5H8.025L6.6 10.05q-.275-1.2.563-2.125T9.2 7q.875 0 1.588.525T11.7 8.95L12.8 14h4.375l4.075 7l-1.75 1Z"/>`),
		g.Group(children),
	)
}