package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowserUpdated(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-2l1-1H4q-.825 0-1.413-.588T2 16V5q0-.825.588-1.413T4 3h8v2H4v11h16v-3h2v3q0 .825-.588 1.413T20 18h-3l1 1v2H6Zm9-6l-5-5l1.4-1.4l2.6 2.575V3h2v8.175L18.6 8.6L20 10l-5 5Z"/>`),
		g.Group(children),
	)
}