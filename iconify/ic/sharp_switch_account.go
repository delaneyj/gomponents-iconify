package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSwitchAccount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6H2v16h16v-2H4V6zm2-4v16h16V2H6zm8 3c1.66 0 3 1.34 3 3s-1.34 3-3 3s-3-1.34-3-3s1.34-3 3-3zM7.76 16c1.47-1.83 3.71-3 6.24-3s4.77 1.17 6.24 3H7.76z"/>`),
		g.Group(children),
	)
}