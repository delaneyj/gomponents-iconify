package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddCol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h10v20H2v-2h8v-4H2v-2h8v-4H2V8h8V4H2V2zm17 9h3v2h-3v3h-2v-3h-3v-2h3V8h2v3z"/>`),
		g.Group(children),
	)
}