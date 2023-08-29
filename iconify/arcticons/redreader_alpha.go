package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedreaderAlpha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.594 41.852c-.516 0-1.032-.516-1.032-1.032v-5.8m0 4.382a2.586 2.586 0 0 1-2.578 2.578h0a2.586 2.586 0 0 1-2.578-2.578v-1.675a2.586 2.586 0 0 1 2.578-2.579h0a2.586 2.586 0 0 1 2.578 2.579m4.938.773a7 7 0 1 1-14 0a7 7 0 0 1 14 0ZM24 28.57l7.827 7.845M15 37.4V10.5h9a9.035 9.035 0 0 1 0 18.07h-9m27.5 4.19V7.5a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h25.299"/>`),
		g.Group(children),
	)
}