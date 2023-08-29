package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUUpLeftThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 144a60.07 60.07 0 0 1-60 60H80a4 4 0 0 1 0-8h88a52 52 0 0 0 0-104H41.66l41.17 41.17a4 4 0 0 1-5.66 5.66l-48-48a4 4 0 0 1 0-5.66l48-48a4 4 0 0 1 5.66 5.66L41.66 84H168a60.07 60.07 0 0 1 60 60Z"/>`),
		g.Group(children),
	)
}