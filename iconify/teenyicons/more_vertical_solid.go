package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 2.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm0 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm0 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}