package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 2a1 1 0 0 0-1 1v3H4a1 1 0 0 0-1 1v9a3 3 0 0 0 3 3h1v2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-2h1a3 3 0 0 0 3-3V7a1 1 0 0 0-1-1h-3V3a1 1 0 0 0-1-1H8Zm0 12a1 1 0 0 0-1 1v2H6a1 1 0 0 1-1-1V8h14v8a1 1 0 0 1-1 1h-1v-2a1 1 0 0 0-1-1H8Zm4-2a1 1 0 1 0 0-2H8a1 1 0 1 0 0 2h4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}