package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowReplyTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.03 5.47a.75.75 0 0 1 0 1.06L5.56 11h8.69C20.187 11 25 15.813 25 21.75a.75.75 0 0 1-1.5 0a9.25 9.25 0 0 0-9.25-9.25H5.56l4.47 4.47a.75.75 0 1 1-1.06 1.06l-5.75-5.75a.75.75 0 0 1 0-1.06l5.75-5.75a.75.75 0 0 1 1.06 0Z"/>`),
		g.Group(children),
	)
}