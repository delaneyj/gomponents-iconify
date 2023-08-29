package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainbowBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M256 172v8a12 12 0 0 1-24 0v-8a104 104 0 0 0-208 0v8a12 12 0 0 1-24 0v-8a128 128 0 0 1 256 0Zm-128-32a36 36 0 0 0-36 36v4a12 12 0 0 0 24 0v-4a12 12 0 0 1 24 0v4a12 12 0 0 0 24 0v-4a36 36 0 0 0-36-36Zm0-48a84.09 84.09 0 0 0-84 84v4a12 12 0 0 0 24 0v-4a60 60 0 0 1 120 0v4a12 12 0 0 0 24 0v-4a84.09 84.09 0 0 0-84-84Z"/>`),
		g.Group(children),
	)
}