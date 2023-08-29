package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GfScanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.12 24h5.122m-5.122 7.88V16.12H35m-11.559 5.22a5.22 5.22 0 0 0-5.22-5.22h0A5.22 5.22 0 0 0 13 21.34v5.32a5.22 5.22 0 0 0 5.22 5.22h0a5.22 5.22 0 0 0 5.221-5.22h-5.22"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.707 37.52A21.4 21.4 0 0 0 45.5 24c0-11.874-9.626-21.5-21.5-21.5S2.5 12.126 2.5 24S12.126 45.5 24 45.5a21.399 21.399 0 0 0 13.943-5.153l4.883 2.044l-2.12-4.87Z"/>`),
		g.Group(children),
	)
}