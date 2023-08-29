package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Batterydoctor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.505 4.5h6.67a1 1 0 0 1 1 1v3h4.63a2 2 0 0 1 2 2v31a2 2 0 0 1-2 2h-17.61a2 2 0 0 1-2-2v-31a2 2 0 0 1 2-2h4.31v-3a1 1 0 0 1 1-1Zm-1 4h8.67"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.868 17.5L20 26h8l-5.868 8.5"/>`),
		g.Group(children),
	)
}