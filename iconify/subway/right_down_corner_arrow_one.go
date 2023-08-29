package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightDownCornerArrowOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m512 512l-23.3-302.5L384 314.2L69.8 0L0 69.8L314.2 384L209.5 488.7z"/>`),
		g.Group(children),
	)
}