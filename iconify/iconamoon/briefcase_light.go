package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="1.5" d="M5 16h14v3a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-3ZM4 7h16v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V7Z"/><path stroke-width="2.25" d="M12 12h.01v.01H12z"/><path stroke-width="1.5" d="M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2H9V5Z"/></g>`),
		g.Group(children),
	)
}