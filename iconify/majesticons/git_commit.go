package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 16a4 4 0 1 1 0-8a4 4 0 0 1 0 8z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16a4 4 0 0 1 0-8m0 8a4 4 0 0 0 0-8m0 8v5m0-13V3"/></g>`),
		g.Group(children),
	)
}