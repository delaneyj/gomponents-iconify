package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 7h16v2H0V7zm7-1h2V3h2L8 0L5 3h2zm2 4H7v3H5l3 3l3-3H9z"/>`),
		g.Group(children),
	)
}