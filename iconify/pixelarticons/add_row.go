package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddRow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 10V2H2v10h20V2h-2v8h-4V2h-2v8h-4V2H8v8H4zm9 9v3h-2v-3H8v-2h3v-3h2v3h3v2h-3z"/>`),
		g.Group(children),
	)
}