package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpStayPrimaryPortrait(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.01 1v22H19V1H5.01zM17 19H7V5h10v14z"/>`),
		g.Group(children),
	)
}