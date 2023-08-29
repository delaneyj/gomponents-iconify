package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightDownCornerArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m511.9 186.1l-93.1-93V349L69.8 0L0 69.8l349 349H93.1l93 93.1l325.9.1z"/>`),
		g.Group(children),
	)
}