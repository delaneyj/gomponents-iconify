package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunningWaterAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M26 12a2 2 0 1 0-4 0v12h4V12Zm2 8v-8a4 4 0 0 0-8 0v8h-6v-1h4v-5a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v5h4v1H7a2 2 0 0 0-2 2v3H4v2h15v3h10v-3h15v-2h-1v-3a2 2 0 0 0-2-2h-4v-1h4v-5a2 2 0 0 0-2-2h-7a2 2 0 0 0-2 2v5h4v1h-6Zm-1 6h-6v2h6v-2Zm1 12a4 4 0 0 1-8 0c0-3.5 4-7 4-7s4 3.5 4 7ZM13 17h-1v-2h-2v2H9v-3h7v3h-1v-2h-2v2Zm22 0h1v-2h2v2h1v-3h-7v3h1v-2h2v2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}