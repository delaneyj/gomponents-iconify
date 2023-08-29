package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderGenderfluid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 15.464a4 4 0 1 0 4-6.928a4 4 0 0 0-4 6.928zM15.464 14l3-5.196M5.536 15.195l3-5.196M12 12h.01M9 9L3 3m2.5 5.5l3-3M21 21l-6-6m2 5l3-3M3 7V3h4"/>`),
		g.Group(children),
	)
}