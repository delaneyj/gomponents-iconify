package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v3.465l7 4.667V2h2v8.132l7-4.667V2h2v13h1v2h-1v5h-2v-5H4v5H2v-5H1v-2h1V2h2Zm0 13h7v-2.465L4 7.87V15Zm9 0h7V7.87l-7 4.666V15Z"/>`),
		g.Group(children),
	)
}