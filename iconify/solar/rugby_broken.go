package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RugbyBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M5.38 5.38c-2.128 2.13-2.993 5.194-3.27 8.077M13.457 2.11c1.67-.16 3.28-.125 4.612-.023a4.136 4.136 0 0 1 3.844 3.844a27.13 27.13 0 0 1-.023 4.612M13.457 2.11l8.433 8.433M13.457 2.11c-1.489.143-3.026.443-4.457.992m12.89 7.44c-.277 2.884-1.142 5.949-3.27 8.077m-8.077 3.271c-1.67.16-3.28.125-4.612.023a4.136 4.136 0 0 1-3.844-3.844a27.141 27.141 0 0 1 .023-4.612m8.433 8.433L2.11 13.457m8.433 8.433c1.489-.143 3.026-.443 4.457-.992m.31-12.208l-6.62 6.62m0-2.837l2.837 2.837m-.945-4.728l2.836 2.836m-.945-4.728l2.837 2.837"/>`),
		g.Group(children),
	)
}