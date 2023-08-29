package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutAlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 8v8h12V8H4zm10 6H6v-4h8v4zm6-10h-2v16h2V4z"/>`),
		g.Group(children),
	)
}