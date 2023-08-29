package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RailPlanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.677 31.401A12.643 12.643 0 0 1 11.3 24c0-7.015 5.686-12.701 12.701-12.701s12.7 5.686 12.7 12.701c0 2.762-.881 5.318-2.379 7.402l5.572 7.077A21.422 21.422 0 0 0 45.5 24c0-11.874-9.626-21.5-21.5-21.5S2.5 12.126 2.5 24a21.422 21.422 0 0 0 5.606 14.48l5.571-7.079Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.012 32.728H17.988l3.345-4.247h5.334l3.345 4.247zm6.689 8.496H11.299l3.345-4.248h18.712l3.345 4.248z"/>`),
		g.Group(children),
	)
}