package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClinicalFOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M19 15v3h-3v2h3v3h2v-3h3v-2h-3v-3h-2Zm-2 11a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H17Zm-1 6a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H17a1 1 0 0 1-1-1Zm1 4a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H17Z"/><path fill-rule="evenodd" d="M17 7a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3h4a3 3 0 0 1 3 3v31a3 3 0 0 1-3 3H13a3 3 0 0 1-3-3V10a3 3 0 0 1 3-3h4Zm11 5a3 3 0 0 0 3-3h4a1 1 0 0 1 1 1v31a1 1 0 0 1-1 1H13a1 1 0 0 1-1-1V10a1 1 0 0 1 1-1h4a3 3 0 0 0 3 3h8Zm-8-6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}