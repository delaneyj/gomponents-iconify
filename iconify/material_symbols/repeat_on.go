package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23q-.825 0-1.413-.588T1 21V3q0-.825.588-1.413T3 1h18q.825 0 1.413.588T23 3v18q0 .825-.588 1.413T21 23H3Zm4-1l1.4-1.45L6.85 19H19v-6h-2v4H6.85l1.55-1.55L7 14l-4 4l4 4ZM5 11h2V7h10.15L15.6 8.55L17 10l4-4l-4-4l-1.4 1.45L17.15 5H5v6Z"/>`),
		g.Group(children),
	)
}