package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L16 18.8l-4 4L1.2 12l4-4L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM14.6 17.4l-8-8L4 12l8 8l2.6-2.6Zm4.25-1.45l-1.4-1.4L20 12l-8-8l-2.55 2.55l-1.4-1.4L12 1.2L22.8 12l-3.95 3.95Zm-5.4-5.4ZM10.6 13.4ZM8 15v-5h2.025l2 2H10v3H8Zm7.45-2.45L17 11l-3.5-3.5V10h-.6l2.55 2.55Z"/>`),
		g.Group(children),
	)
}