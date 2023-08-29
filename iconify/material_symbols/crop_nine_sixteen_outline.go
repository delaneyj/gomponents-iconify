package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropNineSixteenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21q-.825 0-1.413-.588T7 19V5q0-.825.588-1.413T9 3h6q.825 0 1.413.588T17 5v14q0 .825-.588 1.413T15 21H9ZM9 5v14h6V5H9Zm0 0v14V5Z"/>`),
		g.Group(children),
	)
}