package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignFactorialBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="3.75" d="M12 19.5h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M11.523 4.5L12 15l.477-10.5a.478.478 0 1 0-.954 0Z"/></g>`),
		g.Group(children),
	)
}