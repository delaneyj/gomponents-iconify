package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GripVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0zM7 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0zM7 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm-3 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm-3 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`),
		g.Group(children),
	)
}