package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScoreOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm4-8.5l4 4l6-6V5H5v11.5l4-4Zm3-.5V6h1.5v6H12Zm3.5 0l-2-3l2-3h1.7l-2 3l2 3h-1.7ZM7 12V8.25h2.5V7.5H7V6h4v3.75H8.5v.75H11V12H7Zm2 3l-4 4h14v-6l-6 6l-4-4Zm-4 4V5v14Z"/>`),
		g.Group(children),
	)
}