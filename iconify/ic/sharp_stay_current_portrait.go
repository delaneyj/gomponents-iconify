package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpStayCurrentPortrait(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 1.01L5.01 1v22H19V1.01zM17 19H7V5h10v14z"/>`),
		g.Group(children),
	)
}