package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microwave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zm12-1v12m3-6h.01M18 15h.01M18 9h.01"/><path d="M6.5 10.5c1-.667 1.5-.667 2.5 0c.833.347 1.667.926 2.5 0m-5 3c1-.667 1.5-.667 2.5 0c.833.347 1.667.926 2.5 0"/></g>`),
		g.Group(children),
	)
}