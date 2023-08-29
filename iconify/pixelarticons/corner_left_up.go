package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18V8H6V6h2V4h2v2h2v2h-2v10h10v2H8v-2zm4-10h2v2h-2V8zM6 8H4v2h2V8z"/>`),
		g.Group(children),
	)
}