package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoI(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7q-.825 0-1.413-.588T10 5q0-.825.588-1.413T12 3q.825 0 1.413.588T14 5q0 .825-.588 1.413T12 7Zm-1.5 14V9h3v12h-3Z"/>`),
		g.Group(children),
	)
}