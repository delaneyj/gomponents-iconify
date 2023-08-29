package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Biomass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#9CCC65" d="M32 15V7H16v8L6.2 40c-.6 1.5.5 3 2.1 3h31.5c1.6 0 2.6-1.6 2.1-3L32 15z"/><path fill="#8BC34A" d="M32 9H16c-1.1 0-2-.9-2-2s.9-2 2-2h16c1.1 0 2 .9 2 2s-.9 2-2 2z"/><path fill="#2E7D32" d="M28 30c0 4.4-4 8-4 8s-4-3.6-4-8s4-8 4-8s4 3.6 4 8z"/><path fill="#388E3C" d="M31.1 32.6c-2 4-7.1 5.4-7.1 5.4s-2-5 0-8.9s7.1-5.4 7.1-5.4s2 4.9 0 8.9z"/><path fill="#43A047" d="M16.9 32.6c2 4 7.1 5.4 7.1 5.4s2-5 0-8.9s-7.1-5.4-7.1-5.4s-2 4.9 0 8.9z"/>`),
		g.Group(children),
	)
}