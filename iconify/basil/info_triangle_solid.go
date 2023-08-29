package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoTriangleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.73 3.993a2.749 2.749 0 0 1 4.54 0l.432.632a75.95 75.95 0 0 1 6.944 12.563l.09.208a2.511 2.511 0 0 1-2.024 3.497a69.412 69.412 0 0 1-15.424 0a2.511 2.511 0 0 1-2.024-3.497l.09-.208A75.95 75.95 0 0 1 9.298 4.625l.432-.632ZM13 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-1 2.75a.75.75 0 0 1 .75.75v5a.75.75 0 1 1-1.5 0v-5a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}