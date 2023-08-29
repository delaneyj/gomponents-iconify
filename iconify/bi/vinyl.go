package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vinyl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/><path d="M8 6a2 2 0 1 0 0 4a2 2 0 0 0 0-4zM4 8a4 4 0 1 1 8 0a4 4 0 0 1-8 0z"/><path d="M9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/></g>`),
		g.Group(children),
	)
}