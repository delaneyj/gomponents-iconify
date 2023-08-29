package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 4.5q0 .825-.588 1.413T10.5 6.5q-.325 0-.6-.088t-.55-.287q-.6.2-.963.725T8.025 8H21l-1 7h-4.9v-2h3.175l.425-3H5.3l.425 3H8.9v2H4L3 8h3q0-1.225.675-2.225T8.5 4.3q.075-.775.65-1.287T10.5 2.5q.825 0 1.413.588T12.5 4.5ZM9.775 19h4.45l.575-6H9.2l.575 6ZM8 21l-.95-10h9.9L16 21H8Z"/>`),
		g.Group(children),
	)
}