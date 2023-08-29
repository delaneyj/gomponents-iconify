package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjustments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M20 7a1 1 0 1 0 0-2h-9.17a3.001 3.001 0 0 0-5.66 0H4a1 1 0 1 0 0 2h1.17a3.001 3.001 0 0 0 5.66 0H20zm0 6a1 1 0 1 0 0-2h-1.17a3.001 3.001 0 0 0-5.66 0H4a1 1 0 1 0 0 2h9.17a3.001 3.001 0 0 0 5.66 0H20zm0 6a1 1 0 1 0 0-2h-9.17a3.001 3.001 0 0 0-5.66 0H4a1 1 0 1 0 0 2h1.17a3.001 3.001 0 0 0 5.66 0H20z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}