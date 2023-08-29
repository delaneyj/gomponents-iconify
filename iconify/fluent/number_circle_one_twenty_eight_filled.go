package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleOneTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 14C2 7.373 7.373 2 14 2s12 5.373 12 12s-5.373 12-12 12S2 20.627 2 14Zm13.5-5.75a.75.75 0 0 0-1.474-.199l-.004.014a4.619 4.619 0 0 1-.127.35c-.1.246-.26.588-.499.967c-.477.758-1.251 1.638-2.456 2.185a.75.75 0 0 0 .62 1.366A7.02 7.02 0 0 0 14 11.087v8.163a.75.75 0 0 0 1.5 0v-11Zm-1.474-.199Z"/>`),
		g.Group(children),
	)
}