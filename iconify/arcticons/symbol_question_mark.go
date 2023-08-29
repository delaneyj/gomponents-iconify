package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolQuestionMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="34.748" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.08 19.423c0-1.911.774-3.641 2.027-4.894s2.982-2.027 4.893-2.027a6.92 6.92 0 0 1 6.92 6.92c0 1.912-.645 3.786-2.027 4.894c-1.432 1.149-4.918 3.031-4.918 5.96"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}