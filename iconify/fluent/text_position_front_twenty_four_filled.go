package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionFrontTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 3.75a1 1 0 0 0 0 2h16.5a1 1 0 1 0 0-2H3.75Zm8.25 5c-.78 0-1.467.397-1.871 1H7.937a4.252 4.252 0 0 1 8.126 0h-2.192a2.248 2.248 0 0 0-1.871-1Zm-8.25 2a1 1 0 1 0 0 2h16.5a1 1 0 1 0 0-2H3.75Zm4 3h2v2a1 1 0 1 1-2 0v-2Zm8.5 0h-2v2a1 1 0 1 0 2 0v-2Zm-13.5 5a1 1 0 0 1 1-1h16.5a1 1 0 1 1 0 2H3.75a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}