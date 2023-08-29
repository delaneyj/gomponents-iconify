package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.725 18.5L15 17.125l2.275 1.375l-.6-2.6l2-1.725l-2.625-.225L15 11.5l-1.05 2.45l-2.625.225l2 1.725l-.6 2.6ZM6 14v2H4q-.825 0-1.413-.588T2 14V4q0-.825.588-1.413T4 2h10q.825 0 1.413.588T16 4v2h-2V4H4v10h2Zm4 8q-.825 0-1.413-.588T8 20V10q0-.825.588-1.413T10 8h10q.825 0 1.413.588T22 10v10q0 .825-.588 1.413T20 22H10Z"/>`),
		g.Group(children),
	)
}