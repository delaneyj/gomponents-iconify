package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UntisMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.981 2.5A21.481 21.481 0 1 1 2.5 23.981h0A21.491 21.491 0 0 1 23.981 2.5Zm-.005 21.526h17.229M23.976 6.798v17.229m12.183-12.183L23.977 24.026"/>`),
		g.Group(children),
	)
}