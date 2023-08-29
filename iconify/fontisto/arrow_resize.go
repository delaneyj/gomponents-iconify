package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowResize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.327 7.418L7.418 18.327l-1.745-1.745L16.582 5.673l-2.4-2.4h6.545l.001 6.546zM3.272 20.727v-6.545l6.546 6.546z"/>`),
		g.Group(children),
	)
}