package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#FFA000" d="M13.324 45.003L3.002 34.68L34.676 3.007L44.998 13.33z"/><path fill="#9E6400" d="m22.803 24.188l-4.666-4.666l1.414-1.414l4.666 4.666zm2.01-5.99l-2.616-2.616l1.414-1.414l2.616 2.616zm5.991-2.01l-4.666-4.666l1.414-1.414l4.666 4.666zm-.649-8.645l1.415-1.414l2.615 2.616l-1.414 1.414zM8.81 34.198l-2.616-2.616l1.414-1.414l2.616 2.616zm5.991-2.01l-4.666-4.666l1.414-1.414l4.666 4.666zm2.011-5.99l-2.616-2.616l1.414-1.414l2.616 2.616z"/>`),
		g.Group(children),
	)
}