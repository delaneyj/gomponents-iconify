package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitBranchLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 21v-6m0-6v6m0 0h8a2 2 0 0 0 2-2v-2"/><circle cx="7" cy="6" r="3"/><circle cx="17" cy="8" r="3"/></g>`),
		g.Group(children),
	)
}