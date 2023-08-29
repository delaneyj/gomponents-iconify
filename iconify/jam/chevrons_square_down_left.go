package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsSquareDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4zm0 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4zm4 10h3a1 1 0 0 1 0 2H7a1 1 0 0 1-1-1V9a1 1 0 1 1 2 0v3zm3-3h3a1 1 0 0 1 0 2h-4a1 1 0 0 1-1-1V6a1 1 0 1 1 2 0v3z"/>`),
		g.Group(children),
	)
}