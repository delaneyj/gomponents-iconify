package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpMuseum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 11V9L12 2L2 9v2h2v9H2v2h20v-2h-2v-9h2zm-6 7h-2v-4l-2 3l-2-3v4H8v-7h2l2 3l2-3h2v7z"/>`),
		g.Group(children),
	)
}