package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ellipsis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`),
		g.Group(children),
	)
}