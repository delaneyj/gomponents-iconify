package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 20v-2m2-2h2M7.05 11.293l-1.414 1.414a4 4 0 0 0 5.657 5.657l1.415-1.414M6 8H4m4-4v2m3.293 1.05l1.414-1.414a4 4 0 1 1 5.656 5.657l-1.414 1.414"/>`),
		g.Group(children),
	)
}