package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3h1l4.85 4.85l-4.64 4.65l4.64 4.65L12 22h-1v-8.29l-4.45 4.45l-.71-.71l4.95-4.95l-4.95-4.95l.71-.71L11 11.29V3m1 1.41v6.88l3.44-3.44L12 4.41m0 16.18l3.44-3.44L12 13.71v6.88Z"/>`),
		g.Group(children),
	)
}