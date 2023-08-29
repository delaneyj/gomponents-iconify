package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5h2.5a3 3 0 0 1 3-3a3 3 0 0 1 3 3H17a3 3 0 0 1 3 3v11a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3m0 1a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-1v3H7V6H6m2 2h7V6H8v2m3.5-5a2 2 0 0 0-2 2h4a2 2 0 0 0-2-2M8 19v-2H6v-1h2v-2h1v2h2v1H9v2H8Z"/>`),
		g.Group(children),
	)
}