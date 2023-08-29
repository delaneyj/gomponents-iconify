package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.75 3a.75.75 0 0 0 0 1.5h7.67L3.22 19.7a.764.764 0 1 0 1.081 1.081l15.2-15.2v7.669a.75.75 0 0 0 1.5 0v-9.5a.75.75 0 0 0-.75-.75h-9.5Z"/>`),
		g.Group(children),
	)
}