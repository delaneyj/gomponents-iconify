package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h14V3H7v14zm4.25-5.53l1.96 2.36l2.75-3.54L19.5 15h-11l2.75-3.53z" opacity=".3"/><path fill="currentColor" d="M1 21c0 1.1.9 2 2 2h16v-2H3V5H1v16zM21 1H7c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2zm0 16H7V3h14v14zm-5.04-6.71l-2.75 3.54l-1.96-2.36L8.5 15h11z"/>`),
		g.Group(children),
	)
}