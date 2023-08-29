package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmpathizeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8a2.5 2.5 0 1 0-2.5-2.5m2.817 6.815l-.317.317l-.728-.727a3.088 3.088 0 1 0-4.367 4.367L12 21.368l4.689-4.69m1.324-2.673a3.087 3.087 0 0 0-3.021-3.018M3 3l18 18"/>`),
		g.Group(children),
	)
}