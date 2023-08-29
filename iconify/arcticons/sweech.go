package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sweech(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.276 24l-5.846-5.848L42.5 6.077H11.323l-5.089 5.09v7.343L11.723 24m18.707-5.848l-6.146-6.148M11.724 24l5.846 5.848L5.5 41.923h31.177l5.088-5.09V29.49L36.277 24M17.57 29.848l6.146 6.148"/>`),
		g.Group(children),
	)
}