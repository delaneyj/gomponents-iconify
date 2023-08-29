package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TypeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 5a1 1 0 0 0-1 1v1a1 1 0 0 1-2 0V6a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v1a1 1 0 1 1-2 0V6a1 1 0 0 0-1-1h-4v14h2a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h2V5H7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}