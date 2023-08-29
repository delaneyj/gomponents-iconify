package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpTabUnselected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 9h2V7H1v2zm0 4h2v-2H1v2zm8 8h2v-2H9v2zm-8-4h2v-2H1v2zm0 4h2v-2H1v2zM23 3H13v6h10V3zm-2 14h2v-2h-2v2zM9 5h2V3H9v2zM5 21h2v-2H5v2zM5 5h2V3H5v2zM1 5h2V3H1v2zm20 8h2v-2h-2v2zm-8 8h2v-2h-2v2zm4 0h2v-2h-2v2zm4 0h2v-2h-2v2z"/>`),
		g.Group(children),
	)
}