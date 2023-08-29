package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M20 12c0 1.657-3.582 3-8 3s-8-1.343-8-3m16 6c0 1.657-3.582 3-8 3s-8-1.343-8-3"/><ellipse cx="12" cy="6" rx="8" ry="3"/><path d="M4 6v12M20 6v12"/></g>`),
		g.Group(children),
	)
}