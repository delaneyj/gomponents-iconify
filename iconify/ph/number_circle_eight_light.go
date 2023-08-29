package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleEightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm18.57-94.46a30 30 0 1 0-37.14 0a34 34 0 1 0 37.14 0ZM110 100a18 18 0 1 1 18 18a18 18 0 0 1-18-18Zm18 74a22 22 0 1 1 22-22a22 22 0 0 1-22 22Z"/>`),
		g.Group(children),
	)
}