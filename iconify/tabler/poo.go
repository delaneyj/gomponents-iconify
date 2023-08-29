package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 12h.01M14 12h.01M10 16a3.5 3.5 0 0 0 4 0"/><path d="M11 4c2 0 3.5 1.5 3.5 4h.164a2.5 2.5 0 0 1 2.196 3.32a3 3 0 0 1 1.615 3.063a3 3 0 0 1-1.299 5.607H7a3 3 0 0 1-1.474-5.613a3 3 0 0 1 1.615-3.062a2.5 2.5 0 0 1 2.195-3.32H9.5c1.5 0 2.5-2 1.5-4z"/></g>`),
		g.Group(children),
	)
}