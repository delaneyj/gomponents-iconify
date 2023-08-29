package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 80v72a8 8 0 0 1-13.66 5.66L196 127.31l-70.34 70.35a8 8 0 0 1-11.32 0l-96-96a8 8 0 0 1 11.32-11.32L120 180.69L184.69 116l-30.35-30.34A8 8 0 0 1 160 72h72a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}