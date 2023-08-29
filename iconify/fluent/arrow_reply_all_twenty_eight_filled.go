package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowReplyAllTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.457 5.543a1 1 0 0 1 0 1.414L5.414 12l5.043 5.043a1 1 0 0 1-1.414 1.414l-5.75-5.75a1 1 0 0 1 0-1.414l5.75-5.75a1 1 0 0 1 1.414 0Zm5 0a1 1 0 0 1 0 1.414L11.414 11h2.836c6.075 0 11 4.925 11 11a1 1 0 1 1-2 0a9 9 0 0 0-9-9h-2.836l4.043 4.043a1 1 0 0 1-1.414 1.414l-5.75-5.75a1 1 0 0 1 0-1.414l5.75-5.75a1 1 0 0 1 1.414 0Z"/>`),
		g.Group(children),
	)
}