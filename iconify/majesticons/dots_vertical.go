package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M10 5a2 2 0 1 0 4 0a2 2 0 0 0-4 0zm0 7a2 2 0 1 0 4 0a2 2 0 0 0-4 0zm0 7a2 2 0 1 0 4 0a2 2 0 0 0-4 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}