package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm0 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm14 0a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm-7 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/>`),
		g.Group(children),
	)
}