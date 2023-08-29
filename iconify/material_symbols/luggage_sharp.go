package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21H5V6h4V2h6v4h4v15h-2v1h-2v-1H9v1H7v-1Zm3.5-15h3V3.5h-3V6ZM8 18h1.5V9H8v9Zm3.25 0h1.5V9h-1.5v9Zm3.25 0H16V9h-1.5v9Z"/>`),
		g.Group(children),
	)
}