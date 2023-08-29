package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Member(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 7a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm3-5a5 5 0 1 0 0 10a5 5 0 0 0 0-10Zm7 0h-1v2h1a3 3 0 1 1 0 6h-1v2h1a5 5 0 0 0 0-10ZM0 19a5 5 0 0 1 5-5h7a5 5 0 0 1 5 5v2h-2v-2a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2H0v-2Zm24 0a5 5 0 0 0-5-5h-1v2h1a3 3 0 0 1 3 3v2h2v-2Z"/>`),
		g.Group(children),
	)
}