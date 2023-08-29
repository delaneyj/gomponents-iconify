package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftDownCornerArrowOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 69.8L442.2 0L128 314.2L23.3 209.5L0 512l302.5-23.3L197.8 384z"/>`),
		g.Group(children),
	)
}