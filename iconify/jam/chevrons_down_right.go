package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 8V1a1 1 0 1 1 2 0v8a1 1 0 0 1-1 1H1a1 1 0 1 1 0-2h7zm4 4V5a1 1 0 0 1 2 0v8a1 1 0 0 1-1 1H5a1 1 0 0 1 0-2h7z"/>`),
		g.Group(children),
	)
}