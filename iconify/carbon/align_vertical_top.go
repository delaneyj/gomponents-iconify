package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 20h-4a2.002 2.002 0 0 1-2-2v-7a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v7a2.002 2.002 0 0 1-2 2zm-4-9v7h4.001L24 11zm-8 17H8a2.002 2.002 0 0 1-2-2V11a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v15a2.002 2.002 0 0 1-2 2zM8 11v15h4.001L12 11zM2 4h28v2H2z"/>`),
		g.Group(children),
	)
}