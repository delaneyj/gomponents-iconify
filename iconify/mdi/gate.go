package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 6v5H7V7H5v4H3V9H1v12h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h2v2h2V9h-2v2h-2V7h-2v4h-2V6h-2v5h-2V6H9m-6 7h2v4H3v-4m4 0h2v4H7v-4m4 0h2v4h-2v-4m4 0h2v4h-2v-4m4 0h2v4h-2v-4Z"/>`),
		g.Group(children),
	)
}