package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 19V9a2 2 0 0 1 2-2h.5a2 2 0 0 0 1.6-.8l2.22-2.96A.6.6 0 0 1 8.8 3h6.4a.6.6 0 0 1 .48.24L17.9 6.2a2 2 0 0 0 1.6.8h.5a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path d="M12 17a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`),
		g.Group(children),
	)
}