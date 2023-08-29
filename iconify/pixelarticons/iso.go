package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3H6v3H3v2h3v3h2V8h3V6H8V3zm11 2h-2v2h-2v2h-2v2h-2v2H9v2H7v2H5v2h2v-2h2v-2h2v-2h2v-2h2V9h2V7h2V5zm-6 13v-2h8v2h-8z"/>`),
		g.Group(children),
	)
}