package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConciergeBell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 18a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2H2v-2Zm18-2a8 8 0 1 0-16 0m8-12v4m-2-4h4"/>`),
		g.Group(children),
	)
}