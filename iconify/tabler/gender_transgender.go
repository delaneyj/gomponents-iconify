package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderTransgender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12a4 4 0 1 0 8 0a4 4 0 1 0-8 0m7-3l6-6m0 4V3h-4M9 9L3 3m0 4V3h4M5.5 8.5l3-3M12 16v5m-2.5-2h5"/>`),
		g.Group(children),
	)
}