package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalletShoes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#ffa7c0" d="M10 54.904c-2.248-1.853-3-8.255-3-11.396c0-2.79 1.13-4.315 2.959-6.144c0 0 2.658-2.132 4.041-1c4.529 3.71 15.658 15.215 27.004 3c0 0 15.577 2.343 20.32 3.546C67 45 65.741 45.03 65.741 47.733c0 8.651-6.084 7.539-12.81 7.171c-3.707-.202-38.801.226-42.93 0"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 54.904c-2.248-1.853-3-8.255-3-11.396c0-2.79 1.13-4.315 2.959-6.144c0 0 2.658-2.132 4.041-1c4.529 3.71 15.658 14.215 27.004 2c0 0 15.996 2.636 20.32 4.546C67 45 65.741 45.03 65.741 47.733c0 8.651-6.084 7.539-12.81 7.171c-3.707-.202-38.801.226-42.93 0m36.525-15.102l-5.071 5.938m.915-6.136l3.533 4.136"/>`),
		g.Group(children),
	)
}