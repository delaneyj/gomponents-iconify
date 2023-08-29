package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReorderAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm0 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm2 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm5-12a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm2 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm-2 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4z"/>`),
		g.Group(children),
	)
}