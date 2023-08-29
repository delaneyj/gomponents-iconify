package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 26h28v2H2zm22-3h-4a2.002 2.002 0 0 1-2-2v-7a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v7a2.002 2.002 0 0 1-2 2zm-4-9v7h4.001L24 14zm-8 9H8a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v15a2.002 2.002 0 0 1-2 2zM8 6v15h4.001L12 6z"/>`),
		g.Group(children),
	)
}