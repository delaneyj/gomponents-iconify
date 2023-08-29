package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpLabelOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 12l-4.97-7H8.66l10.7 10.73zM2 4l1 1v14h14l2 2l1.41-1.41L3.44 2.62z"/>`),
		g.Group(children),
	)
}