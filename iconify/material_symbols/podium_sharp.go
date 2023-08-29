package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 4.5q0 .825-.588 1.413T10.5 6.5q-.325 0-.6-.088t-.55-.287q-.6.2-.963.725T8.025 8H21l-1 7h-2.675l.475-4.75H6.2L6.675 15H4L3 8h3q0-1.225.675-2.225T8.5 4.3q.075-.775.65-1.287T10.5 2.5q.825 0 1.413.588T12.5 4.5ZM8.675 20.25h6.65l.8-8.5h-8.25l.8 8.5Z"/>`),
		g.Group(children),
	)
}