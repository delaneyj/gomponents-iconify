package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.512 17.023L10 14l-7-3.5a.55.55 0 0 1 0-1L21 3l-4.45 12.324M15 19l2 2l4-4"/>`),
		g.Group(children),
	)
}