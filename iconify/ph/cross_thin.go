package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 76h-40a4 4 0 0 1-4-4V32a12 12 0 0 0-12-12h-32a12 12 0 0 0-12 12v40a4 4 0 0 1-4 4H56a12 12 0 0 0-12 12v32a12 12 0 0 0 12 12h40a4 4 0 0 1 4 4v88a12 12 0 0 0 12 12h32a12 12 0 0 0 12-12v-88a4 4 0 0 1 4-4h40a12 12 0 0 0 12-12V88a12 12 0 0 0-12-12Zm4 44a4 4 0 0 1-4 4h-40a12 12 0 0 0-12 12v88a4 4 0 0 1-4 4h-32a4 4 0 0 1-4-4v-88a12 12 0 0 0-12-12H56a4 4 0 0 1-4-4V88a4 4 0 0 1 4-4h40a12 12 0 0 0 12-12V32a4 4 0 0 1 4-4h32a4 4 0 0 1 4 4v40a12 12 0 0 0 12 12h40a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}