package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitscreenRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 21q-.825 0-1.413-.588T13 19V5q0-.825.588-1.413T15 3h4q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21h-4ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h4q.825 0 1.413.588T11 5v14q0 .825-.588 1.413T9 21H5ZM5 5v14h4V5H5Zm4 0v14H5V5h4Z"/>`),
		g.Group(children),
	)
}