package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Podium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15L3 8h3q0-1.225.675-2.225T8.5 4.3q.075-.775.65-1.287T10.5 2.5q.825 0 1.413.588T12.5 4.5q0 .825-.588 1.413T10.5 6.5q-.325 0-.6-.088t-.55-.287q-.6.2-.963.725T8.025 8H21l-1 7h-2.675l.175-1.725q.125-1.225-.688-2.125t-2.037-.9h-5.55q-1.225 0-2.038.9T6.5 13.275L6.675 15H4Zm4.675 5.25L8 13.125q-.05-.55.3-.963t.925-.412h5.55q.575 0 .925.413t.3.962l-.675 7.125h-6.65Z"/>`),
		g.Group(children),
	)
}