package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineUpgrade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 18v2H8v-2h8zM11 7.99V16h2V7.99h3L12 4L8 7.99h3z"/>`),
		g.Group(children),
	)
}