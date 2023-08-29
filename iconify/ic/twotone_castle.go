package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneCastle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 7H7v6H3v6h5v-1c0-2.21 1.79-4 4-4s4 1.79 4 4v1h5v-6h-4V7zm-6 5H9V9h2v3zm4 0h-2V9h2v3z" opacity=".3"/><path fill="currentColor" d="M21 9v2h-2V3h-2v2h-2V3h-2v2h-2V3H9v2H7V3H5v8H3V9H1v12h9v-3c0-1.1.9-2 2-2s2 .9 2 2v3h9V9h-2zm0 10h-5v-1c0-2.21-1.79-4-4-4s-4 1.79-4 4v1H3v-6h4V7h10v6h4v6z"/><path fill="currentColor" d="M9 9h2v3H9zm4 0h2v3h-2z"/>`),
		g.Group(children),
	)
}