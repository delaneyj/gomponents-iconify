package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDeleteBinThreeRemoveDeleteEmptyBinTrashGarbage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12 .5l-1.4 12.11a1 1 0 0 1-1 .89H4.39a1 1 0 0 1-1-.89L2 .5m-1 0h12m-6 0v13m-4.54-9h9.08M2.98 9h8.04"/>`),
		g.Group(children),
	)
}