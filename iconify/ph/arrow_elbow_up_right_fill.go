package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowUpRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m213.66 85.66l-48 48A8 8 0 0 1 152 128V88H72v136a8 8 0 0 1-16 0V80a8 8 0 0 1 8-8h88V32a8 8 0 0 1 13.66-5.66l48 48a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}