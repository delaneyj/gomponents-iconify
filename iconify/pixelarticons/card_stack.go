package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardStack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h18v12H2V4h2zm16 10V6H4v8h16zm2 4H2v2h20v-2z"/>`),
		g.Group(children),
	)
}