package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 6a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm0 14a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm0-7a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/>`),
		g.Group(children),
	)
}