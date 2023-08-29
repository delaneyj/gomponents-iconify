package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/><path d="M10 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0zM8 4a4 4 0 0 0-4 4a.5.5 0 0 1-1 0a5 5 0 0 1 5-5a.5.5 0 0 1 0 1zm4.5 3.5a.5.5 0 0 1 .5.5a5 5 0 0 1-5 5a.5.5 0 0 1 0-1a4 4 0 0 0 4-4a.5.5 0 0 1 .5-.5z"/></g>`),
		g.Group(children),
	)
}