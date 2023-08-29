package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaShareOutline0"><g id="evaShareOutline1"><path id="evaShareOutline2" fill="currentColor" d="M18 15a3 3 0 0 0-2.1.86L8 12.34v-.67l7.9-3.53A3 3 0 1 0 15 6v.34L7.1 9.86a3 3 0 1 0 0 4.28l7.9 3.53V18a3 3 0 1 0 3-3Zm0-10a1 1 0 1 1-1 1a1 1 0 0 1 1-1ZM5 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm13 6a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`),
		g.Group(children),
	)
}