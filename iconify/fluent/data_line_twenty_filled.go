package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataLineTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 5.5a2.5 2.5 0 0 1-3.443 2.316l-1.282 1.923a2.5 2.5 0 1 1-3.83 3.185l-2.496 1.07a2.5 2.5 0 1 1-.394-.92l2.496-1.069a2.5 2.5 0 0 1 3.392-2.822l1.282-1.922A2.5 2.5 0 1 1 18 5.5Z"/>`),
		g.Group(children),
	)
}