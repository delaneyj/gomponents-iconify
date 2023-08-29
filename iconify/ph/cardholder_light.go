package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardholderLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 50H48a22 22 0 0 0-22 22v112a22 22 0 0 0 22 22h160a22 22 0 0 0 22-22V72a22 22 0 0 0-22-22ZM38 94h180v20h-50.81a14 14 0 0 0-13.71 11.21a26 26 0 0 1-51 0A14 14 0 0 0 88.81 114H38Zm10-32h160a10 10 0 0 1 10 10v10H38V72a10 10 0 0 1 10-10Zm160 132H48a10 10 0 0 1-10-10v-58h50.81a2 2 0 0 1 2 1.59a38 38 0 0 0 74.48 0a2 2 0 0 1 1.95-1.59H218v58a10 10 0 0 1-10 10Z"/>`),
		g.Group(children),
	)
}