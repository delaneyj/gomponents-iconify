package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M8 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8zm0 1A5 5 0 1 0 8 3a5 5 0 0 0 0 10z"/><path d="M10 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0z"/></g>`),
		g.Group(children),
	)
}