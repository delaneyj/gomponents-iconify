package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.79 4.46l.71-.71L20.25 20.5l-.71.71L14.33 16H12v4.5l-.5 1.5l-.5-1.5V16H6v-2l2-2V9.66l-5.21-5.2M14 12.41V5h1V4H8v1h1v2.84l-1-1V6h-.84L7 5.84V3h9v3h-1v6l2 2v1.83l-1-.99v-.43l-2-2m-5 0l-2 2V15h6.34L9 10.66v1.75Z"/>`),
		g.Group(children),
	)
}