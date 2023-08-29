package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 16V8H8v8h12zm-10-6h8v4h-8v-4zM4 20h2V4H4v16z"/>`),
		g.Group(children),
	)
}