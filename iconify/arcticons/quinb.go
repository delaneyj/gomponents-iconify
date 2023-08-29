package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quinb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="M20.368 24.89a13.6 13.6 0 0 1-3.48 3.08a8.274 8.274 0 1 1 3.07-11.294c1.136 1.983 4.274 7.682 4.274 7.682s2.692 4.634 3.841 6.609a8.275 8.275 0 1 0 2.989-11.315a13.534 13.534 0 0 0-3.43 3.05M19.203 35.22a2.342 2.342 0 1 1 3.085 3.026m6.626-25.466a2.342 2.342 0 1 1-3.086-3.026"/>`),
		g.Group(children),
	)
}