package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoLuggageOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 16.125l-2-2V8h-6.125l-2-2V2H15v4h4v10.125Zm-3-3l-1.5-1.5V9H16v4.125Zm-3.25-3.25L11.875 9h.875v.875ZM10.375 6H13.5V3.5h-3.125V6ZM9 22H7v-1H5V6.175h1.2L8.025 8H7v11h10v-2.025l2 2V21h-2v1h-2v-1H9v1Zm-1-4V9.475h1.5V18H8Zm3.25 0v-6.775l1.5 1.525V18h-1.5Zm3.25 0v-3.5L16 16v2h-1.5Zm-.55-6.95ZM11.6 14.4Zm8.875 8.9L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425Z"/>`),
		g.Group(children),
	)
}