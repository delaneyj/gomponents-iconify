package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="6" cy="6" r="3" fill="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18" cy="6" r="3" fill="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12" cy="18" r="3" fill="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path d="M6 9v1a2 2 0 0 0 2 2h4m6-3v1a2 2 0 0 1-2 2h-4m0 0v3"/></g>`),
		g.Group(children),
	)
}