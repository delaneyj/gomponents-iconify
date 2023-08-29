package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCollapseAllTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.75A.75.75 0 0 1 2.75 4h18.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 4.75Zm4.22 3.47a.75.75 0 0 1 1.06 0l3 3a.75.75 0 1 1-1.06 1.06L7.5 10.56v8.19a.75.75 0 0 1-1.5 0v-8.19l-1.72 1.72a.75.75 0 0 1-1.06-1.06l3-3Zm5.28.53a.75.75 0 0 1 .75-.75h9a.75.75 0 0 1 0 1.5h-9a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}