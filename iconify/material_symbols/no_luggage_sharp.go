package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoLuggageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.375 6H13.5V3.5h-3.125V6ZM7 22v-1H5V6.175h1.2l12.8 12.8V21h-2v1h-2v-1H9v1H7Zm1-4h1.5V9.475H8V18Zm3.25 0h1.5v-5.25l-1.5-1.525V18Zm3.25 0H16v-2l-1.5-1.5V18Zm5.975 5.3L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM19 16.125l-3-3V9h-1.5v2.625l-1.75-1.75V9h-.875l-3-3V2H15v4h4v10.125Z"/>`),
		g.Group(children),
	)
}