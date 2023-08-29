package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPhonelinkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.56 4l-2.5-2.49L4.56 4zM24 8h-8v4.61l2 2V10h4v7h-1.61l3 3H24zm-2-2V4H7.39l2 2zM2.06 1.51L.65 2.92L2 4.27V17H0v3h17.73l2.35 2.35l1.41-1.41L2.06 1.51zM4 17V6.27L14.73 17H4z"/>`),
		g.Group(children),
	)
}