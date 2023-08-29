package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagePhotographyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 8c-.5 0-1.5.3-1.5 1.5S6.5 11 7 11m-5 6v3a2 2 0 0 0 2 2v0h11M2 17V6a2 2 0 0 1 2-2h3M2 17c1.403-.234 3.637-.293 5.945.243M15 22h3a2 2 0 0 0 2-2v-6m-5 8c-1.704-2.768-4.427-4.148-7.055-4.757m0 0c1.095-1.268 2.73-2.45 5.096-3.243M10 10V5a1 1 0 0 1 1-1h1l2-2h4l2 2h1a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H11a1 1 0 0 1-1-1z"/><circle cx="16" cy="7" r="1"/></g>`),
		g.Group(children),
	)
}