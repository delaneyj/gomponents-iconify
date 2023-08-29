package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Users(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="9" r="4" fill="currentColor"/><path fill="currentColor" d="M9 13c-3.866 0-7 2.686-7 6h14c0-3.314-3.134-6-7-6z"/><path d="M15 13a4 4 0 1 0-3-6.646m0 5.411c.897.942 2.193 1.235 3 1.235c3.866 0 7 2.686 7 6h-6"/></g>`),
		g.Group(children),
	)
}