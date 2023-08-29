package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartHighLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 8h-3V6h-2v14H7v2h3v2h2V10h3V8zm12 2h-3V6h-2v12h-3v2h3v4h2V12h3v-2z"/><path fill="currentColor" d="M30 30H4a2 2 0 0 1-2-2V2h2v26h26Z"/>`),
		g.Group(children),
	)
}