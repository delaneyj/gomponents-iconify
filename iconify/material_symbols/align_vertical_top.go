package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22V6h3v16H7Zm7-6V6h3v10h-3ZM2 4V2h20v2H2Z"/>`),
		g.Group(children),
	)
}