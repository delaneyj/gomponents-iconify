package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 8V6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v2m6 0h4a2 2 0 0 1 2 2v3m-6-5H9m0 0H5a2 2 0 0 0-2 2v3m0 0v5a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5M3 13h7m11 0h-7m0 0v-2h-4v2m4 0v2h-4v-2"/>`),
		g.Group(children),
	)
}