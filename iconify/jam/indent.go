package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Indent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 0h12a1 1 0 0 1 0 2H2a1 1 0 1 1 0-2zm0 8h12a1 1 0 0 1 0 2H2a1 1 0 1 1 0-2zm6-4h6a1 1 0 0 1 0 2H8a1 1 0 1 1 0-2zM4.44 5.857L2.382 7.091a1 1 0 0 1-1.515-.857V3.766a1 1 0 0 1 1.515-.857l2.056 1.234a1 1 0 0 1 0 1.714z"/>`),
		g.Group(children),
	)
}