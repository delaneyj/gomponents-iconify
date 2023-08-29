package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutAlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20H8V8h8v12zm-6-10v8h4v-8h-4zm10-6v2H4V4h16z"/>`),
		g.Group(children),
	)
}