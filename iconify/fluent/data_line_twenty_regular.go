package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataLineTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.5 4a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM13 5.5a2.5 2.5 0 1 1 1.557 2.316l-1.282 1.923a2.5 2.5 0 1 1-3.83 3.185l-2.496 1.07a2.5 2.5 0 1 1-.394-.919l2.496-1.07a2.5 2.5 0 0 1 3.392-2.822l1.282-1.922A2.492 2.492 0 0 1 13 5.5Zm-3 6a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Zm-7 3a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Z"/>`),
		g.Group(children),
	)
}