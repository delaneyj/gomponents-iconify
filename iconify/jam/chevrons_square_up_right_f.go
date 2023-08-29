package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsSquareUpRightF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11v3a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1H6a1 1 0 1 0 0 2h3zm3-3v3a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1H9a1 1 0 1 0 0 2h3zM4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4z"/>`),
		g.Group(children),
	)
}