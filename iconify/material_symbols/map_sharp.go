package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 21l-6-2.1l-6 2.325V5.05L9 3l6 2.1l6-2.325V18.95L15 21Zm-1-2.45V6.85l-4-1.4v11.7l4 1.4Z"/>`),
		g.Group(children),
	)
}