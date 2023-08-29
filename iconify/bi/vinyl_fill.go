package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M8 6a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2z"/><path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zM4 8a4 4 0 1 0 8 0a4 4 0 0 0-8 0z"/></g>`),
		g.Group(children),
	)
}