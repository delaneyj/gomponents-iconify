package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Temperature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 15a5 5 0 1 1-8-4V3a3 3 0 1 1 6 0v8c1.214.912 2 2.364 2 4zm-3.201-2.401l-.799-.6V3a1 1 0 1 0-2 0v8.999l-.799.6a3 3 0 1 0 3.598 0zM5 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4z"/>`),
		g.Group(children),
	)
}