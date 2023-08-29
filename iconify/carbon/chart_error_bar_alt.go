package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartErrorBarAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 10V8H8v2h2v4.184a2.983 2.983 0 0 0 0 5.632V24H8v2h6v-2h-2v-4.184a2.983 2.983 0 0 0 0-5.632V10zm12-4V4h-6v2h2v2.184a2.983 2.983 0 0 0 0 5.632V18h-2v2h6v-2h-2v-4.184a2.983 2.983 0 0 0 0-5.632V6z"/><path fill="currentColor" d="M30 30H4a2 2 0 0 1-2-2V2h2v26h26Z"/>`),
		g.Group(children),
	)
}