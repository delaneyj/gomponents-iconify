package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232.49 104.49a12 12 0 0 1-17 0L188 77v115a12 12 0 0 1-12 12H32a12 12 0 0 1 0-24h132V77l-27.51 27.52a12 12 0 1 1-17-17l48-48a12 12 0 0 1 17 0l48 48a12 12 0 0 1 0 16.97Z"/>`),
		g.Group(children),
	)
}