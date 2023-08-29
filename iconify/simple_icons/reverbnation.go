package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reverbnation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m24 9.324l-9.143-.03L11.971.57L9.143 9.294L0 9.324h.031l7.367 5.355l-2.855 8.749h.029l7.459-5.386l7.396 5.386l-2.855-8.73L24 9.315"/>`),
		g.Group(children),
	)
}