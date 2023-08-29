package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GripHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 8a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm3 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm3 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm3 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm3 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2a1 1 0 0 1 0-2z"/>`),
		g.Group(children),
	)
}