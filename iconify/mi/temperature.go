package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Temperature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 5a3 3 0 1 1 6 0v8a5 5 0 1 1-6 0V5zm3-1a1 1 0 0 0-1 1v8.535a1 1 0 0 1-.5.866a3 3 0 1 0 2.999 0a1 1 0 0 1-.499-.866V5a1 1 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}