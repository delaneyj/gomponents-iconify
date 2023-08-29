package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBarLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4v16H4V4h2zm14 7v2h-8v2h-2v-2H8v-2h2V9h2v2h8zm-8-2V7h2v2h-2zm0 6h2v2h-2v-2z"/>`),
		g.Group(children),
	)
}