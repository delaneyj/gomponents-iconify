package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3h12v18H6V3zm2 2v4h3V5H8zm5 0v4h3V5h-3zm3 6H8v8h8v-8z"/>`),
		g.Group(children),
	)
}