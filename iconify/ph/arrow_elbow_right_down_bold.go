package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightDownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m232.49 168.49l-48 48a12 12 0 0 1-17 0l-48-48a12 12 0 1 1 17-17L164 179V76H32a12 12 0 0 1 0-24h144a12 12 0 0 1 12 12v115l27.51-27.52a12 12 0 0 1 17 17Z"/>`),
		g.Group(children),
	)
}