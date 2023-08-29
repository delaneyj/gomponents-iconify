package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsGymnastics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22l-.5-10L8 11H1V9h6l7-5l1.3 1.525L11.15 8.5H14L21.8 4L23 5.4L14.5 12L14 22h-2ZM6 8q-.825 0-1.413-.588T4 6q0-.825.588-1.413T6 4q.825 0 1.413.588T8 6q0 .825-.588 1.413T6 8Z"/>`),
		g.Group(children),
	)
}