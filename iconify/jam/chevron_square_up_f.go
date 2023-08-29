package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronSquareUpF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 9.172l3.536 3.535a1 1 0 0 0 1.414-1.414L10.707 7.05a1 1 0 0 0-1.414 0L5.05 11.293a1 1 0 0 0 1.414 1.414L10 9.172zM4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4z"/>`),
		g.Group(children),
	)
}