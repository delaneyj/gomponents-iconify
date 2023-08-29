package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowArcLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 184a8 8 0 0 1-16 0a88 88 0 0 0-148.53-63.84l26.19 26.18A8 8 0 0 1 88 160H24a8 8 0 0 1-8-8V88a8 8 0 0 1 13.66-5.66l26.48 26.48A104 104 0 0 1 232 184Z"/>`),
		g.Group(children),
	)
}