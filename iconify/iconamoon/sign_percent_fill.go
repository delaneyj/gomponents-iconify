package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignPercentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.894 4.447a1 1 0 1 0-1.788-.894l-8 16a1 1 0 1 0 1.788.894l8-16ZM8 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM5 6a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm11 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}